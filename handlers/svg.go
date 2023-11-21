package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"mrm/ent"
	"mrm/ent/meetingdateroom"
	"mrm/ent/room"
	"net/http"
	"strconv"
)

const (
	Intervals        = 24
	IntervalsPerHour = 2

	cellWidth       = 30
	cellHeight      = 30
	colHeaderHeight = cellHeight
	colHeaderWidth  = cellWidth * IntervalsPerHour
	rowHeaderHeight = cellHeight
	rowHeaderWidth  = 100

	tableSpacingX = 5
	tableSpacingY = 5

	tableWidth = rowHeaderWidth + cellWidth*Intervals + tableSpacingX*2

	startHour = 7

	colorsNum = 16
)

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

	buffer := GenSvg(roomOccupiedArrays)
	c.Data(http.StatusOK, "image/svg+xml", buffer.Bytes())
}

func GenSvg(rooms []roomOccupiedArray) *bytes.Buffer {
	tableHeight := colHeaderHeight + cellHeight*len(rooms) + tableSpacingY*2

	buffer := bytes.NewBufferString(fmt.Sprintf(SVG_HEAD, tableWidth, tableHeight, tableWidth, tableHeight))

	// Style
	buffer.WriteString(getStyle())

	// Col Header
	currentX := rowHeaderWidth + tableSpacingX
	currentY := tableSpacingY
	for i := 0; i < Intervals/IntervalsPerHour; i++ {
		buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="header col-header" />`+"\n",
			currentX+colHeaderWidth*i, currentY,
			colHeaderWidth, colHeaderHeight))
		buffer.WriteString(fmt.Sprintf(`<text x="%d" y="%d" class="header-text col-header-text">%d</text>`+"\n",
			currentX+colHeaderWidth*i+colHeaderWidth/2, currentY+colHeaderHeight/2, startHour+i))
	}

	// Row Header
	currentX = tableSpacingX
	currentY = colHeaderHeight + tableSpacingY

	for i, room := range rooms {
		buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="header row-header" />`+"\n",
			currentX, currentY+cellHeight*i,
			rowHeaderWidth, cellHeight))
		buffer.WriteString(fmt.Sprintf(`<text x="%d" y="%d" class="header-text row-header-text">%s</text>`+"\n",
			currentX+rowHeaderWidth/2, currentY+cellHeight*i+cellHeight/2, room.roomName))
	}

	// cells
	currentX = rowHeaderWidth + tableSpacingX
	currentY = colHeaderHeight + tableSpacingY
	currentMeetingID := 0
	currentColorIndex := 0
	for i, room := range rooms {
		for j := 0; j < Intervals; j++ {
			if room.Occupied[j] > 0 {
				if currentMeetingID != room.Occupied[j] {
					// change color
					currentColorIndex = (currentColorIndex + 1) % colorsNum
				}
				buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="cell cell-occupied color-fill-%d" />`+"\n",
					currentX+cellWidth*j, currentY+cellHeight*i,
					cellWidth, cellHeight, currentColorIndex))
			} else {
				buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="cell cell-idle"/>`+"\n",
					currentX+cellWidth*j, currentY+cellHeight*i,
					cellWidth, cellHeight))
			}
			currentMeetingID = room.Occupied[j]
		}
	}

	// SVG END
	buffer.WriteString(SVG_END)

	return buffer
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

const SVG_HEAD = `<svg preserveAspectRatio="xMinYMim meet" viewBox="0 0 %d %d" width="%d" height="%d"   xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`
const SVG_END = `</svg>`
