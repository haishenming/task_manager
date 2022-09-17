package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"task_manager/internal/biz"
)

type hospitalRepo struct {
	data *Data
	log  *log.Helper
}

// NewHospitalRepo .
func NewHospitalRepo(data *Data, logger log.Logger) biz.HospitalRepo {
	return &hospitalRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h hospitalRepo) Save(ctx context.Context, hospital *biz.Hospital) (*biz.Hospital, error) {
	hos, err := h.data.Client().Hospital.Create().
		SetName(hospital.Name).
		SetAddress(hospital.Address).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	hospital.ID = hos.ID
	hospital.CreatedAt = hos.CreatedAt
	hospital.UpdatedAt = hos.UpdatedAt

	return hospital, nil
}

func (h hospitalRepo) Update(ctx context.Context, hospital *biz.Hospital) (*biz.Hospital, error) {
	// TODO implement me
	panic("implement me")
}

func (h hospitalRepo) FindByID(ctx context.Context, i int) (*biz.Hospital, error) {
	// 通过ID查找
	hos, err := h.data.Client().Hospital.Get(ctx, i)
	if err != nil {
		return nil, err
	}

	return &biz.Hospital{
		ID:        hos.ID,
		Name:      hos.Name,
		Address:   hos.Address,
		CreatedAt: hos.CreatedAt,
		UpdatedAt: hos.UpdatedAt,
	}, nil
}

func (h hospitalRepo) ListByHello(ctx context.Context, s string) ([]*biz.Hospital, error) {
	// TODO implement me
	panic("implement me")
}

func (h hospitalRepo) ListAll(ctx context.Context) ([]*biz.Hospital, error) {
	// TODO implement me
	panic("implement me")
}
