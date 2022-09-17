package schema

import "entgo.io/ent"

// Hospital holds the schema definition for the Hospital entity.
type Hospital struct {
	ent.Schema
}

// Fields of the Hospital.
func (Hospital) Fields() []ent.Field {
	return nil
}

// Edges of the Hospital.
func (Hospital) Edges() []ent.Edge {
	return nil
}
