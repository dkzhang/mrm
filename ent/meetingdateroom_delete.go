// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"mrm/ent/meetingdateroom"
	"mrm/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MeetingDateRoomDelete is the builder for deleting a MeetingDateRoom entity.
type MeetingDateRoomDelete struct {
	config
	hooks    []Hook
	mutation *MeetingDateRoomMutation
}

// Where appends a list predicates to the MeetingDateRoomDelete builder.
func (mdrd *MeetingDateRoomDelete) Where(ps ...predicate.MeetingDateRoom) *MeetingDateRoomDelete {
	mdrd.mutation.Where(ps...)
	return mdrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mdrd *MeetingDateRoomDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mdrd.sqlExec, mdrd.mutation, mdrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mdrd *MeetingDateRoomDelete) ExecX(ctx context.Context) int {
	n, err := mdrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mdrd *MeetingDateRoomDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(meetingdateroom.Table, sqlgraph.NewFieldSpec(meetingdateroom.FieldID, field.TypeInt))
	if ps := mdrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mdrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mdrd.mutation.done = true
	return affected, err
}

// MeetingDateRoomDeleteOne is the builder for deleting a single MeetingDateRoom entity.
type MeetingDateRoomDeleteOne struct {
	mdrd *MeetingDateRoomDelete
}

// Where appends a list predicates to the MeetingDateRoomDelete builder.
func (mdrdo *MeetingDateRoomDeleteOne) Where(ps ...predicate.MeetingDateRoom) *MeetingDateRoomDeleteOne {
	mdrdo.mdrd.mutation.Where(ps...)
	return mdrdo
}

// Exec executes the deletion query.
func (mdrdo *MeetingDateRoomDeleteOne) Exec(ctx context.Context) error {
	n, err := mdrdo.mdrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{meetingdateroom.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdrdo *MeetingDateRoomDeleteOne) ExecX(ctx context.Context) {
	if err := mdrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
