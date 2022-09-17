package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"task_manager/internal/biz"
)

type employeeRepo struct {
	data *Data
	log  *log.Helper
}

// NewEmployeeRepo .
func NewEmployeeRepo(data *Data, logger log.Logger) biz.EmployeeRepo {
	return &employeeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (e employeeRepo) Save(ctx context.Context, employee *biz.Employee) (*biz.Employee, error) {
	emp, err := e.data.Client().Employee.Create().SetName(employee.Name).SetHospitalID(employee.HospitalID).Save(ctx)
	if err != nil {
		return nil, err
	}

	employee.ID = emp.ID
	employee.CreatedAt = emp.CreatedAt
	employee.UpdatedAt = emp.UpdatedAt

	return employee, nil
}

func (e employeeRepo) Update(ctx context.Context, employee *biz.Employee) (*biz.Employee, error) {
	// TODO implement me
	panic("implement me")
}

func (e employeeRepo) FindByID(ctx context.Context, i int) (*biz.Employee, error) {
	// 通过ID查找
	emp, err := e.data.Client().Employee.Get(ctx, i)
	if err != nil {
		return nil, err
	}

	return &biz.Employee{
		ID:         emp.ID,
		Name:       emp.Name,
		HospitalID: emp.HospitalID,
		CreatedAt:  emp.CreatedAt,
		UpdatedAt:  emp.UpdatedAt,
	}, nil
}

func (e employeeRepo) ListAll(ctx context.Context) ([]*biz.Employee, error) {
	// TODO implement me
	panic("implement me")
}
