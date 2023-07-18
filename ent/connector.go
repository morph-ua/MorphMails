// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"helium/ent/connector"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Connector is the model entity for the Connector schema.
type Connector struct {
	config `form:"-" json:"-" query:"-"`
	// ID of the ent.
	ID string `json:"id" query:"id" form:"id"`
	// Name holds the value of the "name" field.
	Name string `json:"name" query:"name" form:"name"`
	// URL holds the value of the "url" field.
	URL string `json:"url" query:"url" form:"url"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"secret" query:"secret" form:"secret"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConnectorQuery when eager-loading is set.
	Edges        ConnectorEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ConnectorEdges holds the relations/edges for other nodes in the graph.
type ConnectorEdges struct {
	// Receivers holds the value of the receivers edge.
	Receivers []*Receiver `json:"receivers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReceiversOrErr returns the Receivers value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectorEdges) ReceiversOrErr() ([]*Receiver, error) {
	if e.loadedTypes[0] {
		return e.Receivers, nil
	}
	return nil, &NotLoadedError{edge: "receivers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Connector) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case connector.FieldID, connector.FieldName, connector.FieldURL, connector.FieldSecret:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Connector fields.
func (c *Connector) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case connector.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case connector.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case connector.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				c.URL = value.String
			}
		case connector.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				c.Secret = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Connector.
// This includes values selected through modifiers, order, etc.
func (c *Connector) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryReceivers queries the "receivers" edge of the Connector entity.
func (c *Connector) QueryReceivers() *ReceiverQuery {
	return NewConnectorClient(c.config).QueryReceivers(c)
}

// Update returns a builder for updating this Connector.
// Note that you need to call Connector.Unwrap() before calling this method if this Connector
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Connector) Update() *ConnectorUpdateOne {
	return NewConnectorClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Connector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Connector) Unwrap() *Connector {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Connector is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Connector) String() string {
	var builder strings.Builder
	builder.WriteString("Connector(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(c.URL)
	builder.WriteString(", ")
	builder.WriteString("secret=")
	builder.WriteString(c.Secret)
	builder.WriteByte(')')
	return builder.String()
}

// Connectors is a parsable slice of Connector.
type Connectors []*Connector
