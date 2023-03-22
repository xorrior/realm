// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/predicate"
	"github.com/kcarretto/realm/tavern/ent/session"
	"github.com/kcarretto/realm/tavern/ent/tag"
)

// TagQuery is the builder for querying Tag entities.
type TagQuery struct {
	config
	ctx               *QueryContext
	order             []OrderFunc
	inters            []Interceptor
	predicates        []predicate.Tag
	withSessions      *SessionQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Tag) error
	withNamedSessions map[string]*SessionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TagQuery builder.
func (tq *TagQuery) Where(ps ...predicate.Tag) *TagQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TagQuery) Limit(limit int) *TagQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TagQuery) Offset(offset int) *TagQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TagQuery) Unique(unique bool) *TagQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TagQuery) Order(o ...OrderFunc) *TagQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QuerySessions chains the current query on the "sessions" edge.
func (tq *TagQuery) QuerySessions() *SessionQuery {
	query := (&SessionClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, selector),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.SessionsTable, tag.SessionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tag entity from the query.
// Returns a *NotFoundError when no Tag was found.
func (tq *TagQuery) First(ctx context.Context) (*Tag, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TagQuery) FirstX(ctx context.Context) *Tag {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tag ID from the query.
// Returns a *NotFoundError when no Tag ID was found.
func (tq *TagQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TagQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tag entity is found.
// Returns a *NotFoundError when no Tag entities are found.
func (tq *TagQuery) Only(ctx context.Context) (*Tag, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tag.Label}
	default:
		return nil, &NotSingularError{tag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TagQuery) OnlyX(ctx context.Context) *Tag {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tag ID in the query.
// Returns a *NotSingularError when more than one Tag ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TagQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tag.Label}
	default:
		err = &NotSingularError{tag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TagQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tags.
func (tq *TagQuery) All(ctx context.Context) ([]*Tag, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tag, *TagQuery]()
	return withInterceptors[[]*Tag](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TagQuery) AllX(ctx context.Context) []*Tag {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tag IDs.
func (tq *TagQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(tag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TagQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TagQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TagQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TagQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TagQuery) Clone() *TagQuery {
	if tq == nil {
		return nil
	}
	return &TagQuery{
		config:       tq.config,
		ctx:          tq.ctx.Clone(),
		order:        append([]OrderFunc{}, tq.order...),
		inters:       append([]Interceptor{}, tq.inters...),
		predicates:   append([]predicate.Tag{}, tq.predicates...),
		withSessions: tq.withSessions.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithSessions tells the query-builder to eager-load the nodes that are connected to
// the "sessions" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TagQuery) WithSessions(opts ...func(*SessionQuery)) *TagQuery {
	query := (&SessionClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withSessions = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Tag.Query().
//		GroupBy(tag.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TagQuery) GroupBy(field string, fields ...string) *TagGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TagGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = tag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Tag.Query().
//		Select(tag.FieldName).
//		Scan(ctx, &v)
func (tq *TagQuery) Select(fields ...string) *TagSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TagSelect{TagQuery: tq}
	sbuild.label = tag.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TagSelect configured with the given aggregations.
func (tq *TagQuery) Aggregate(fns ...AggregateFunc) *TagSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !tag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tag, error) {
	var (
		nodes       = []*Tag{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withSessions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tag{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withSessions; query != nil {
		if err := tq.loadSessions(ctx, query, nodes,
			func(n *Tag) { n.Edges.Sessions = []*Session{} },
			func(n *Tag, e *Session) { n.Edges.Sessions = append(n.Edges.Sessions, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedSessions {
		if err := tq.loadSessions(ctx, query, nodes,
			func(n *Tag) { n.appendNamedSessions(name) },
			func(n *Tag, e *Session) { n.appendNamedSessions(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TagQuery) loadSessions(ctx context.Context, query *SessionQuery, nodes []*Tag, init func(*Tag), assign func(*Tag, *Session)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Tag)
	nids := make(map[int]map[*Tag]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tag.SessionsTable)
		s.Join(joinT).On(s.C(session.FieldID), joinT.C(tag.SessionsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tag.SessionsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tag.SessionsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tag]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Session](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "sessions" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tq *TagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for i := range fields {
			if fields[i] != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tag.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = tag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSessions tells the query-builder to eager-load the nodes that are connected to the "sessions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TagQuery) WithNamedSessions(name string, opts ...func(*SessionQuery)) *TagQuery {
	query := (&SessionClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedSessions == nil {
		tq.withNamedSessions = make(map[string]*SessionQuery)
	}
	tq.withNamedSessions[name] = query
	return tq
}

// TagGroupBy is the group-by builder for Tag entities.
type TagGroupBy struct {
	selector
	build *TagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TagGroupBy) Aggregate(fns ...AggregateFunc) *TagGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TagQuery, *TagGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TagGroupBy) sqlScan(ctx context.Context, root *TagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TagSelect is the builder for selecting fields of Tag entities.
type TagSelect struct {
	*TagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TagSelect) Aggregate(fns ...AggregateFunc) *TagSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TagQuery, *TagSelect](ctx, ts.TagQuery, ts, ts.inters, v)
}

func (ts *TagSelect) sqlScan(ctx context.Context, root *TagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
