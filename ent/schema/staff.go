package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").MaxLen(50),
		field.Int32("role").Positive(),
		field.String("lastName").MaxLen(20),
		field.String("firstName").MaxLen(20),
		field.Int32("gender").Positive(),
		field.String("phoneNumber").MaxLen(20).Optional(),
		field.String("description").MaxLen(500).Optional(),
		field.UUID("userId", uuid.New()),
	}
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return nil
}
