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

// Save 保存医院
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

// Update 更新医院
func (h hospitalRepo) Update(ctx context.Context, hospital *biz.Hospital) (*biz.Hospital, error) {
	// TODO implement me
	panic("implement me")
}

// FindByID 根据ID查找医院
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

// ListAll 列出所有医院
func (h hospitalRepo) ListAll(ctx context.Context, limit, offset int) ([]*biz.Hospital, int, error) {
	var err error

	q := h.data.Client().Hospital.Query()

	hoss, err := q.Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	count, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	hospitals := make([]*biz.Hospital, 0, len(hoss))
	for _, hos := range hoss {
		hospitals = append(hospitals, &biz.Hospital{
			ID:        hos.ID,
			Name:      hos.Name,
			Address:   hos.Address,
			CreatedAt: hos.CreatedAt,
			UpdatedAt: hos.UpdatedAt,
		})
	}

	return hospitals, count, nil
}
