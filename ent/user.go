// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"helium/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Emails holds the value of the "emails" field.
	Emails []string `json:"emails,omitempty"`
	// Forward holds the value of the "forward" field.
	Forward bool `json:"forward,omitempty"`
	// Paid holds the value of the "paid" field.
	Paid bool `json:"paid,omitempty"`
	// Counter holds the value of the "counter" field.
	Counter int8 `json:"counter,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Receivers holds the value of the receivers edge.
	Receivers []*Receiver `json:"receivers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReceiversOrErr returns the Receivers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReceiversOrErr() ([]*Receiver, error) {
	if e.loadedTypes[0] {
		return e.Receivers, nil
	}
	return nil, &NotLoadedError{edge: "receivers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmails:
			values[i] = new([]byte)
		case user.FieldForward, user.FieldPaid:
			values[i] = new(sql.NullBool)
		case user.FieldCounter:
			values[i] = new(sql.NullInt64)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldEmails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field emails", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Emails); err != nil {
					return fmt.Errorf("unmarshal field emails: %w", err)
				}
			}
		case user.FieldForward:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field forward", values[i])
			} else if value.Valid {
				u.Forward = value.Bool
			}
		case user.FieldPaid:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field paid", values[i])
			} else if value.Valid {
				u.Paid = value.Bool
			}
		case user.FieldCounter:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field counter", values[i])
			} else if value.Valid {
				u.Counter = int8(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryReceivers queries the "receivers" edge of the User entity.
func (u *User) QueryReceivers() *ReceiverQuery {
	return NewUserClient(u.config).QueryReceivers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("emails=")
	builder.WriteString(fmt.Sprintf("%v", u.Emails))
	builder.WriteString(", ")
	builder.WriteString("forward=")
	builder.WriteString(fmt.Sprintf("%v", u.Forward))
	builder.WriteString(", ")
	builder.WriteString("paid=")
	builder.WriteString(fmt.Sprintf("%v", u.Paid))
	builder.WriteString(", ")
	builder.WriteString("counter=")
	builder.WriteString(fmt.Sprintf("%v", u.Counter))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
