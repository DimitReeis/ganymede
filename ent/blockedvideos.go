// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/zibbp/ganymede/ent/blockedvideos"
)

// BlockedVideos is the model entity for the BlockedVideos schema.
type BlockedVideos struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the blocked vod.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlockedVideos) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blockedvideos.FieldID:
			values[i] = new(sql.NullString)
		case blockedvideos.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlockedVideos fields.
func (bv *BlockedVideos) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blockedvideos.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bv.ID = value.String
			}
		case blockedvideos.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bv.CreatedAt = value.Time
			}
		default:
			bv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BlockedVideos.
// This includes values selected through modifiers, order, etc.
func (bv *BlockedVideos) Value(name string) (ent.Value, error) {
	return bv.selectValues.Get(name)
}

// Update returns a builder for updating this BlockedVideos.
// Note that you need to call BlockedVideos.Unwrap() before calling this method if this BlockedVideos
// was returned from a transaction, and the transaction was committed or rolled back.
func (bv *BlockedVideos) Update() *BlockedVideosUpdateOne {
	return NewBlockedVideosClient(bv.config).UpdateOne(bv)
}

// Unwrap unwraps the BlockedVideos entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bv *BlockedVideos) Unwrap() *BlockedVideos {
	_tx, ok := bv.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlockedVideos is not a transactional entity")
	}
	bv.config.driver = _tx.drv
	return bv
}

// String implements the fmt.Stringer.
func (bv *BlockedVideos) String() string {
	var builder strings.Builder
	builder.WriteString("BlockedVideos(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bv.ID))
	builder.WriteString("created_at=")
	builder.WriteString(bv.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BlockedVideosSlice is a parsable slice of BlockedVideos.
type BlockedVideosSlice []*BlockedVideos
