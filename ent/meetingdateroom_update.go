// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mrm/ent/meeting"
	"mrm/ent/meetingdateroom"
	"mrm/ent/predicate"
	"mrm/ent/room"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MeetingDateRoomUpdate is the builder for updating MeetingDateRoom entities.
type MeetingDateRoomUpdate struct {
	config
	hooks    []Hook
	mutation *MeetingDateRoomMutation
}

// Where appends a list predicates to the MeetingDateRoomUpdate builder.
func (mdru *MeetingDateRoomUpdate) Where(ps ...predicate.MeetingDateRoom) *MeetingDateRoomUpdate {
	mdru.mutation.Where(ps...)
	return mdru
}

// SetDate sets the "date" field.
func (mdru *MeetingDateRoomUpdate) SetDate(i int) *MeetingDateRoomUpdate {
	mdru.mutation.ResetDate()
	mdru.mutation.SetDate(i)
	return mdru
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (mdru *MeetingDateRoomUpdate) SetNillableDate(i *int) *MeetingDateRoomUpdate {
	if i != nil {
		mdru.SetDate(*i)
	}
	return mdru
}

// AddDate adds i to the "date" field.
func (mdru *MeetingDateRoomUpdate) AddDate(i int) *MeetingDateRoomUpdate {
	mdru.mutation.AddDate(i)
	return mdru
}

// SetStartTime sets the "start_time" field.
func (mdru *MeetingDateRoomUpdate) SetStartTime(i int) *MeetingDateRoomUpdate {
	mdru.mutation.ResetStartTime()
	mdru.mutation.SetStartTime(i)
	return mdru
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (mdru *MeetingDateRoomUpdate) SetNillableStartTime(i *int) *MeetingDateRoomUpdate {
	if i != nil {
		mdru.SetStartTime(*i)
	}
	return mdru
}

// AddStartTime adds i to the "start_time" field.
func (mdru *MeetingDateRoomUpdate) AddStartTime(i int) *MeetingDateRoomUpdate {
	mdru.mutation.AddStartTime(i)
	return mdru
}

// SetEndTime sets the "end_time" field.
func (mdru *MeetingDateRoomUpdate) SetEndTime(i int) *MeetingDateRoomUpdate {
	mdru.mutation.ResetEndTime()
	mdru.mutation.SetEndTime(i)
	return mdru
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (mdru *MeetingDateRoomUpdate) SetNillableEndTime(i *int) *MeetingDateRoomUpdate {
	if i != nil {
		mdru.SetEndTime(*i)
	}
	return mdru
}

// AddEndTime adds i to the "end_time" field.
func (mdru *MeetingDateRoomUpdate) AddEndTime(i int) *MeetingDateRoomUpdate {
	mdru.mutation.AddEndTime(i)
	return mdru
}

// SetMeetingID sets the "meeting" edge to the Meeting entity by ID.
func (mdru *MeetingDateRoomUpdate) SetMeetingID(id int64) *MeetingDateRoomUpdate {
	mdru.mutation.SetMeetingID(id)
	return mdru
}

// SetNillableMeetingID sets the "meeting" edge to the Meeting entity by ID if the given value is not nil.
func (mdru *MeetingDateRoomUpdate) SetNillableMeetingID(id *int64) *MeetingDateRoomUpdate {
	if id != nil {
		mdru = mdru.SetMeetingID(*id)
	}
	return mdru
}

// SetMeeting sets the "meeting" edge to the Meeting entity.
func (mdru *MeetingDateRoomUpdate) SetMeeting(m *Meeting) *MeetingDateRoomUpdate {
	return mdru.SetMeetingID(m.ID)
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (mdru *MeetingDateRoomUpdate) SetRoomID(id int) *MeetingDateRoomUpdate {
	mdru.mutation.SetRoomID(id)
	return mdru
}

// SetNillableRoomID sets the "room" edge to the Room entity by ID if the given value is not nil.
func (mdru *MeetingDateRoomUpdate) SetNillableRoomID(id *int) *MeetingDateRoomUpdate {
	if id != nil {
		mdru = mdru.SetRoomID(*id)
	}
	return mdru
}

// SetRoom sets the "room" edge to the Room entity.
func (mdru *MeetingDateRoomUpdate) SetRoom(r *Room) *MeetingDateRoomUpdate {
	return mdru.SetRoomID(r.ID)
}

// Mutation returns the MeetingDateRoomMutation object of the builder.
func (mdru *MeetingDateRoomUpdate) Mutation() *MeetingDateRoomMutation {
	return mdru.mutation
}

// ClearMeeting clears the "meeting" edge to the Meeting entity.
func (mdru *MeetingDateRoomUpdate) ClearMeeting() *MeetingDateRoomUpdate {
	mdru.mutation.ClearMeeting()
	return mdru
}

// ClearRoom clears the "room" edge to the Room entity.
func (mdru *MeetingDateRoomUpdate) ClearRoom() *MeetingDateRoomUpdate {
	mdru.mutation.ClearRoom()
	return mdru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mdru *MeetingDateRoomUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mdru.sqlSave, mdru.mutation, mdru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mdru *MeetingDateRoomUpdate) SaveX(ctx context.Context) int {
	affected, err := mdru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mdru *MeetingDateRoomUpdate) Exec(ctx context.Context) error {
	_, err := mdru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdru *MeetingDateRoomUpdate) ExecX(ctx context.Context) {
	if err := mdru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mdru *MeetingDateRoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(meetingdateroom.Table, meetingdateroom.Columns, sqlgraph.NewFieldSpec(meetingdateroom.FieldID, field.TypeInt))
	if ps := mdru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mdru.mutation.Date(); ok {
		_spec.SetField(meetingdateroom.FieldDate, field.TypeInt, value)
	}
	if value, ok := mdru.mutation.AddedDate(); ok {
		_spec.AddField(meetingdateroom.FieldDate, field.TypeInt, value)
	}
	if value, ok := mdru.mutation.StartTime(); ok {
		_spec.SetField(meetingdateroom.FieldStartTime, field.TypeInt, value)
	}
	if value, ok := mdru.mutation.AddedStartTime(); ok {
		_spec.AddField(meetingdateroom.FieldStartTime, field.TypeInt, value)
	}
	if value, ok := mdru.mutation.EndTime(); ok {
		_spec.SetField(meetingdateroom.FieldEndTime, field.TypeInt, value)
	}
	if value, ok := mdru.mutation.AddedEndTime(); ok {
		_spec.AddField(meetingdateroom.FieldEndTime, field.TypeInt, value)
	}
	if mdru.mutation.MeetingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.MeetingTable,
			Columns: []string{meetingdateroom.MeetingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mdru.mutation.MeetingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.MeetingTable,
			Columns: []string{meetingdateroom.MeetingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mdru.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.RoomTable,
			Columns: []string{meetingdateroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mdru.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.RoomTable,
			Columns: []string{meetingdateroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mdru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meetingdateroom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mdru.mutation.done = true
	return n, nil
}

// MeetingDateRoomUpdateOne is the builder for updating a single MeetingDateRoom entity.
type MeetingDateRoomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MeetingDateRoomMutation
}

// SetDate sets the "date" field.
func (mdruo *MeetingDateRoomUpdateOne) SetDate(i int) *MeetingDateRoomUpdateOne {
	mdruo.mutation.ResetDate()
	mdruo.mutation.SetDate(i)
	return mdruo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (mdruo *MeetingDateRoomUpdateOne) SetNillableDate(i *int) *MeetingDateRoomUpdateOne {
	if i != nil {
		mdruo.SetDate(*i)
	}
	return mdruo
}

// AddDate adds i to the "date" field.
func (mdruo *MeetingDateRoomUpdateOne) AddDate(i int) *MeetingDateRoomUpdateOne {
	mdruo.mutation.AddDate(i)
	return mdruo
}

// SetStartTime sets the "start_time" field.
func (mdruo *MeetingDateRoomUpdateOne) SetStartTime(i int) *MeetingDateRoomUpdateOne {
	mdruo.mutation.ResetStartTime()
	mdruo.mutation.SetStartTime(i)
	return mdruo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (mdruo *MeetingDateRoomUpdateOne) SetNillableStartTime(i *int) *MeetingDateRoomUpdateOne {
	if i != nil {
		mdruo.SetStartTime(*i)
	}
	return mdruo
}

// AddStartTime adds i to the "start_time" field.
func (mdruo *MeetingDateRoomUpdateOne) AddStartTime(i int) *MeetingDateRoomUpdateOne {
	mdruo.mutation.AddStartTime(i)
	return mdruo
}

// SetEndTime sets the "end_time" field.
func (mdruo *MeetingDateRoomUpdateOne) SetEndTime(i int) *MeetingDateRoomUpdateOne {
	mdruo.mutation.ResetEndTime()
	mdruo.mutation.SetEndTime(i)
	return mdruo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (mdruo *MeetingDateRoomUpdateOne) SetNillableEndTime(i *int) *MeetingDateRoomUpdateOne {
	if i != nil {
		mdruo.SetEndTime(*i)
	}
	return mdruo
}

// AddEndTime adds i to the "end_time" field.
func (mdruo *MeetingDateRoomUpdateOne) AddEndTime(i int) *MeetingDateRoomUpdateOne {
	mdruo.mutation.AddEndTime(i)
	return mdruo
}

// SetMeetingID sets the "meeting" edge to the Meeting entity by ID.
func (mdruo *MeetingDateRoomUpdateOne) SetMeetingID(id int64) *MeetingDateRoomUpdateOne {
	mdruo.mutation.SetMeetingID(id)
	return mdruo
}

// SetNillableMeetingID sets the "meeting" edge to the Meeting entity by ID if the given value is not nil.
func (mdruo *MeetingDateRoomUpdateOne) SetNillableMeetingID(id *int64) *MeetingDateRoomUpdateOne {
	if id != nil {
		mdruo = mdruo.SetMeetingID(*id)
	}
	return mdruo
}

// SetMeeting sets the "meeting" edge to the Meeting entity.
func (mdruo *MeetingDateRoomUpdateOne) SetMeeting(m *Meeting) *MeetingDateRoomUpdateOne {
	return mdruo.SetMeetingID(m.ID)
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (mdruo *MeetingDateRoomUpdateOne) SetRoomID(id int) *MeetingDateRoomUpdateOne {
	mdruo.mutation.SetRoomID(id)
	return mdruo
}

// SetNillableRoomID sets the "room" edge to the Room entity by ID if the given value is not nil.
func (mdruo *MeetingDateRoomUpdateOne) SetNillableRoomID(id *int) *MeetingDateRoomUpdateOne {
	if id != nil {
		mdruo = mdruo.SetRoomID(*id)
	}
	return mdruo
}

// SetRoom sets the "room" edge to the Room entity.
func (mdruo *MeetingDateRoomUpdateOne) SetRoom(r *Room) *MeetingDateRoomUpdateOne {
	return mdruo.SetRoomID(r.ID)
}

// Mutation returns the MeetingDateRoomMutation object of the builder.
func (mdruo *MeetingDateRoomUpdateOne) Mutation() *MeetingDateRoomMutation {
	return mdruo.mutation
}

// ClearMeeting clears the "meeting" edge to the Meeting entity.
func (mdruo *MeetingDateRoomUpdateOne) ClearMeeting() *MeetingDateRoomUpdateOne {
	mdruo.mutation.ClearMeeting()
	return mdruo
}

// ClearRoom clears the "room" edge to the Room entity.
func (mdruo *MeetingDateRoomUpdateOne) ClearRoom() *MeetingDateRoomUpdateOne {
	mdruo.mutation.ClearRoom()
	return mdruo
}

// Where appends a list predicates to the MeetingDateRoomUpdate builder.
func (mdruo *MeetingDateRoomUpdateOne) Where(ps ...predicate.MeetingDateRoom) *MeetingDateRoomUpdateOne {
	mdruo.mutation.Where(ps...)
	return mdruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mdruo *MeetingDateRoomUpdateOne) Select(field string, fields ...string) *MeetingDateRoomUpdateOne {
	mdruo.fields = append([]string{field}, fields...)
	return mdruo
}

// Save executes the query and returns the updated MeetingDateRoom entity.
func (mdruo *MeetingDateRoomUpdateOne) Save(ctx context.Context) (*MeetingDateRoom, error) {
	return withHooks(ctx, mdruo.sqlSave, mdruo.mutation, mdruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mdruo *MeetingDateRoomUpdateOne) SaveX(ctx context.Context) *MeetingDateRoom {
	node, err := mdruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mdruo *MeetingDateRoomUpdateOne) Exec(ctx context.Context) error {
	_, err := mdruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdruo *MeetingDateRoomUpdateOne) ExecX(ctx context.Context) {
	if err := mdruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mdruo *MeetingDateRoomUpdateOne) sqlSave(ctx context.Context) (_node *MeetingDateRoom, err error) {
	_spec := sqlgraph.NewUpdateSpec(meetingdateroom.Table, meetingdateroom.Columns, sqlgraph.NewFieldSpec(meetingdateroom.FieldID, field.TypeInt))
	id, ok := mdruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MeetingDateRoom.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mdruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meetingdateroom.FieldID)
		for _, f := range fields {
			if !meetingdateroom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != meetingdateroom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mdruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mdruo.mutation.Date(); ok {
		_spec.SetField(meetingdateroom.FieldDate, field.TypeInt, value)
	}
	if value, ok := mdruo.mutation.AddedDate(); ok {
		_spec.AddField(meetingdateroom.FieldDate, field.TypeInt, value)
	}
	if value, ok := mdruo.mutation.StartTime(); ok {
		_spec.SetField(meetingdateroom.FieldStartTime, field.TypeInt, value)
	}
	if value, ok := mdruo.mutation.AddedStartTime(); ok {
		_spec.AddField(meetingdateroom.FieldStartTime, field.TypeInt, value)
	}
	if value, ok := mdruo.mutation.EndTime(); ok {
		_spec.SetField(meetingdateroom.FieldEndTime, field.TypeInt, value)
	}
	if value, ok := mdruo.mutation.AddedEndTime(); ok {
		_spec.AddField(meetingdateroom.FieldEndTime, field.TypeInt, value)
	}
	if mdruo.mutation.MeetingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.MeetingTable,
			Columns: []string{meetingdateroom.MeetingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mdruo.mutation.MeetingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.MeetingTable,
			Columns: []string{meetingdateroom.MeetingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mdruo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.RoomTable,
			Columns: []string{meetingdateroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mdruo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meetingdateroom.RoomTable,
			Columns: []string{meetingdateroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MeetingDateRoom{config: mdruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mdruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meetingdateroom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mdruo.mutation.done = true
	return _node, nil
}
