package handlers

import (
	"github.com/gin-gonic/gin"
	"mrm/ent/meetingdateroom"
	"mrm/ent/room"
	"net/http"
	"strconv"
)

type ResponseMeeting struct {
	MeetingID   int    `json:"meeting_id"`
	MeetingName string `json:"meeting_name"`
	Applicant   string `json:"applicant"`
	Date        int    `json:"date"`
	StartTime   int    `json:"start_time"`
	EndTime     int    `json:"end_time"`
	RoomID      int    `json:"room_id"`
	RoomName    string `json:"room_name"`
}

func (h *Handler) QueryAllocateByDate(c *gin.Context) {
	meetingDateStr := c.Param("date")
	meetingDate, err := strconv.Atoi(meetingDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mdrs, err := h.DbClient.MeetingDateRoom.Query().
		Where(meetingdateroom.Date(meetingDate)).WithRoom().WithMeeting().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res []ResponseMeeting
	for _, mdr := range mdrs {
		res = append(res, ResponseMeeting{
			MeetingID:   mdr.Edges.Meeting.ID,
			MeetingName: mdr.Edges.Meeting.Name,
			Applicant:   mdr.Edges.Meeting.Applicant,
			Date:        mdr.Date,
			StartTime:   mdr.StartTime,
			EndTime:     mdr.EndTime,
			RoomID:      mdr.Edges.Room.ID,
			RoomName:    mdr.Edges.Room.Name,
		})
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) QueryAllocateByRoom(c *gin.Context) {
	// get params and query
	roomIdStr := c.Param("id")
	fromStr := c.DefaultQuery("from", "20000101")
	toStr := c.DefaultQuery("to", "20991231")

	// convert string to int
	roomID, err := strconv.Atoi(roomIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomID convert to int error: " + err.Error()})
		return
	}
	from, err := strconv.Atoi(fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "<from> convert to int error: " + err.Error()})
		return
	}
	to, err := strconv.Atoi(toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "<to> convert to int error: " + err.Error()})
		return
	}

	// query in db
	mdrs, err := h.DbClient.Room.Query().Where(room.ID(roomID)).QueryMdrs().WithRoom().WithMeeting().All(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// convert to response
	var res []ResponseMeeting
	for _, mdr := range mdrs {
		if mdr.Date >= from && mdr.Date <= to {
			res = append(res, ResponseMeeting{
				MeetingID:   mdr.Edges.Meeting.ID,
				MeetingName: mdr.Edges.Meeting.Name,
				Applicant:   mdr.Edges.Meeting.Applicant,
				Date:        mdr.Date,
				StartTime:   mdr.StartTime,
				EndTime:     mdr.EndTime,
				RoomID:      mdr.Edges.Room.ID,
				RoomName:    mdr.Edges.Room.Name,
			})
		}
	}

	c.JSON(http.StatusOK, res)
}
