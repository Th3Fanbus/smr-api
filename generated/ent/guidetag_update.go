// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/satisfactorymodding/smr-api/generated/ent/guide"
	"github.com/satisfactorymodding/smr-api/generated/ent/guidetag"
	"github.com/satisfactorymodding/smr-api/generated/ent/predicate"
	"github.com/satisfactorymodding/smr-api/generated/ent/tag"
)

// GuideTagUpdate is the builder for updating GuideTag entities.
type GuideTagUpdate struct {
	config
	hooks     []Hook
	mutation  *GuideTagMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GuideTagUpdate builder.
func (gtu *GuideTagUpdate) Where(ps ...predicate.GuideTag) *GuideTagUpdate {
	gtu.mutation.Where(ps...)
	return gtu
}

// SetGuideID sets the "guide_id" field.
func (gtu *GuideTagUpdate) SetGuideID(s string) *GuideTagUpdate {
	gtu.mutation.SetGuideID(s)
	return gtu
}

// SetTagID sets the "tag_id" field.
func (gtu *GuideTagUpdate) SetTagID(s string) *GuideTagUpdate {
	gtu.mutation.SetTagID(s)
	return gtu
}

// SetGuide sets the "guide" edge to the Guide entity.
func (gtu *GuideTagUpdate) SetGuide(g *Guide) *GuideTagUpdate {
	return gtu.SetGuideID(g.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (gtu *GuideTagUpdate) SetTag(t *Tag) *GuideTagUpdate {
	return gtu.SetTagID(t.ID)
}

// Mutation returns the GuideTagMutation object of the builder.
func (gtu *GuideTagUpdate) Mutation() *GuideTagMutation {
	return gtu.mutation
}

// ClearGuide clears the "guide" edge to the Guide entity.
func (gtu *GuideTagUpdate) ClearGuide() *GuideTagUpdate {
	gtu.mutation.ClearGuide()
	return gtu
}

// ClearTag clears the "tag" edge to the Tag entity.
func (gtu *GuideTagUpdate) ClearTag() *GuideTagUpdate {
	gtu.mutation.ClearTag()
	return gtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gtu *GuideTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gtu.sqlSave, gtu.mutation, gtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gtu *GuideTagUpdate) SaveX(ctx context.Context) int {
	affected, err := gtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gtu *GuideTagUpdate) Exec(ctx context.Context) error {
	_, err := gtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtu *GuideTagUpdate) ExecX(ctx context.Context) {
	if err := gtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtu *GuideTagUpdate) check() error {
	if _, ok := gtu.mutation.GuideID(); gtu.mutation.GuideCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GuideTag.guide"`)
	}
	if _, ok := gtu.mutation.TagID(); gtu.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GuideTag.tag"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gtu *GuideTagUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GuideTagUpdate {
	gtu.modifiers = append(gtu.modifiers, modifiers...)
	return gtu
}

func (gtu *GuideTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(guidetag.Table, guidetag.Columns, sqlgraph.NewFieldSpec(guidetag.FieldGuideID, field.TypeString), sqlgraph.NewFieldSpec(guidetag.FieldTagID, field.TypeString))
	if ps := gtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gtu.mutation.GuideCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.GuideTable,
			Columns: []string{guidetag.GuideColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(guide.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtu.mutation.GuideIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.GuideTable,
			Columns: []string{guidetag.GuideColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(guide.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gtu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.TagTable,
			Columns: []string{guidetag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.TagTable,
			Columns: []string{guidetag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(gtu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, gtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guidetag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gtu.mutation.done = true
	return n, nil
}

// GuideTagUpdateOne is the builder for updating a single GuideTag entity.
type GuideTagUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GuideTagMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetGuideID sets the "guide_id" field.
func (gtuo *GuideTagUpdateOne) SetGuideID(s string) *GuideTagUpdateOne {
	gtuo.mutation.SetGuideID(s)
	return gtuo
}

// SetTagID sets the "tag_id" field.
func (gtuo *GuideTagUpdateOne) SetTagID(s string) *GuideTagUpdateOne {
	gtuo.mutation.SetTagID(s)
	return gtuo
}

// SetGuide sets the "guide" edge to the Guide entity.
func (gtuo *GuideTagUpdateOne) SetGuide(g *Guide) *GuideTagUpdateOne {
	return gtuo.SetGuideID(g.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (gtuo *GuideTagUpdateOne) SetTag(t *Tag) *GuideTagUpdateOne {
	return gtuo.SetTagID(t.ID)
}

// Mutation returns the GuideTagMutation object of the builder.
func (gtuo *GuideTagUpdateOne) Mutation() *GuideTagMutation {
	return gtuo.mutation
}

// ClearGuide clears the "guide" edge to the Guide entity.
func (gtuo *GuideTagUpdateOne) ClearGuide() *GuideTagUpdateOne {
	gtuo.mutation.ClearGuide()
	return gtuo
}

// ClearTag clears the "tag" edge to the Tag entity.
func (gtuo *GuideTagUpdateOne) ClearTag() *GuideTagUpdateOne {
	gtuo.mutation.ClearTag()
	return gtuo
}

// Where appends a list predicates to the GuideTagUpdate builder.
func (gtuo *GuideTagUpdateOne) Where(ps ...predicate.GuideTag) *GuideTagUpdateOne {
	gtuo.mutation.Where(ps...)
	return gtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gtuo *GuideTagUpdateOne) Select(field string, fields ...string) *GuideTagUpdateOne {
	gtuo.fields = append([]string{field}, fields...)
	return gtuo
}

// Save executes the query and returns the updated GuideTag entity.
func (gtuo *GuideTagUpdateOne) Save(ctx context.Context) (*GuideTag, error) {
	return withHooks(ctx, gtuo.sqlSave, gtuo.mutation, gtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gtuo *GuideTagUpdateOne) SaveX(ctx context.Context) *GuideTag {
	node, err := gtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gtuo *GuideTagUpdateOne) Exec(ctx context.Context) error {
	_, err := gtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtuo *GuideTagUpdateOne) ExecX(ctx context.Context) {
	if err := gtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtuo *GuideTagUpdateOne) check() error {
	if _, ok := gtuo.mutation.GuideID(); gtuo.mutation.GuideCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GuideTag.guide"`)
	}
	if _, ok := gtuo.mutation.TagID(); gtuo.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GuideTag.tag"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gtuo *GuideTagUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GuideTagUpdateOne {
	gtuo.modifiers = append(gtuo.modifiers, modifiers...)
	return gtuo
}

func (gtuo *GuideTagUpdateOne) sqlSave(ctx context.Context) (_node *GuideTag, err error) {
	if err := gtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(guidetag.Table, guidetag.Columns, sqlgraph.NewFieldSpec(guidetag.FieldGuideID, field.TypeString), sqlgraph.NewFieldSpec(guidetag.FieldTagID, field.TypeString))
	if id, ok := gtuo.mutation.GuideID(); !ok {
		return nil, &ValidationError{Name: "guide_id", err: errors.New(`ent: missing "GuideTag.guide_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := gtuo.mutation.TagID(); !ok {
		return nil, &ValidationError{Name: "tag_id", err: errors.New(`ent: missing "GuideTag.tag_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := gtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !guidetag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := gtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gtuo.mutation.GuideCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.GuideTable,
			Columns: []string{guidetag.GuideColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(guide.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtuo.mutation.GuideIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.GuideTable,
			Columns: []string{guidetag.GuideColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(guide.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gtuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.TagTable,
			Columns: []string{guidetag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   guidetag.TagTable,
			Columns: []string{guidetag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(gtuo.modifiers...)
	_node = &GuideTag{config: gtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guidetag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gtuo.mutation.done = true
	return _node, nil
}
