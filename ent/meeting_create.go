// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mrm/ent/meeting"
	"mrm/ent/meetingdateroom"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MeetingCreate is the builder for creating a Meeting entity.
type MeetingCreate struct {
	config
	mutation *MeetingMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mc *MeetingCreate) SetName(s string) *MeetingCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetApplicant sets the "applicant" field.
func (mc *MeetingCreate) SetApplicant(s string) *MeetingCreate {
	mc.mutation.SetApplicant(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MeetingCreate) SetID(i int) *MeetingCreate {
	mc.mutation.SetID(i)
	return mc
}

// AddMdrIDs adds the "mdrs" edge to the MeetingDateRoom entity by IDs.
func (mc *MeetingCreate) AddMdrIDs(ids ...int) *MeetingCreate {
	mc.mutation.AddMdrIDs(ids...)
	return mc
}

// AddMdrs adds the "mdrs" edges to the MeetingDateRoom entity.
func (mc *MeetingCreate) AddMdrs(m ...*MeetingDateRoom) *MeetingCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMdrIDs(ids...)
}

// Mutation returns the MeetingMutation object of the builder.
func (mc *MeetingCreate) Mutation() *MeetingMutation {
	return mc.mutation
}

// Save creates the Meeting in the database.
func (mc *MeetingCreate) Save(ctx context.Context) (*Meeting, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MeetingCreate) SaveX(ctx context.Context) *Meeting {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MeetingCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MeetingCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MeetingCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Meeting.name"`)}
	}
	if _, ok := mc.mutation.Applicant(); !ok {
		return &ValidationError{Name: "applicant", err: errors.New(`ent: missing required field "Meeting.applicant"`)}
	}
	return nil
}

func (mc *MeetingCreate) sqlSave(ctx context.Context) (*Meeting, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MeetingCreate) createSpec() (*Meeting, *sqlgraph.CreateSpec) {
	var (
		_node = &Meeting{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(meeting.Table, sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(meeting.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Applicant(); ok {
		_spec.SetField(meeting.FieldApplicant, field.TypeString, value)
		_node.Applicant = value
	}
	if nodes := mc.mutation.MdrsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meeting.MdrsTable,
			Columns: []string{meeting.MdrsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meetingdateroom.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MeetingCreateBulk is the builder for creating many Meeting entities in bulk.
type MeetingCreateBulk struct {
	config
	err      error
	builders []*MeetingCreate
}

// Save creates the Meeting entities in the database.
func (mcb *MeetingCreateBulk) Save(ctx context.Context) ([]*Meeting, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Meeting, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MeetingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MeetingCreateBulk) SaveX(ctx context.Context) []*Meeting {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MeetingCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MeetingCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
