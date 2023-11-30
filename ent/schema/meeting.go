package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Meeting holds the schema definition for the Meeting entity.
type Meeting struct {
	ent.Schema
}

// Fields of the Meeting.
func (Meeting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(), // 会议ID
		field.String("name"),       // 会议名称
		field.String("applicant"),  // 申请人
	}
}

// Edges of the Meeting.
func (Meeting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mdrs", MeetingDateRoom.Type),
	}
}
