package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"task_manager/ent/employee"
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

// Save 保存医院
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

// Update 更新医院
func (e employeeRepo) Update(ctx context.Context, employee *biz.Employee) (*biz.Employee, error) {
	// TODO implement me
	panic("implement me")
}

// FindByID 根据ID查找医院
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

// ListAll 列出医院的所有员工
func (e employeeRepo) ListAll(ctx context.Context, hospitalID, limit, offset int) ([]*biz.Employee, int, error) {

	q := e.data.Client().Employee.Query().
		Where(employee.HospitalIDEQ(hospitalID))

	emps, err := q.Limit(limit).Offset(offset).All(ctx)

	if err != nil {
		return nil, 0, err
	}

	count, err := q.Count(ctx)

	employees := make([]*biz.Employee, 0, len(emps))
	for _, v := range emps {
		employees = append(employees, &biz.Employee{
			ID:         v.ID,
			Name:       v.Name,
			HospitalID: v.HospitalID,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		})
	}

	return employees, count, nil
}
