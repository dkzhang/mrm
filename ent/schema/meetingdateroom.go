package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MeetingDateRoom holds the schema definition for the MeetingDateRoom entity.
type MeetingDateRoom struct {
	ent.Schema
}

// Fields of the MeetingDateRoom.
func (MeetingDateRoom) Fields() []ent.Field {
	return []ent.Field{
		field.Int("date"),       // 会议日期
		field.Int("start_time"), // 会议开始时间
		field.Int("end_time"),   // 会议结束时间
	}
}

// Edges of the MeetingDateRoom.
func (MeetingDateRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("meeting", Meeting.Type).
			Ref("mdrs").
			Unique(),
		edge.From("room", Room.Type).
			Ref("mdrs").
			Unique(),
	}
}

func (MeetingDateRoom) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("date"),
	}
}
