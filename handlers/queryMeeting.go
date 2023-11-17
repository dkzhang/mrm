package handlers

type ResponseMeetingEx struct {
	MeetingID   int
	MeetingName string
	Applicant   string
}

type ResponseMeetingDateTime struct {
	Date      int
	StartTime int
	EndTime   int

	RoomID   int
	RoomName string
}
