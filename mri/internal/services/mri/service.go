package mri

import (
	"context"
	"errors"
	"fmt"
	"time"

	"mri/internal/domain"
	"mri/internal/repository"
	"mri/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreateMri(ctx context.Context, mri domain.Mri) (uuid.UUID, error)
	GetMriByID(ctx context.Context, id uuid.UUID) (domain.Mri, error)
	GetMrisByPatientID(ctx context.Context, patientID uuid.UUID) ([]domain.Mri, error)
	GetMriEchographicsByID(ctx context.Context, id uuid.UUID) (domain.Echographic, error)
	UpdateMri(ctx context.Context, id uuid.UUID, update UpdateMri) (domain.Mri, error)
	UpdateEchographic(ctx context.Context, id uuid.UUID, update UpdateEchographic) (domain.Echographic, error)
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) CreateMri(ctx context.Context, mri domain.Mri) (uuid.UUID, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("begin transaction: %w", err)
	}

	mri.Id = uuid.New()
	mri.Checked = false
	mri.CreateAt = time.Now()

	if err := s.dao.NewMriQuery(ctx).InsertMri(entity.Mri{}.FromDomain(mri)); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return uuid.Nil, fmt.Errorf("insert mri: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.NewEchographicQuery(ctx).InsertEchographic(entity.Echographic{Id: mri.Id}); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return uuid.Nil, fmt.Errorf("insert echographic: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit transaction: %w", err)
	}

	return mri.Id, nil
}

func (s *service) GetMriByID(ctx context.Context, id uuid.UUID) (domain.Mri, error) {
	mri, err := s.dao.NewMriQuery(ctx).GetMriByPK(id)
	if err != nil {
		return domain.Mri{}, fmt.Errorf("get mri by pk: %w", err)
	}

	return mri.ToDomain(), nil
}

func (s *service) GetMrisByPatientID(ctx context.Context, patientID uuid.UUID) ([]domain.Mri, error) {
	mris, err := s.dao.NewMriQuery(ctx).GetMrisByPatientID(patientID)
	if err != nil {
		return nil, fmt.Errorf("get mri by pk: %w", err)
	}

	domainMris := make([]domain.Mri, 0, len(mris))
	for _, v := range mris {
		domainMris = append(domainMris, v.ToDomain())
	}

	return domainMris, nil
}

func (s *service) GetMriEchographicsByID(ctx context.Context, id uuid.UUID) (domain.Echographic, error) {
	echographics, err := s.dao.NewEchographicQuery(ctx).GetEchographicByPK(id)
	if err != nil {
		return domain.Echographic{}, fmt.Errorf("get mri echographics pk: %w", err)
	}

	return echographics.ToDomain(), nil
}

func (s *service) UpdateMri(ctx context.Context, id uuid.UUID, update UpdateMri) (domain.Mri, error) {
	mri, err := s.GetMriByID(ctx, id)
	if err != nil {
		return domain.Mri{}, fmt.Errorf("get mri by id: %w", err)
	}
	update.Update(&mri)

	if _, err := s.dao.NewMriQuery(ctx).UpdateMri(entity.Mri{}.FromDomain(mri)); err != nil {
		return domain.Mri{}, fmt.Errorf("update mri: %w", err)
	}

	return mri, nil
}

func (s *service) UpdateEchographic(ctx context.Context, id uuid.UUID, update UpdateEchographic) (domain.Echographic, error) {
	echographic, err := s.GetMriEchographicsByID(ctx, id)
	if err != nil {
		return domain.Echographic{}, fmt.Errorf("get mri by id: %w", err)
	}
	update.Update(&echographic)

	if _, err := s.dao.NewEchographicQuery(ctx).UpdateEchographic(entity.Echographic{}.FromDomain(echographic)); err != nil {
		return domain.Echographic{}, fmt.Errorf("update echographic: %w", err)
	}

	return echographic, nil
}
