// Code generated by entc, DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/api/src/infra/db/model/authentication"
	"github.com/Southclaws/storyden/api/src/infra/db/model/user"
	"github.com/google/uuid"
)

// Authentication is the model entity for the Authentication schema.
type Authentication struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Service holds the value of the "service" field.
	// The authentication service name, such as GitHub, Twitter, Discord, etc. Or, 'password' for password auth and 'api_token' for token auth
	Service string `json:"service,omitempty"`
	// Identifier holds the value of the "identifier" field.
	// The identifier, usually a user/account ID on some OAuth service or API token name. If it's a password, this is blank.
	Identifier string `json:"identifier,omitempty"`
	// Token holds the value of the "token" field.
	// The actual authentication token/password/key/etc. If OAuth, it'll be the access_token value, if it's a password, a hash and if it's an api_token type then the API token string.
	Token string `json:"-"`
	// Metadata holds the value of the "metadata" field.
	// Any necessary metadata specific to the authentication method.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthenticationQuery when eager-loading is set.
	Edges               AuthenticationEdges `json:"edges"`
	user_authentication *uuid.UUID
}

// AuthenticationEdges holds the relations/edges for other nodes in the graph.
type AuthenticationEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthenticationEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Authentication) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case authentication.FieldMetadata:
			values[i] = new([]byte)
		case authentication.FieldService, authentication.FieldIdentifier, authentication.FieldToken:
			values[i] = new(sql.NullString)
		case authentication.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case authentication.FieldID:
			values[i] = new(uuid.UUID)
		case authentication.ForeignKeys[0]: // user_authentication
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Authentication", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Authentication fields.
func (a *Authentication) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authentication.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case authentication.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				a.CreateTime = value.Time
			}
		case authentication.FieldService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service", values[i])
			} else if value.Valid {
				a.Service = value.String
			}
		case authentication.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				a.Identifier = value.String
			}
		case authentication.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				a.Token = value.String
			}
		case authentication.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case authentication.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_authentication", values[i])
			} else if value.Valid {
				a.user_authentication = new(uuid.UUID)
				*a.user_authentication = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Authentication entity.
func (a *Authentication) QueryUser() *UserQuery {
	return (&AuthenticationClient{config: a.config}).QueryUser(a)
}

// Update returns a builder for updating this Authentication.
// Note that you need to call Authentication.Unwrap() before calling this method if this Authentication
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Authentication) Update() *AuthenticationUpdateOne {
	return (&AuthenticationClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Authentication entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Authentication) Unwrap() *Authentication {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("model: Authentication is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Authentication) String() string {
	var builder strings.Builder
	builder.WriteString("Authentication(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(a.CreateTime.Format(time.ANSIC))
	builder.WriteString(", service=")
	builder.WriteString(a.Service)
	builder.WriteString(", identifier=")
	builder.WriteString(a.Identifier)
	builder.WriteString(", token=<sensitive>")
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", a.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// Authentications is a parsable slice of Authentication.
type Authentications []*Authentication

func (a Authentications) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
