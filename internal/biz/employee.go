package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "task_manager/api/task/v1"
)

var (
	// ErrEmployeeNotFound is Employee not found.
	ErrEmployeeNotFound = errors.NotFound(v1.ErrorReason_EMPLOYEE_NOT_FOUND.String(), "employee not found")
)

// Employee is a Employee model.
type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	HospitalID int       `json:"hospital_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// EmployeeRepo is a Employee repo.
type EmployeeRepo interface {
	Save(context.Context, *Employee) (*Employee, error)
	Update(context.Context, *Employee) (*Employee, error)
	FindByID(context.Context, int) (*Employee, error)
	ListAll(context.Context) ([]*Employee, error)
}

// CreateEmployee creates a Employee, and returns the new Employee.
func (uc *TaskUsecase) CreateEmployee(ctx context.Context, g *Employee) (*Employee, error) {
	uc.log.WithContext(ctx).Infof("CreateEmployee: %v", g.Name)

	// 检查医院是否存在
	if _, err := uc.hr.FindByID(ctx, g.HospitalID); err != nil {
		uc.log.WithContext(ctx).Errorw("CreateEmployee: hospital not found", "hospital_id", g.HospitalID, "err", err)
		return nil, ErrHospitalNotFound
	}

	return uc.er.Save(ctx, g)
}
