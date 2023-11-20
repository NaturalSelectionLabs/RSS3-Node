// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/naturalselectionlabs/rss3-node/internal/database/ent/feed"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/shopspring/decimal"
)

// Feed is the model entity for the Feed schema.
type Feed struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Chain holds the value of the "chain" field.
	Chain string `json:"chain,omitempty"`
	// Platform holds the value of the "platform" field.
	Platform string `json:"platform,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Index holds the value of the "index" field.
	Index uint `json:"index,omitempty"`
	// TotalActions holds the value of the "total_actions" field.
	TotalActions uint `json:"total_actions,omitempty"`
	// Actions holds the value of the "actions" field.
	Actions []schema.Action `json:"actions,omitempty"`
	// FeeValue holds the value of the "fee_value" field.
	FeeValue decimal.Decimal `json:"fee_value,omitempty"`
	// FeeToken holds the value of the "fee_token" field.
	FeeToken string `json:"fee_token,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Feed) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case feed.FieldActions:
			values[i] = new([]byte)
		case feed.FieldFeeValue:
			values[i] = new(decimal.Decimal)
		case feed.FieldIndex, feed.FieldTotalActions:
			values[i] = new(sql.NullInt64)
		case feed.FieldID, feed.FieldChain, feed.FieldPlatform, feed.FieldFrom, feed.FieldTo, feed.FieldTag, feed.FieldType, feed.FieldStatus, feed.FieldFeeToken:
			values[i] = new(sql.NullString)
		case feed.FieldTimestamp, feed.FieldCreatedAt, feed.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Feed fields.
func (f *Feed) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feed.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case feed.FieldChain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain", values[i])
			} else if value.Valid {
				f.Chain = value.String
			}
		case feed.FieldPlatform:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform", values[i])
			} else if value.Valid {
				f.Platform = value.String
			}
		case feed.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				f.From = value.String
			}
		case feed.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				f.To = value.String
			}
		case feed.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				f.Tag = value.String
			}
		case feed.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				f.Type = value.String
			}
		case feed.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = value.String
			}
		case feed.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				f.Index = uint(value.Int64)
			}
		case feed.FieldTotalActions:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_actions", values[i])
			} else if value.Valid {
				f.TotalActions = uint(value.Int64)
			}
		case feed.FieldActions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field actions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &f.Actions); err != nil {
					return fmt.Errorf("unmarshal field actions: %w", err)
				}
			}
		case feed.FieldFeeValue:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field fee_value", values[i])
			} else if value != nil {
				f.FeeValue = *value
			}
		case feed.FieldFeeToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fee_token", values[i])
			} else if value.Valid {
				f.FeeToken = value.String
			}
		case feed.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				f.Timestamp = value.Time
			}
		case feed.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case feed.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Feed.
// This includes values selected through modifiers, order, etc.
func (f *Feed) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this Feed.
// Note that you need to call Feed.Unwrap() before calling this method if this Feed
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Feed) Update() *FeedUpdateOne {
	return NewFeedClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Feed entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Feed) Unwrap() *Feed {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Feed is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Feed) String() string {
	var builder strings.Builder
	builder.WriteString("Feed(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("chain=")
	builder.WriteString(f.Chain)
	builder.WriteString(", ")
	builder.WriteString("platform=")
	builder.WriteString(f.Platform)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(f.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(f.To)
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(f.Tag)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(f.Type)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(f.Status)
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", f.Index))
	builder.WriteString(", ")
	builder.WriteString("total_actions=")
	builder.WriteString(fmt.Sprintf("%v", f.TotalActions))
	builder.WriteString(", ")
	builder.WriteString("actions=")
	builder.WriteString(fmt.Sprintf("%v", f.Actions))
	builder.WriteString(", ")
	builder.WriteString("fee_value=")
	builder.WriteString(fmt.Sprintf("%v", f.FeeValue))
	builder.WriteString(", ")
	builder.WriteString("fee_token=")
	builder.WriteString(f.FeeToken)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(f.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Feeds is a parsable slice of Feed.
type Feeds []*Feed
