package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Default("").Comment("任务标题"),
		field.String("description").NotEmpty().Default("").Comment("任务描述"),
		field.Int("employee_id").Default(0).Comment("员工id"),
		field.Int("hospital_id").Default(0).Comment("医院id"),
		field.Int8("status").Default(0).Comment("任务状态"),
		field.Int8("priority").Default(0).Comment("任务优先级"),
		field.Time("created_at").Default(time.Now()).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now()).Comment("更新时间"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
