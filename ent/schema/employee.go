package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Default("").Comment("员工名称"),
		field.Int("hospital_id").Comment("医院id"),
		field.Time("created_at").Default(time.Now()).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now()).Comment("更新时间"),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}
