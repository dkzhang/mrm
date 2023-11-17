package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(), // 会议室ID
		field.String("name"),     // 会议室名称
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mdrs", MeetingDateRoom.Type),
	}
}
