package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "task_manager/api/task/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Task is a Task model.
type Task struct {
	Hello string
}

// TaskRepo is a Greater repo.
type TaskRepo interface {
	Save(context.Context, *Task) (*Task, error)
	Update(context.Context, *Task) (*Task, error)
	FindByID(context.Context, int64) (*Task, error)
	ListByHello(context.Context, string) ([]*Task, error)
	ListAll(context.Context) ([]*Task, error)
}

// TaskUsecase is a Task usecase.
type TaskUsecase struct {
	repo TaskRepo
	log  *log.Helper
}

// NewTaskUsecase new a Task usecase.
func NewTaskUsecase(repo TaskRepo, logger log.Logger) *TaskUsecase {
	return &TaskUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateTask creates a Task, and returns the new Task.
func (uc *TaskUsecase) CreateTask(ctx context.Context, g *Task) (*Task, error) {
	uc.log.WithContext(ctx).Infof("CreateTask: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
