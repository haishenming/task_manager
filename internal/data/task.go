package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "task_manager/api/task/v1"
	"task_manager/ent/task"
	"task_manager/internal/biz"
)

type taskRepo struct {
	data *Data
	log  *log.Helper
}

// NewTaskRepo .
func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *taskRepo) Save(ctx context.Context, g *biz.Task) (*biz.Task, error) {
	t, err := r.data.Client().Task.Create().
		SetTitle(g.Title).
		SetDescription(g.Description).
		SetPriority(int8(g.Priority)).
		SetHospitalID(g.HospitalID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	g.ID = t.ID
	g.CreatedAt = t.CreatedAt
	g.UpdatedAt = t.UpdatedAt

	return g, nil
}

func (r *taskRepo) Update(ctx context.Context, g *biz.Task) (*biz.Task, error) {
	t, err := r.data.Client().Task.UpdateOneID(g.ID).SetTitle(g.Title).
		SetDescription(g.Description).
		SetPriority(int8(g.Priority)).
		SetStatus(int8(g.Status)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	g.Title = t.Title
	g.Description = t.Description
	g.Priority = v1.TaskPriority(t.Priority)
	g.Status = v1.TaskStatus(t.Status)
	g.HospitalID = t.HospitalID
	g.EmployeeID = t.EmployeeID
	g.UpdatedAt = t.UpdatedAt
	g.CreatedAt = t.CreatedAt

	return g, nil
}

func (r *taskRepo) FindByID(ctx context.Context, i int) (*biz.Task, error) {
	// 通过ID查找
	tas, err := r.data.Client().Task.Get(ctx, i)
	if err != nil {
		return nil, err
	}

	return &biz.Task{
		ID:          tas.ID,
		Title:       tas.Title,
		Description: tas.Description,
		Priority:    v1.TaskPriority(tas.Priority),
		Status:      v1.TaskStatus(tas.Status),
		EmployeeID:  tas.EmployeeID,
		HospitalID:  tas.HospitalID,
		CreatedAt:   tas.CreatedAt,
		UpdatedAt:   tas.UpdatedAt,
	}, nil
}

// ListAll 任务列表
func (r *taskRepo) ListAll(ctx context.Context, query *biz.TaskQuery, limit, offset int) ([]*biz.Task, int, error) {
	var err error

	q := r.data.Client().Task.Query()
	if query != nil {
		if query.HospitalID != 0 {
			q = q.Where(task.HospitalIDEQ(query.HospitalID))
		}

		if query.EmployeeID != 0 {
			q = q.Where(task.EmployeeIDEQ(query.EmployeeID))
		}
	}

	tasks, err := q.Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	count, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	var res []*biz.Task
	for _, t := range tasks {
		res = append(res, &biz.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Priority:    v1.TaskPriority(t.Priority),
			Status:      v1.TaskStatus(t.Status),
			EmployeeID:  t.EmployeeID,
			HospitalID:  t.HospitalID,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	return res, count, nil
}

// AssignTask 分配任务
func (r *taskRepo) AssignTask(ctx context.Context, taskID int, employeeID int) error {
	_, err := r.data.Client().Task.UpdateOneID(taskID).
		SetEmployeeID(employeeID).
		SetStatus(int8(v1.TaskStatus_TASK_STATUS_OPEN)).
		Save(ctx)
	return err
}
