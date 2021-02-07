// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjob"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjobinstance"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/predicate"
)

// DetectionJobQuery is the builder for querying DetectionJob entities.
type DetectionJobQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.DetectionJob
	// eager-loading edges.
	withInstance *DetectionJobInstanceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DetectionJobQuery builder.
func (djq *DetectionJobQuery) Where(ps ...predicate.DetectionJob) *DetectionJobQuery {
	djq.predicates = append(djq.predicates, ps...)
	return djq
}

// Limit adds a limit step to the query.
func (djq *DetectionJobQuery) Limit(limit int) *DetectionJobQuery {
	djq.limit = &limit
	return djq
}

// Offset adds an offset step to the query.
func (djq *DetectionJobQuery) Offset(offset int) *DetectionJobQuery {
	djq.offset = &offset
	return djq
}

// Order adds an order step to the query.
func (djq *DetectionJobQuery) Order(o ...OrderFunc) *DetectionJobQuery {
	djq.order = append(djq.order, o...)
	return djq
}

// QueryInstance chains the current query on the "instance" edge.
func (djq *DetectionJobQuery) QueryInstance() *DetectionJobInstanceQuery {
	query := &DetectionJobInstanceQuery{config: djq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := djq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := djq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(detectionjob.Table, detectionjob.FieldID, selector),
			sqlgraph.To(detectionjobinstance.Table, detectionjobinstance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, detectionjob.InstanceTable, detectionjob.InstanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(djq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DetectionJob entity from the query.
// Returns a *NotFoundError when no DetectionJob was found.
func (djq *DetectionJobQuery) First(ctx context.Context) (*DetectionJob, error) {
	nodes, err := djq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{detectionjob.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (djq *DetectionJobQuery) FirstX(ctx context.Context) *DetectionJob {
	node, err := djq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DetectionJob ID from the query.
// Returns a *NotFoundError when no DetectionJob ID was found.
func (djq *DetectionJobQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = djq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{detectionjob.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (djq *DetectionJobQuery) FirstIDX(ctx context.Context) int {
	id, err := djq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DetectionJob entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DetectionJob entity is not found.
// Returns a *NotFoundError when no DetectionJob entities are found.
func (djq *DetectionJobQuery) Only(ctx context.Context) (*DetectionJob, error) {
	nodes, err := djq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{detectionjob.Label}
	default:
		return nil, &NotSingularError{detectionjob.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (djq *DetectionJobQuery) OnlyX(ctx context.Context) *DetectionJob {
	node, err := djq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DetectionJob ID in the query.
// Returns a *NotSingularError when exactly one DetectionJob ID is not found.
// Returns a *NotFoundError when no entities are found.
func (djq *DetectionJobQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = djq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = &NotSingularError{detectionjob.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (djq *DetectionJobQuery) OnlyIDX(ctx context.Context) int {
	id, err := djq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DetectionJobs.
func (djq *DetectionJobQuery) All(ctx context.Context) ([]*DetectionJob, error) {
	if err := djq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return djq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (djq *DetectionJobQuery) AllX(ctx context.Context) []*DetectionJob {
	nodes, err := djq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DetectionJob IDs.
func (djq *DetectionJobQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := djq.Select(detectionjob.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (djq *DetectionJobQuery) IDsX(ctx context.Context) []int {
	ids, err := djq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (djq *DetectionJobQuery) Count(ctx context.Context) (int, error) {
	if err := djq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return djq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (djq *DetectionJobQuery) CountX(ctx context.Context) int {
	count, err := djq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (djq *DetectionJobQuery) Exist(ctx context.Context) (bool, error) {
	if err := djq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return djq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (djq *DetectionJobQuery) ExistX(ctx context.Context) bool {
	exist, err := djq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DetectionJobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (djq *DetectionJobQuery) Clone() *DetectionJobQuery {
	if djq == nil {
		return nil
	}
	return &DetectionJobQuery{
		config:       djq.config,
		limit:        djq.limit,
		offset:       djq.offset,
		order:        append([]OrderFunc{}, djq.order...),
		predicates:   append([]predicate.DetectionJob{}, djq.predicates...),
		withInstance: djq.withInstance.Clone(),
		// clone intermediate query.
		sql:  djq.sql.Clone(),
		path: djq.path,
	}
}

// WithInstance tells the query-builder to eager-load the nodes that are connected to
// the "instance" edge. The optional arguments are used to configure the query builder of the edge.
func (djq *DetectionJobQuery) WithInstance(opts ...func(*DetectionJobInstanceQuery)) *DetectionJobQuery {
	query := &DetectionJobInstanceQuery{config: djq.config}
	for _, opt := range opts {
		opt(query)
	}
	djq.withInstance = query
	return djq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DetectionJob.Query().
//		GroupBy(detectionjob.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (djq *DetectionJobQuery) GroupBy(field string, fields ...string) *DetectionJobGroupBy {
	group := &DetectionJobGroupBy{config: djq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := djq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return djq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.DetectionJob.Query().
//		Select(detectionjob.FieldCreateTime).
//		Scan(ctx, &v)
//
func (djq *DetectionJobQuery) Select(field string, fields ...string) *DetectionJobSelect {
	djq.fields = append([]string{field}, fields...)
	return &DetectionJobSelect{DetectionJobQuery: djq}
}

func (djq *DetectionJobQuery) prepareQuery(ctx context.Context) error {
	for _, f := range djq.fields {
		if !detectionjob.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if djq.path != nil {
		prev, err := djq.path(ctx)
		if err != nil {
			return err
		}
		djq.sql = prev
	}
	return nil
}

func (djq *DetectionJobQuery) sqlAll(ctx context.Context) ([]*DetectionJob, error) {
	var (
		nodes       = []*DetectionJob{}
		_spec       = djq.querySpec()
		loadedTypes = [1]bool{
			djq.withInstance != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DetectionJob{config: djq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, djq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := djq.withInstance; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*DetectionJob)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Instance = []*DetectionJobInstance{}
		}
		query.withFKs = true
		query.Where(predicate.DetectionJobInstance(func(s *sql.Selector) {
			s.Where(sql.InValues(detectionjob.InstanceColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.detection_job_instance
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "detection_job_instance" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "detection_job_instance" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Instance = append(node.Edges.Instance, n)
		}
	}

	return nodes, nil
}

func (djq *DetectionJobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := djq.querySpec()
	return sqlgraph.CountNodes(ctx, djq.driver, _spec)
}

func (djq *DetectionJobQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := djq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (djq *DetectionJobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   detectionjob.Table,
			Columns: detectionjob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: detectionjob.FieldID,
			},
		},
		From:   djq.sql,
		Unique: true,
	}
	if fields := djq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, detectionjob.FieldID)
		for i := range fields {
			if fields[i] != detectionjob.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := djq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := djq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := djq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := djq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, detectionjob.ValidColumn)
			}
		}
	}
	return _spec
}

func (djq *DetectionJobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(djq.driver.Dialect())
	t1 := builder.Table(detectionjob.Table)
	selector := builder.Select(t1.Columns(detectionjob.Columns...)...).From(t1)
	if djq.sql != nil {
		selector = djq.sql
		selector.Select(selector.Columns(detectionjob.Columns...)...)
	}
	for _, p := range djq.predicates {
		p(selector)
	}
	for _, p := range djq.order {
		p(selector, detectionjob.ValidColumn)
	}
	if offset := djq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := djq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DetectionJobGroupBy is the group-by builder for DetectionJob entities.
type DetectionJobGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (djgb *DetectionJobGroupBy) Aggregate(fns ...AggregateFunc) *DetectionJobGroupBy {
	djgb.fns = append(djgb.fns, fns...)
	return djgb
}

// Scan applies the group-by query and scans the result into the given value.
func (djgb *DetectionJobGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := djgb.path(ctx)
	if err != nil {
		return err
	}
	djgb.sql = query
	return djgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := djgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(djgb.fields) > 1 {
		return nil, errors.New("ent: DetectionJobGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := djgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) StringsX(ctx context.Context) []string {
	v, err := djgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = djgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) StringX(ctx context.Context) string {
	v, err := djgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(djgb.fields) > 1 {
		return nil, errors.New("ent: DetectionJobGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := djgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) IntsX(ctx context.Context) []int {
	v, err := djgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = djgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) IntX(ctx context.Context) int {
	v, err := djgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(djgb.fields) > 1 {
		return nil, errors.New("ent: DetectionJobGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := djgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := djgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = djgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) Float64X(ctx context.Context) float64 {
	v, err := djgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(djgb.fields) > 1 {
		return nil, errors.New("ent: DetectionJobGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := djgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := djgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (djgb *DetectionJobGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = djgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (djgb *DetectionJobGroupBy) BoolX(ctx context.Context) bool {
	v, err := djgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (djgb *DetectionJobGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range djgb.fields {
		if !detectionjob.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := djgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := djgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (djgb *DetectionJobGroupBy) sqlQuery() *sql.Selector {
	selector := djgb.sql
	columns := make([]string, 0, len(djgb.fields)+len(djgb.fns))
	columns = append(columns, djgb.fields...)
	for _, fn := range djgb.fns {
		columns = append(columns, fn(selector, detectionjob.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(djgb.fields...)
}

// DetectionJobSelect is the builder for selecting fields of DetectionJob entities.
type DetectionJobSelect struct {
	*DetectionJobQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (djs *DetectionJobSelect) Scan(ctx context.Context, v interface{}) error {
	if err := djs.prepareQuery(ctx); err != nil {
		return err
	}
	djs.sql = djs.DetectionJobQuery.sqlQuery(ctx)
	return djs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (djs *DetectionJobSelect) ScanX(ctx context.Context, v interface{}) {
	if err := djs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) Strings(ctx context.Context) ([]string, error) {
	if len(djs.fields) > 1 {
		return nil, errors.New("ent: DetectionJobSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := djs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (djs *DetectionJobSelect) StringsX(ctx context.Context) []string {
	v, err := djs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = djs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (djs *DetectionJobSelect) StringX(ctx context.Context) string {
	v, err := djs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) Ints(ctx context.Context) ([]int, error) {
	if len(djs.fields) > 1 {
		return nil, errors.New("ent: DetectionJobSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := djs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (djs *DetectionJobSelect) IntsX(ctx context.Context) []int {
	v, err := djs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = djs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (djs *DetectionJobSelect) IntX(ctx context.Context) int {
	v, err := djs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(djs.fields) > 1 {
		return nil, errors.New("ent: DetectionJobSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := djs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (djs *DetectionJobSelect) Float64sX(ctx context.Context) []float64 {
	v, err := djs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = djs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (djs *DetectionJobSelect) Float64X(ctx context.Context) float64 {
	v, err := djs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(djs.fields) > 1 {
		return nil, errors.New("ent: DetectionJobSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := djs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (djs *DetectionJobSelect) BoolsX(ctx context.Context) []bool {
	v, err := djs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (djs *DetectionJobSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = djs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{detectionjob.Label}
	default:
		err = fmt.Errorf("ent: DetectionJobSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (djs *DetectionJobSelect) BoolX(ctx context.Context) bool {
	v, err := djs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (djs *DetectionJobSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := djs.sqlQuery().Query()
	if err := djs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (djs *DetectionJobSelect) sqlQuery() sql.Querier {
	selector := djs.sql
	selector.Select(selector.Columns(djs.fields...)...)
	return selector
}
