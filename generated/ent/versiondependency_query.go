// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/satisfactorymodding/smr-api/generated/ent/mod"
	"github.com/satisfactorymodding/smr-api/generated/ent/predicate"
	"github.com/satisfactorymodding/smr-api/generated/ent/version"
	"github.com/satisfactorymodding/smr-api/generated/ent/versiondependency"
)

// VersionDependencyQuery is the builder for querying VersionDependency entities.
type VersionDependencyQuery struct {
	config
	ctx         *QueryContext
	order       []versiondependency.OrderOption
	inters      []Interceptor
	predicates  []predicate.VersionDependency
	withVersion *VersionQuery
	withMod     *ModQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VersionDependencyQuery builder.
func (vdq *VersionDependencyQuery) Where(ps ...predicate.VersionDependency) *VersionDependencyQuery {
	vdq.predicates = append(vdq.predicates, ps...)
	return vdq
}

// Limit the number of records to be returned by this query.
func (vdq *VersionDependencyQuery) Limit(limit int) *VersionDependencyQuery {
	vdq.ctx.Limit = &limit
	return vdq
}

// Offset to start from.
func (vdq *VersionDependencyQuery) Offset(offset int) *VersionDependencyQuery {
	vdq.ctx.Offset = &offset
	return vdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vdq *VersionDependencyQuery) Unique(unique bool) *VersionDependencyQuery {
	vdq.ctx.Unique = &unique
	return vdq
}

// Order specifies how the records should be ordered.
func (vdq *VersionDependencyQuery) Order(o ...versiondependency.OrderOption) *VersionDependencyQuery {
	vdq.order = append(vdq.order, o...)
	return vdq
}

// QueryVersion chains the current query on the "version" edge.
func (vdq *VersionDependencyQuery) QueryVersion() *VersionQuery {
	query := (&VersionClient{config: vdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(versiondependency.Table, versiondependency.VersionColumn, selector),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, versiondependency.VersionTable, versiondependency.VersionColumn),
		)
		fromU = sqlgraph.SetNeighbors(vdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMod chains the current query on the "mod" edge.
func (vdq *VersionDependencyQuery) QueryMod() *ModQuery {
	query := (&ModClient{config: vdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(versiondependency.Table, versiondependency.ModColumn, selector),
			sqlgraph.To(mod.Table, mod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, versiondependency.ModTable, versiondependency.ModColumn),
		)
		fromU = sqlgraph.SetNeighbors(vdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VersionDependency entity from the query.
// Returns a *NotFoundError when no VersionDependency was found.
func (vdq *VersionDependencyQuery) First(ctx context.Context) (*VersionDependency, error) {
	nodes, err := vdq.Limit(1).All(setContextOp(ctx, vdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{versiondependency.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vdq *VersionDependencyQuery) FirstX(ctx context.Context) *VersionDependency {
	node, err := vdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single VersionDependency entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VersionDependency entity is found.
// Returns a *NotFoundError when no VersionDependency entities are found.
func (vdq *VersionDependencyQuery) Only(ctx context.Context) (*VersionDependency, error) {
	nodes, err := vdq.Limit(2).All(setContextOp(ctx, vdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{versiondependency.Label}
	default:
		return nil, &NotSingularError{versiondependency.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vdq *VersionDependencyQuery) OnlyX(ctx context.Context) *VersionDependency {
	node, err := vdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of VersionDependencies.
func (vdq *VersionDependencyQuery) All(ctx context.Context) ([]*VersionDependency, error) {
	ctx = setContextOp(ctx, vdq.ctx, "All")
	if err := vdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VersionDependency, *VersionDependencyQuery]()
	return withInterceptors[[]*VersionDependency](ctx, vdq, qr, vdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vdq *VersionDependencyQuery) AllX(ctx context.Context) []*VersionDependency {
	nodes, err := vdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (vdq *VersionDependencyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vdq.ctx, "Count")
	if err := vdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vdq, querierCount[*VersionDependencyQuery](), vdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vdq *VersionDependencyQuery) CountX(ctx context.Context) int {
	count, err := vdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vdq *VersionDependencyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vdq.ctx, "Exist")
	switch _, err := vdq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vdq *VersionDependencyQuery) ExistX(ctx context.Context) bool {
	exist, err := vdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VersionDependencyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vdq *VersionDependencyQuery) Clone() *VersionDependencyQuery {
	if vdq == nil {
		return nil
	}
	return &VersionDependencyQuery{
		config:      vdq.config,
		ctx:         vdq.ctx.Clone(),
		order:       append([]versiondependency.OrderOption{}, vdq.order...),
		inters:      append([]Interceptor{}, vdq.inters...),
		predicates:  append([]predicate.VersionDependency{}, vdq.predicates...),
		withVersion: vdq.withVersion.Clone(),
		withMod:     vdq.withMod.Clone(),
		// clone intermediate query.
		sql:  vdq.sql.Clone(),
		path: vdq.path,
	}
}

// WithVersion tells the query-builder to eager-load the nodes that are connected to
// the "version" edge. The optional arguments are used to configure the query builder of the edge.
func (vdq *VersionDependencyQuery) WithVersion(opts ...func(*VersionQuery)) *VersionDependencyQuery {
	query := (&VersionClient{config: vdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vdq.withVersion = query
	return vdq
}

// WithMod tells the query-builder to eager-load the nodes that are connected to
// the "mod" edge. The optional arguments are used to configure the query builder of the edge.
func (vdq *VersionDependencyQuery) WithMod(opts ...func(*ModQuery)) *VersionDependencyQuery {
	query := (&ModClient{config: vdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vdq.withMod = query
	return vdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VersionDependency.Query().
//		GroupBy(versiondependency.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vdq *VersionDependencyQuery) GroupBy(field string, fields ...string) *VersionDependencyGroupBy {
	vdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VersionDependencyGroupBy{build: vdq}
	grbuild.flds = &vdq.ctx.Fields
	grbuild.label = versiondependency.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.VersionDependency.Query().
//		Select(versiondependency.FieldCreatedAt).
//		Scan(ctx, &v)
func (vdq *VersionDependencyQuery) Select(fields ...string) *VersionDependencySelect {
	vdq.ctx.Fields = append(vdq.ctx.Fields, fields...)
	sbuild := &VersionDependencySelect{VersionDependencyQuery: vdq}
	sbuild.label = versiondependency.Label
	sbuild.flds, sbuild.scan = &vdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VersionDependencySelect configured with the given aggregations.
func (vdq *VersionDependencyQuery) Aggregate(fns ...AggregateFunc) *VersionDependencySelect {
	return vdq.Select().Aggregate(fns...)
}

func (vdq *VersionDependencyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vdq); err != nil {
				return err
			}
		}
	}
	for _, f := range vdq.ctx.Fields {
		if !versiondependency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vdq.path != nil {
		prev, err := vdq.path(ctx)
		if err != nil {
			return err
		}
		vdq.sql = prev
	}
	return nil
}

func (vdq *VersionDependencyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VersionDependency, error) {
	var (
		nodes       = []*VersionDependency{}
		_spec       = vdq.querySpec()
		loadedTypes = [2]bool{
			vdq.withVersion != nil,
			vdq.withMod != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VersionDependency).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VersionDependency{config: vdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vdq.modifiers) > 0 {
		_spec.Modifiers = vdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vdq.withVersion; query != nil {
		if err := vdq.loadVersion(ctx, query, nodes, nil,
			func(n *VersionDependency, e *Version) { n.Edges.Version = e }); err != nil {
			return nil, err
		}
	}
	if query := vdq.withMod; query != nil {
		if err := vdq.loadMod(ctx, query, nodes, nil,
			func(n *VersionDependency, e *Mod) { n.Edges.Mod = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vdq *VersionDependencyQuery) loadVersion(ctx context.Context, query *VersionQuery, nodes []*VersionDependency, init func(*VersionDependency), assign func(*VersionDependency, *Version)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*VersionDependency)
	for i := range nodes {
		fk := nodes[i].VersionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(version.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "version_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vdq *VersionDependencyQuery) loadMod(ctx context.Context, query *ModQuery, nodes []*VersionDependency, init func(*VersionDependency), assign func(*VersionDependency, *Mod)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*VersionDependency)
	for i := range nodes {
		fk := nodes[i].ModID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(mod.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mod_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vdq *VersionDependencyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vdq.querySpec()
	if len(vdq.modifiers) > 0 {
		_spec.Modifiers = vdq.modifiers
	}
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, vdq.driver, _spec)
}

func (vdq *VersionDependencyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(versiondependency.Table, versiondependency.Columns, nil)
	_spec.From = vdq.sql
	if unique := vdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vdq.path != nil {
		_spec.Unique = true
	}
	if fields := vdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if vdq.withVersion != nil {
			_spec.Node.AddColumnOnce(versiondependency.FieldVersionID)
		}
		if vdq.withMod != nil {
			_spec.Node.AddColumnOnce(versiondependency.FieldModID)
		}
	}
	if ps := vdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vdq *VersionDependencyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vdq.driver.Dialect())
	t1 := builder.Table(versiondependency.Table)
	columns := vdq.ctx.Fields
	if len(columns) == 0 {
		columns = versiondependency.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vdq.sql != nil {
		selector = vdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vdq.ctx.Unique != nil && *vdq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range vdq.modifiers {
		m(selector)
	}
	for _, p := range vdq.predicates {
		p(selector)
	}
	for _, p := range vdq.order {
		p(selector)
	}
	if offset := vdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vdq *VersionDependencyQuery) Modify(modifiers ...func(s *sql.Selector)) *VersionDependencySelect {
	vdq.modifiers = append(vdq.modifiers, modifiers...)
	return vdq.Select()
}

// VersionDependencyGroupBy is the group-by builder for VersionDependency entities.
type VersionDependencyGroupBy struct {
	selector
	build *VersionDependencyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vdgb *VersionDependencyGroupBy) Aggregate(fns ...AggregateFunc) *VersionDependencyGroupBy {
	vdgb.fns = append(vdgb.fns, fns...)
	return vdgb
}

// Scan applies the selector query and scans the result into the given value.
func (vdgb *VersionDependencyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vdgb.build.ctx, "GroupBy")
	if err := vdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VersionDependencyQuery, *VersionDependencyGroupBy](ctx, vdgb.build, vdgb, vdgb.build.inters, v)
}

func (vdgb *VersionDependencyGroupBy) sqlScan(ctx context.Context, root *VersionDependencyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vdgb.fns))
	for _, fn := range vdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vdgb.flds)+len(vdgb.fns))
		for _, f := range *vdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VersionDependencySelect is the builder for selecting fields of VersionDependency entities.
type VersionDependencySelect struct {
	*VersionDependencyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vds *VersionDependencySelect) Aggregate(fns ...AggregateFunc) *VersionDependencySelect {
	vds.fns = append(vds.fns, fns...)
	return vds
}

// Scan applies the selector query and scans the result into the given value.
func (vds *VersionDependencySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vds.ctx, "Select")
	if err := vds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VersionDependencyQuery, *VersionDependencySelect](ctx, vds.VersionDependencyQuery, vds, vds.inters, v)
}

func (vds *VersionDependencySelect) sqlScan(ctx context.Context, root *VersionDependencyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vds.fns))
	for _, fn := range vds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vds *VersionDependencySelect) Modify(modifiers ...func(s *sql.Selector)) *VersionDependencySelect {
	vds.modifiers = append(vds.modifiers, modifiers...)
	return vds
}
