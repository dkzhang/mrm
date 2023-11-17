package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mrm/ent"
	"mrm/ent/meetingdateroom"
	"mrm/ent/room"
	"net/http"
	"strconv"
)

const Intervals = 24

func (h *Handler) SVG(c *gin.Context) {
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

	rooms, err := h.DbClient.Room.Query().Order(ent.Asc(room.FieldID)).All(c)

	roomOccupiedArrays := make([]roomOccupiedArray, len(rooms))
	for i, room := range rooms {
		roomOccupiedArrays[i].roomID = room.ID
		roomOccupiedArrays[i].roomName = room.Name
		roomOccupiedArrays[i].Occupied = make([]int, Intervals)
	}

	for _, mdr := range mdrs {
		for i, room := range rooms {
			if room.ID == mdr.Edges.Room.ID {
				roomOccupiedArrays[i].Occupied = mergeArray(roomOccupiedArrays[i].Occupied, t2t(mdr.StartTime, mdr.EndTime, mdr.Edges.Meeting.ID))
			}
		}
	}
	fmt.Printf("%v\n", roomOccupiedArrays)

	c.JSON(http.StatusOK, roomOccupiedArrays)

}

type roomOccupiedArray struct {
	roomID   int
	roomName string
	Occupied []int
}

func t2i(t int) int {
	// Convert time 700 to 0, 730 to 1, like this
	return t/100*2 + t%100/30 - 7*2
}

func t2t(from int, to int, id int) []int {
	array := make([]int, Intervals)
	for i := t2i(from); i < t2i(to); i++ {
		array[i] = id
	}
	return array
}

func mergeArray(a []int, b []int) []int {
	array := make([]int, Intervals)
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			array[i] = b[i]
		} else {
			array[i] = a[i]
		}
	}
	return array
}
