// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"task_manager/ent/employee"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Employee is the model entity for the Employee schema.
type Employee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 员工名称
	Name string `json:"name,omitempty"`
	// 医院id
	HospitalID int `json:"hospital_id,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employee) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case employee.FieldID, employee.FieldHospitalID:
			values[i] = new(sql.NullInt64)
		case employee.FieldName:
			values[i] = new(sql.NullString)
		case employee.FieldCreatedAt, employee.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Employee", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employee fields.
func (e *Employee) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case employee.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case employee.FieldHospitalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hospital_id", values[i])
			} else if value.Valid {
				e.HospitalID = int(value.Int64)
			}
		case employee.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case employee.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Employee.
// Note that you need to call Employee.Unwrap() before calling this method if this Employee
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employee) Update() *EmployeeUpdateOne {
	return (&EmployeeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Employee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Employee) Unwrap() *Employee {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employee is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employee) String() string {
	var builder strings.Builder
	builder.WriteString("Employee(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("hospital_id=")
	builder.WriteString(fmt.Sprintf("%v", e.HospitalID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Employees is a parsable slice of Employee.
type Employees []*Employee

func (e Employees) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
