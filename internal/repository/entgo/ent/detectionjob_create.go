// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjob"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjobinstance"
)

// DetectionJobCreate is the builder for creating a DetectionJob entity.
type DetectionJobCreate struct {
	config
	mutation *DetectionJobMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (djc *DetectionJobCreate) SetCreateTime(t time.Time) *DetectionJobCreate {
	djc.mutation.SetCreateTime(t)
	return djc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (djc *DetectionJobCreate) SetNillableCreateTime(t *time.Time) *DetectionJobCreate {
	if t != nil {
		djc.SetCreateTime(*t)
	}
	return djc
}

// SetUpdateTime sets the "update_time" field.
func (djc *DetectionJobCreate) SetUpdateTime(t time.Time) *DetectionJobCreate {
	djc.mutation.SetUpdateTime(t)
	return djc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (djc *DetectionJobCreate) SetNillableUpdateTime(t *time.Time) *DetectionJobCreate {
	if t != nil {
		djc.SetUpdateTime(*t)
	}
	return djc
}

// SetSchedule sets the "schedule" field.
func (djc *DetectionJobCreate) SetSchedule(s string) *DetectionJobCreate {
	djc.mutation.SetSchedule(s)
	return djc
}

// SetNillableSchedule sets the "schedule" field if the given value is not nil.
func (djc *DetectionJobCreate) SetNillableSchedule(s *string) *DetectionJobCreate {
	if s != nil {
		djc.SetSchedule(*s)
	}
	return djc
}

// SetMethod sets the "method" field.
func (djc *DetectionJobCreate) SetMethod(s string) *DetectionJobCreate {
	djc.mutation.SetMethod(s)
	return djc
}

// SetSiteID sets the "site_id" field.
func (djc *DetectionJobCreate) SetSiteID(s string) *DetectionJobCreate {
	djc.mutation.SetSiteID(s)
	return djc
}

// SetMetric sets the "metric" field.
func (djc *DetectionJobCreate) SetMetric(s string) *DetectionJobCreate {
	djc.mutation.SetMetric(s)
	return djc
}

// SetAttribute sets the "attribute" field.
func (djc *DetectionJobCreate) SetAttribute(s string) *DetectionJobCreate {
	djc.mutation.SetAttribute(s)
	return djc
}

// SetDescription sets the "description" field.
func (djc *DetectionJobCreate) SetDescription(s string) *DetectionJobCreate {
	djc.mutation.SetDescription(s)
	return djc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (djc *DetectionJobCreate) SetNillableDescription(s *string) *DetectionJobCreate {
	if s != nil {
		djc.SetDescription(*s)
	}
	return djc
}

// AddInstanceIDs adds the "instance" edge to the DetectionJobInstance entity by IDs.
func (djc *DetectionJobCreate) AddInstanceIDs(ids ...int) *DetectionJobCreate {
	djc.mutation.AddInstanceIDs(ids...)
	return djc
}

// AddInstance adds the "instance" edges to the DetectionJobInstance entity.
func (djc *DetectionJobCreate) AddInstance(d ...*DetectionJobInstance) *DetectionJobCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return djc.AddInstanceIDs(ids...)
}

// Mutation returns the DetectionJobMutation object of the builder.
func (djc *DetectionJobCreate) Mutation() *DetectionJobMutation {
	return djc.mutation
}

// Save creates the DetectionJob in the database.
func (djc *DetectionJobCreate) Save(ctx context.Context) (*DetectionJob, error) {
	var (
		err  error
		node *DetectionJob
	)
	djc.defaults()
	if len(djc.hooks) == 0 {
		if err = djc.check(); err != nil {
			return nil, err
		}
		node, err = djc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetectionJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = djc.check(); err != nil {
				return nil, err
			}
			djc.mutation = mutation
			node, err = djc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(djc.hooks) - 1; i >= 0; i-- {
			mut = djc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, djc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (djc *DetectionJobCreate) SaveX(ctx context.Context) *DetectionJob {
	v, err := djc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (djc *DetectionJobCreate) defaults() {
	if _, ok := djc.mutation.CreateTime(); !ok {
		v := detectionjob.DefaultCreateTime()
		djc.mutation.SetCreateTime(v)
	}
	if _, ok := djc.mutation.UpdateTime(); !ok {
		v := detectionjob.DefaultUpdateTime()
		djc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (djc *DetectionJobCreate) check() error {
	if _, ok := djc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := djc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := djc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New("ent: missing required field \"method\"")}
	}
	if v, ok := djc.mutation.Method(); ok {
		if err := detectionjob.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf("ent: validator failed for field \"method\": %w", err)}
		}
	}
	if _, ok := djc.mutation.SiteID(); !ok {
		return &ValidationError{Name: "site_id", err: errors.New("ent: missing required field \"site_id\"")}
	}
	if v, ok := djc.mutation.SiteID(); ok {
		if err := detectionjob.SiteIDValidator(v); err != nil {
			return &ValidationError{Name: "site_id", err: fmt.Errorf("ent: validator failed for field \"site_id\": %w", err)}
		}
	}
	if _, ok := djc.mutation.Metric(); !ok {
		return &ValidationError{Name: "metric", err: errors.New("ent: missing required field \"metric\"")}
	}
	if v, ok := djc.mutation.Metric(); ok {
		if err := detectionjob.MetricValidator(v); err != nil {
			return &ValidationError{Name: "metric", err: fmt.Errorf("ent: validator failed for field \"metric\": %w", err)}
		}
	}
	if _, ok := djc.mutation.Attribute(); !ok {
		return &ValidationError{Name: "attribute", err: errors.New("ent: missing required field \"attribute\"")}
	}
	if v, ok := djc.mutation.Attribute(); ok {
		if err := detectionjob.AttributeValidator(v); err != nil {
			return &ValidationError{Name: "attribute", err: fmt.Errorf("ent: validator failed for field \"attribute\": %w", err)}
		}
	}
	return nil
}

func (djc *DetectionJobCreate) sqlSave(ctx context.Context) (*DetectionJob, error) {
	_node, _spec := djc.createSpec()
	if err := sqlgraph.CreateNode(ctx, djc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (djc *DetectionJobCreate) createSpec() (*DetectionJob, *sqlgraph.CreateSpec) {
	var (
		_node = &DetectionJob{config: djc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: detectionjob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: detectionjob.FieldID,
			},
		}
	)
	if value, ok := djc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: detectionjob.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := djc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: detectionjob.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := djc.mutation.Schedule(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detectionjob.FieldSchedule,
		})
		_node.Schedule = &value
	}
	if value, ok := djc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detectionjob.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := djc.mutation.SiteID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detectionjob.FieldSiteID,
		})
		_node.SiteID = value
	}
	if value, ok := djc.mutation.Metric(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detectionjob.FieldMetric,
		})
		_node.Metric = value
	}
	if value, ok := djc.mutation.Attribute(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detectionjob.FieldAttribute,
		})
		_node.Attribute = value
	}
	if value, ok := djc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detectionjob.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := djc.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   detectionjob.InstanceTable,
			Columns: []string{detectionjob.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: detectionjobinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DetectionJobCreateBulk is the builder for creating many DetectionJob entities in bulk.
type DetectionJobCreateBulk struct {
	config
	builders []*DetectionJobCreate
}

// Save creates the DetectionJob entities in the database.
func (djcb *DetectionJobCreateBulk) Save(ctx context.Context) ([]*DetectionJob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(djcb.builders))
	nodes := make([]*DetectionJob, len(djcb.builders))
	mutators := make([]Mutator, len(djcb.builders))
	for i := range djcb.builders {
		func(i int, root context.Context) {
			builder := djcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DetectionJobMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, djcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, djcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, djcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (djcb *DetectionJobCreateBulk) SaveX(ctx context.Context) []*DetectionJob {
	v, err := djcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
