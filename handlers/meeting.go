package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mrm/ent/meeting"
	"mrm/ent/meetingdateroom"
	"net/http"
	"strconv"
)

type ResponseMeetingEx struct {
	MeetingID   int    `json:"meeting_id"`
	MeetingName string `json:"meeting_name"`
	Applicant   string `json:"applicant"`

	DateTimes []ResponseMeetingDateTime `json:"date_times"`
}

type ResponseMeetingDateTime struct {
	Date      int `json:"date"`
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`

	RoomID   int    `json:"room_id"`
	RoomName string `json:"room_name"`
}

func (h *Handler) QueryMeeting(c *gin.Context) {
	meetingIdStr := c.Param("id")
	meetingId, err := strconv.Atoi(meetingIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MeetingID convert to int error: " + err.Error()})
		return
	}

	mdrs, err := h.DbClient.MeetingDateRoom.Query().
		Where(meetingdateroom.HasMeetingWith(meeting.ID(meetingId))).
		WithRoom().WithMeeting().All(c)
	if err != nil || len(mdrs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Query error: " + err.Error()})
		return
	}

	res := ResponseMeetingEx{
		MeetingID:   mdrs[0].Edges.Meeting.ID,
		MeetingName: mdrs[0].Edges.Meeting.Name,
		Applicant:   mdrs[0].Edges.Meeting.Applicant,
	}

	for _, mdr := range mdrs {
		res.DateTimes = append(res.DateTimes, ResponseMeetingDateTime{
			Date:      mdr.Date,
			StartTime: mdr.StartTime,
			EndTime:   mdr.EndTime,
			RoomID:    mdr.Edges.Room.ID,
			RoomName:  mdr.Edges.Room.Name,
		})
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteMeeting(c *gin.Context) {
	meetingIdStr := c.Param("id")
	meetingId, err := strconv.Atoi(meetingIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MeetingID convert to int error: " + err.Error()})
		return
	}

	code, obj := h.deleteMeeting(meetingId, c)
	c.JSON(code, obj)
	return
}

func (h *Handler) deleteMeeting(meetingId int, c *gin.Context) (code int, obj any) {
	// Delete Meeting and MeetingDateRoom
	tx, err := h.DbClient.Tx(c)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Tx error: %s", err.Error())}
	}

	// delete meetingDateRoom
	_, err = tx.MeetingDateRoom.Delete().Where(meetingdateroom.HasMeetingWith(meeting.ID(meetingId))).Exec(c)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Delete meetingDateRoom error: %v", err)}
	}

	// delete meeting
	_, err = tx.Meeting.Delete().Where(meeting.ID(meetingId)).Exec(c)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Delete meeting error: %v", err)}
	}

	// commit
	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Commit error: %v", err)}
	}
	return http.StatusOK, gin.H{"message": "Delete success"}
}
