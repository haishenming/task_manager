package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "task_manager/api/task/v1"
)

var (
	// ErrTaskNotFound is task not found.
	ErrTaskNotFound = errors.NotFound(v1.ErrorReason_TASK_NOT_FOUND.String(), "task not found")
)

// Task is a Task model.
type Task struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Priority    v1.TaskPriority `json:"priority"`
	Status      v1.TaskStatus   `json:"status"`
	EmployeeID  int             `json:"employee_id"`
	HospitalID  int             `json:"hospital_id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// TaskRepo is a Greater repo.
type TaskRepo interface {
	Save(context.Context, *Task) (*Task, error)
	Update(context.Context, *Task) (*Task, error)
	FindByID(context.Context, int) (*Task, error)
	ListAll(ctx context.Context, query *TaskQuery, limit, offset int) ([]*Task, int, error)
	AssignTask(context.Context, int, int) error
}

// TaskUsecase is a Task usecase.
type TaskUsecase struct {
	tr TaskRepo
	hr HospitalRepo
	er EmployeeRepo

	log *log.Helper
}

// TaskQuery is a Task query.
type TaskQuery struct {
	EmployeeID int
	HospitalID int
}

// NewTaskUsecase new a Task usecase.
func NewTaskUsecase(
	tr TaskRepo,
	er EmployeeRepo,
	hr HospitalRepo,
	logger log.Logger) *TaskUsecase {

	return &TaskUsecase{
		tr:  tr,
		er:  er,
		hr:  hr,
		log: log.NewHelper(logger)}
}

// CreateTask creates a Task, and returns the new Task.
func (uc *TaskUsecase) CreateTask(ctx context.Context, g *Task) (*Task, error) {
	uc.log.WithContext(ctx).Infof("CreateTask: %v", g.Title)

	// 检查医院是否存在
	if _, err := uc.hr.FindByID(ctx, g.HospitalID); err != nil {
		uc.log.WithContext(ctx).Errorw("CreateTask: hospital not found", "hospital_id", g.HospitalID, "err", err)
		return nil, ErrHospitalNotFound
	}

	return uc.tr.Save(ctx, g)
}

// UpdateTask updates a Task, and returns the new Task.
func (uc *TaskUsecase) UpdateTask(ctx context.Context, g *Task) (*Task, error) {
	uc.log.WithContext(ctx).Infof("UpdateTask: %v", g.Title)

	// 检查Task是否存在
	if _, err := uc.tr.FindByID(ctx, g.ID); err != nil {
		uc.log.WithContext(ctx).Errorw("UpdateTask: task not found", "task_id", g.ID, "err", err)
		return nil, ErrTaskNotFound
	}

	return uc.tr.Update(ctx, g)
}

// AssignTask 分配任务
func (uc *TaskUsecase) AssignTask(ctx context.Context, taskID int, employeeID int) error {
	uc.log.WithContext(ctx).Infof("AssignTask: %v", taskID)

	// 检查Task是否存在
	if _, err := uc.tr.FindByID(ctx, taskID); err != nil {
		uc.log.WithContext(ctx).Errorw("AssignTask: task not found", "task_id", taskID, "err", err)
		return ErrTaskNotFound
	}

	// 检查员工是否存在
	if _, err := uc.er.FindByID(ctx, employeeID); err != nil {
		uc.log.WithContext(ctx).Errorw("AssignTask: employee not found", "employee_id", employeeID, "err", err)
		return ErrEmployeeNotFound
	}

	return uc.tr.AssignTask(ctx, taskID, employeeID)
}

// ListByEmployeeID 根据员工ID获取任务列表
func (uc *TaskUsecase) ListByEmployeeID(ctx context.Context, employeeID int, limit, offset int) ([]*Task, int, error) {
	uc.log.WithContext(ctx).Infof("ListByEmployeeID: %v", employeeID)

	// 检查员工是否存在
	if _, err := uc.er.FindByID(ctx, employeeID); err != nil {
		uc.log.WithContext(ctx).Errorw("ListByEmployeeID: employee not found", "employee_id", employeeID, "err", err)
		return nil, 0, ErrEmployeeNotFound
	}

	return uc.tr.ListAll(ctx, &TaskQuery{EmployeeID: employeeID}, limit, offset)
}

// ListByHospitalID 根据医院ID获取任务列表
func (uc *TaskUsecase) ListByHospitalID(ctx context.Context, hospitalID int, limit, offset int) ([]*Task, int, error) {
	uc.log.WithContext(ctx).Infof("ListByHospitalID: %v", hospitalID)

	// 检查医院是否存在
	if _, err := uc.hr.FindByID(ctx, hospitalID); err != nil {
		uc.log.WithContext(ctx).Errorw("ListByHospitalID: hospital not found", "hospital_id", hospitalID, "err", err)
		return nil, 0, ErrHospitalNotFound
	}

	return uc.tr.ListAll(ctx, &TaskQuery{HospitalID: hospitalID}, limit, offset)
}
