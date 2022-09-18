package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "task_manager/api/task/v1"
)

var (
	// ErrHospitalNotFound is task not found.
	ErrHospitalNotFound = errors.NotFound(v1.ErrorReason_HOSPITAL_NOT_FOUND.String(), "hospital not found")
)

// Hospital is a Hospital model.
type Hospital struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HospitalRepo is a Hospital repo.
type HospitalRepo interface {
	Save(context.Context, *Hospital) (*Hospital, error)
	Update(context.Context, *Hospital) (*Hospital, error)
	FindByID(context.Context, int) (*Hospital, error)
	ListAll(ctx context.Context, limit, offset int) ([]*Hospital, int, error)
}

// CreateHospital creates a Hospital, and returns the new Hospital.
func (uc *TaskUsecase) CreateHospital(ctx context.Context, g *Hospital) (*Hospital, error) {
	uc.log.WithContext(ctx).Infof("CreateHospital: %v", g.Name)
	return uc.hr.Save(ctx, g)
}

func (uc *TaskUsecase) GetHospitals(ctx context.Context, limit, offset int) ([]*Hospital, int, error) {
	uc.log.WithContext(ctx).Infof("GetHospitals")
	return uc.hr.ListAll(ctx, limit, offset)
}
