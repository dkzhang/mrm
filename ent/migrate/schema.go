// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MeetingsColumns holds the columns for the "meetings" table.
	MeetingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "applicant", Type: field.TypeString},
	}
	// MeetingsTable holds the schema information for the "meetings" table.
	MeetingsTable = &schema.Table{
		Name:       "meetings",
		Columns:    MeetingsColumns,
		PrimaryKey: []*schema.Column{MeetingsColumns[0]},
	}
	// MeetingDateRoomsColumns holds the columns for the "meeting_date_rooms" table.
	MeetingDateRoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeInt},
		{Name: "start_time", Type: field.TypeInt},
		{Name: "end_time", Type: field.TypeInt},
		{Name: "meeting_mdrs", Type: field.TypeInt64, Nullable: true},
		{Name: "room_mdrs", Type: field.TypeInt, Nullable: true},
	}
	// MeetingDateRoomsTable holds the schema information for the "meeting_date_rooms" table.
	MeetingDateRoomsTable = &schema.Table{
		Name:       "meeting_date_rooms",
		Columns:    MeetingDateRoomsColumns,
		PrimaryKey: []*schema.Column{MeetingDateRoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "meeting_date_rooms_meetings_mdrs",
				Columns:    []*schema.Column{MeetingDateRoomsColumns[4]},
				RefColumns: []*schema.Column{MeetingsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "meeting_date_rooms_rooms_mdrs",
				Columns:    []*schema.Column{MeetingDateRoomsColumns[5]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "meetingdateroom_date",
				Unique:  false,
				Columns: []*schema.Column{MeetingDateRoomsColumns[1]},
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MeetingsTable,
		MeetingDateRoomsTable,
		RoomsTable,
	}
)

func init() {
	MeetingDateRoomsTable.ForeignKeys[0].RefTable = MeetingsTable
	MeetingDateRoomsTable.ForeignKeys[1].RefTable = RoomsTable
}
