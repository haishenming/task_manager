package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Hospital holds the schema definition for the Hospital entity.
type Hospital struct {
	ent.Schema
}

// Fields of the Hospital.
func (Hospital) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Default("").Comment("医院名称"),
		field.String("address").NotEmpty().Default("").Comment("医院地址"),
		field.Time("created_at").Default(time.Now()).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now()).Comment("更新时间"),
	}
}

// Edges of the Hospital.
func (Hospital) Edges() []ent.Edge {
	return nil
}
