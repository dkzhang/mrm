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
	titleHeight   = 50

	tableWidth = rowHeaderWidth + cellWidth*Intervals + tableSpacingX*2

	startHour = 7

	colorsNum = 16
)

type Meeting struct {
	Name      string
	Applicant string
}

func (h *Handler) SVG(c *gin.Context) {
	meetingDateStr := c.Param("date")
	meetingDate, err := strconv.Atoi(meetingDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	mdrs, err := h.DbClient.MeetingDateRoom.Query().
		Where(meetingdateroom.Date(meetingDate)).WithRoom().WithMeeting().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	rooms, err := h.DbClient.Room.Query().Order(ent.Asc(room.FieldID)).All(c)

	roomOccupiedArrays := make([]roomOccupiedArray, len(rooms))
	for i, room := range rooms {
		roomOccupiedArrays[i].roomID = room.ID
		roomOccupiedArrays[i].roomName = room.Name
		roomOccupiedArrays[i].Occupied = make([]int64, Intervals)
	}

	meetings := make(map[int64]Meeting)

	for _, mdr := range mdrs {
		for i, room := range rooms {
			if room.ID == mdr.Edges.Room.ID {
				roomOccupiedArrays[i].Occupied = mergeArray(roomOccupiedArrays[i].Occupied, t2t(mdr.StartTime, mdr.EndTime, mdr.Edges.Meeting.ID))
				if _, ok := meetings[mdr.Edges.Meeting.ID]; !ok {
					meetings[mdr.Edges.Meeting.ID] = Meeting{
						Name:      mdr.Edges.Meeting.Name,
						Applicant: mdr.Edges.Meeting.Applicant,
					}
				}
			}
		}
	}

	svgTitle := fmt.Sprintf("%d年%d月%d日会议室占用情况", meetingDate/10000, meetingDate%10000/100, meetingDate%100)
	buffer := GenSvg(svgTitle, roomOccupiedArrays, meetings)
	c.Data(http.StatusOK, "image/svg+xml", buffer.Bytes())
}

func GenSvg(svgTitle string, rooms []roomOccupiedArray, meetings map[int64]Meeting) *bytes.Buffer {
	tableHeight := titleHeight + colHeaderHeight + cellHeight*len(rooms) + tableSpacingY*2

	buffer := bytes.NewBufferString(fmt.Sprintf(SVG_HEAD, tableWidth, tableHeight, tableWidth, tableHeight))

	// Style
	buffer.WriteString(getStyle())

	// Title
	buffer.WriteString(fmt.Sprintf(`<text x="%d" y="%d" class="title">%s</text>`+"\n", tableWidth/2, titleHeight/2, svgTitle))

	// Col Header
	currentX := rowHeaderWidth + tableSpacingX
	currentY := titleHeight + tableSpacingY
	for i := 0; i < Intervals/IntervalsPerHour; i++ {
		buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="header col-header" />`+"\n",
			currentX+colHeaderWidth*i, currentY,
			colHeaderWidth, colHeaderHeight,
		))
		buffer.WriteString(fmt.Sprintf(`<text x="%d" y="%d" class="header-text col-header-text">%d</text>`+"\n",
			currentX+colHeaderWidth*i+colHeaderWidth/2, currentY+colHeaderHeight/2, startHour+i))
	}

	// Row Header
	currentX = tableSpacingX
	currentY = titleHeight + colHeaderHeight + tableSpacingY

	for i, room := range rooms {
		buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="header row-header" />`+"\n",
			currentX, currentY+cellHeight*i,
			rowHeaderWidth, cellHeight))
		buffer.WriteString(fmt.Sprintf(`<text x="%d" y="%d" class="header-text row-header-text">%s</text>`+"\n",
			currentX+rowHeaderWidth/2, currentY+cellHeight*i+cellHeight/2, room.roomName))
	}

	// cells
	currentX = rowHeaderWidth + tableSpacingX
	currentY = titleHeight + colHeaderHeight + tableSpacingY
	currentMeetingID := int64(0)
	currentColorIndex := 0
	for i, room := range rooms {
		for j := 0; j < Intervals; j++ {
			if room.Occupied[j] > 0 {
				if currentMeetingID != room.Occupied[j] {
					// change color
					currentColorIndex = (currentColorIndex + 1) % colorsNum
				}
				buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" data-name="%s" data-applicant="%s" class="cell cell-occupied color-fill-%d" />`+"\n",
					currentX+cellWidth*j, currentY+cellHeight*i,
					cellWidth, cellHeight,
					meetings[room.Occupied[j]].Name, meetings[room.Occupied[j]].Applicant,
					currentColorIndex))
			} else {
				buffer.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="cell cell-idle"/>`+"\n",
					currentX+cellWidth*j, currentY+cellHeight*i,
					cellWidth, cellHeight))
			}
			currentMeetingID = room.Occupied[j]
		}
	}

	// insert newRect
	buffer.WriteString(`
	<rect id="newRect" x="0" y="0" width="100" height="50" style="fill:yellow; visibility:hidden"/>
    <text id="textName" x="0" y="0" visibility="hidden"></text>
    <text id="textApplicant" text-anchor="middle" x="0" y="0" visibility="hidden"></text>
`)

	// SVG END
	buffer.WriteString(SVG_END)

	return buffer
}

type roomOccupiedArray struct {
	roomID   int
	roomName string
	Occupied []int64
}

func t2i(t int) int {
	// Convert time 700 to 0, 730 to 1, like this
	return t/100*2 + t%100/30 - 7*2
}

func t2t(from int, to int, id int64) []int64 {
	array := make([]int64, Intervals)
	for i := t2i(from); i < t2i(to); i++ {
		array[i] = id
	}
	return array
}

func mergeArray(a []int64, b []int64) []int64 {
	array := make([]int64, Intervals)
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
