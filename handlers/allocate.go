package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mrm/ent"
	"mrm/ent/room"
	"net/http"
)

type AllocateMeeting struct {
	ID        int               `json:"id"`
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

	if err := c.ShouldBindJSON(&am); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check DateTimes []MeetingDateTime is valid

	// DateTime conflict detection
	var conflictedMdrs []*ent.MeetingDateRoom
	for _, dt := range am.DateTimes {
		mdrs, err := h.DbClient.Room.Query().Where(room.ID(dt.RoomID)).QueryMdrs().All(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, mdr := range mdrs {
			if IsConflict(&dt, mdr) {
				if am.IsMandatory {
					conflictedMdrs = append(conflictedMdrs, mdr)
				} else {
					c.JSON(http.StatusConflict, gin.H{"error": "DateTime Conflict"})
					return
				}
			}
		}
	}

	// Create Meeting and MeetingDateRoom
	tx, err := h.DbClient.Tx(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create meeting.
	meeting, err := tx.Meeting.Create().
		SetID(am.ID).
		SetName(am.Name).
		SetApplicant(am.Applicant).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Create MeetingDateRoom error: %s", err.Error())})
			return
		}
	}

	// remove conflicted meetingDateRoom if isMandatory is true.
	if am.IsMandatory {
		for _, mdr := range conflictedMdrs {
			err := tx.MeetingDateRoom.DeleteOne(mdr).Exec(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	// commit the transaction.
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
