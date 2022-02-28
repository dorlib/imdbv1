// Code generated by entc, DO NOT EDIT.

package ent

import (
	"IMDB/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// BirthDay holds the value of the "birthDay" field.
	BirthDay string `json:"birthDay,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[0] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldFirstname, user.FieldLastname, user.FieldNickname, user.FieldDescription, user.FieldPassword, user.FieldEmail, user.FieldBirthDay:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldFirstname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstname", values[i])
			} else if value.Valid {
				u.Firstname = value.String
			}
		case user.FieldLastname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastname", values[i])
			} else if value.Valid {
				u.Lastname = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldBirthDay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birthDay", values[i])
			} else if value.Valid {
				u.BirthDay = value.String
			}
		}
	}
	return nil
}

// QueryReviews queries the "reviews" edge of the User entity.
func (u *User) QueryReviews() *ReviewQuery {
	return (&UserClient{config: u.config}).QueryReviews(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", firstname=")
	builder.WriteString(u.Firstname)
	builder.WriteString(", lastname=")
	builder.WriteString(u.Lastname)
	builder.WriteString(", nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", description=")
	builder.WriteString(u.Description)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", birthDay=")
	builder.WriteString(u.BirthDay)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
