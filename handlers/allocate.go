package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mrm/ent"
	"mrm/ent/meeting"
	"mrm/ent/meetingdateroom"
	"mrm/ent/room"
	"net/http"
	"time"
)

type RespA struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AllocateMeeting struct {
	ID        int64             `json:"id"`
	Name      string            `json:"name"`
	Applicant string            `json:"applicant"`
	DateTimes []MeetingDateTime `json:"date_times"`

	IsMandatory bool `json:"is_mandatory"` // 是否强制分配
}

type MeetingDateTime struct {
	Date      int `json:"date"`
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`

	RoomID int `json:"room_id"`
}

func (h *Handler) Allocate(c *gin.Context) {
	am := AllocateMeeting{}

	var code int
	var message string

	if err := c.ShouldBindJSON(&am); err != nil {
		code = http.StatusBadRequest
		message = fmt.Sprintf("Bind JSON error: %s", err.Error())
		return
	}

	// Check DateTimes []MeetingDateTime is valid
	for _, dt := range am.DateTimes {
		dateStr := fmt.Sprintf("%d", dt.Date)
		_, err := time.Parse("20060102", dateStr)
		if err != nil {
			code = http.StatusBadRequest
			message = fmt.Sprintf("MeetingDateTime Date is not in YYYYMMDD integer format")
			return
		}

		if dt.StartTime >= dt.EndTime {
			code = http.StatusBadRequest
			message = fmt.Sprintf("MeetingDateTime StartTime >= EndTime")
			return
		}
		if dt.StartTime < 0 || dt.StartTime > 2400 || dt.EndTime < 0 || dt.EndTime > 2400 {
			code = http.StatusBadRequest
			message = fmt.Sprintf("MeetingDateTime StartTime or EndTime is not in [0, 2400]")
			return
		}
	}

	// create tx
	tx, err := h.DbClient.Tx(c)
	if err != nil {
		code = http.StatusInternalServerError
		message = fmt.Sprintf("Tx error: %s", err.Error())
		return
	}

	// rollback if panic or error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, RespA{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("Panic: %v", p),
			})
			panic(p) // re-throw panic after Rollback
		} else if code != http.StatusOK {
			// code is http error code
			tx.Rollback()
			c.JSON(http.StatusOK, RespA{
				Code:    code,
				Message: message,
			})
		} else {
			err = tx.Commit() // code is http.StatusOK, commit the transaction.
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, RespA{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("Commit error: %v", err),
				})
			} else {
				c.JSON(http.StatusOK, RespA{
					Code:    http.StatusOK,
					Message: fmt.Sprintf("Allocate success"),
				})
			}
		}
	}()

	// check meeting exists
	_, err = tx.Meeting.Query().Where(meeting.ID(am.ID)).Only(c)
	if err != nil {
		if !ent.IsNotFound(err) {
			code = http.StatusInternalServerError
			message = fmt.Sprintf("Query meeting error: %s", err.Error())
			return
		}
	} else {
		// meeting exists, delete it.
		// delete meetingDateRoom
		_, err = tx.MeetingDateRoom.Delete().Where(meetingdateroom.HasMeetingWith(meeting.ID(am.ID))).Exec(c)
		if err != nil {
			code = http.StatusInternalServerError
			message = fmt.Sprintf("Delete meetingDateRoom error: %v", err)
			return
		}

		// delete meeting
		_, err = tx.Meeting.Delete().Where(meeting.ID(am.ID)).Exec(c)
		if err != nil {
			code = http.StatusInternalServerError
			message = fmt.Sprintf("Delete meeting error: %v", err)
			return
		}
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// DateTime conflict detection
	var conflictedMdrs []*ent.MeetingDateRoom
	for _, dt := range am.DateTimes {
		mdrs, err := tx.Room.Query().Where(room.ID(dt.RoomID)).
			QueryMdrs().All(c)
		if err != nil {
			code = http.StatusInternalServerError
			message = fmt.Sprintf("Query mdrs error: %s", err.Error())
			return
		}

		for _, mdr := range mdrs {
			// continue if in the same meeting.
			m, err := mdr.QueryMeeting().Only(c)
			if err != nil {
				code = http.StatusInternalServerError
				message = fmt.Sprintf("Query meeting ID from mdr error: %s", err.Error())
				return
			}
			if m.ID == am.ID {
				continue
			}

			if IsConflict(&dt, mdr) {
				if am.IsMandatory {
					conflictedMdrs = append(conflictedMdrs, mdr)
				} else {
					code = http.StatusConflict
					message = fmt.Sprintf("Meeting DateTime Conflict")
					return
				}
			}
		}
	}

	// create meeting.
	meeting, err := tx.Meeting.Create().
		SetID(am.ID).
		SetName(am.Name).
		SetApplicant(am.Applicant).
		Save(c)
	if err != nil {
		code = http.StatusInternalServerError
		message = fmt.Sprintf("Create meeting error: %s", err.Error())
		return
	}

	// create meetingDateRoom.
	for _, dt := range am.DateTimes {
		_, err := tx.MeetingDateRoom.Create().
			SetDate(dt.Date).
			SetStartTime(dt.StartTime).
			SetEndTime(dt.EndTime).
			SetMeeting(meeting).
			SetRoomID(dt.RoomID).
			Save(c)
		if err != nil {
			code = http.StatusInternalServerError
			message = fmt.Sprintf("Create meetingDateRoom error: %s", err.Error())
			return
		}
	}

	// remove conflicted meetingDateRoom if isMandatory is true.
	if am.IsMandatory {
		for _, mdr := range conflictedMdrs {
			err := tx.MeetingDateRoom.DeleteOne(mdr).Exec(c)
			if err != nil {
				code = http.StatusInternalServerError
				message = fmt.Sprintf("Delete conflicted meetingDateRoom error: %s", err.Error())
				return
			}
		}
	}

	code = http.StatusOK
	return
}

func IsConflict(a *MeetingDateTime, b *ent.MeetingDateRoom) bool {
	if a.Date != b.Date {
		return false
	}
	if a.StartTime >= b.EndTime || a.EndTime <= b.StartTime {
		return false
	}
	return true
}
