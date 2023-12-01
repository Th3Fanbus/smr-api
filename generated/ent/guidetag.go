// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/satisfactorymodding/smr-api/generated/ent/guide"
	"github.com/satisfactorymodding/smr-api/generated/ent/guidetag"
	"github.com/satisfactorymodding/smr-api/generated/ent/tag"
)

// GuideTag is the model entity for the GuideTag schema.
type GuideTag struct {
	config `json:"-"`
	// GuideID holds the value of the "guide_id" field.
	GuideID string `json:"guide_id,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID string `json:"tag_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GuideTagQuery when eager-loading is set.
	Edges        GuideTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GuideTagEdges holds the relations/edges for other nodes in the graph.
type GuideTagEdges struct {
	// Guide holds the value of the guide edge.
	Guide *Guide `json:"guide,omitempty"`
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GuideOrErr returns the Guide value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GuideTagEdges) GuideOrErr() (*Guide, error) {
	if e.loadedTypes[0] {
		if e.Guide == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: guide.Label}
		}
		return e.Guide, nil
	}
	return nil, &NotLoadedError{edge: "guide"}
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GuideTagEdges) TagOrErr() (*Tag, error) {
	if e.loadedTypes[1] {
		if e.Tag == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tag.Label}
		}
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GuideTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case guidetag.FieldGuideID, guidetag.FieldTagID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GuideTag fields.
func (gt *GuideTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case guidetag.FieldGuideID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field guide_id", values[i])
			} else if value.Valid {
				gt.GuideID = value.String
			}
		case guidetag.FieldTagID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value.Valid {
				gt.TagID = value.String
			}
		default:
			gt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GuideTag.
// This includes values selected through modifiers, order, etc.
func (gt *GuideTag) Value(name string) (ent.Value, error) {
	return gt.selectValues.Get(name)
}

// QueryGuide queries the "guide" edge of the GuideTag entity.
func (gt *GuideTag) QueryGuide() *GuideQuery {
	return NewGuideTagClient(gt.config).QueryGuide(gt)
}

// QueryTag queries the "tag" edge of the GuideTag entity.
func (gt *GuideTag) QueryTag() *TagQuery {
	return NewGuideTagClient(gt.config).QueryTag(gt)
}

// Update returns a builder for updating this GuideTag.
// Note that you need to call GuideTag.Unwrap() before calling this method if this GuideTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (gt *GuideTag) Update() *GuideTagUpdateOne {
	return NewGuideTagClient(gt.config).UpdateOne(gt)
}

// Unwrap unwraps the GuideTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gt *GuideTag) Unwrap() *GuideTag {
	_tx, ok := gt.config.driver.(*txDriver)
	if !ok {
		panic("ent: GuideTag is not a transactional entity")
	}
	gt.config.driver = _tx.drv
	return gt
}

// String implements the fmt.Stringer.
func (gt *GuideTag) String() string {
	var builder strings.Builder
	builder.WriteString("GuideTag(")
	builder.WriteString("guide_id=")
	builder.WriteString(gt.GuideID)
	builder.WriteString(", ")
	builder.WriteString("tag_id=")
	builder.WriteString(gt.TagID)
	builder.WriteByte(')')
	return builder.String()
}

// GuideTags is a parsable slice of GuideTag.
type GuideTags []*GuideTag
