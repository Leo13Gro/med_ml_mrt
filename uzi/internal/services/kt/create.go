package kt

import (
	"context"
	"fmt"
	"time"

	"uzi/internal/domain"
	uziEntity "uzi/internal/repository/uzi/entity"

	"github.com/google/uuid"
)

func (s *service) CreateUzi(ctx context.Context, arg CreateKTArg) (uuid.UUID, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	kt := domain.KT{
		Id:          uuid.New(),
		Checked:     false,
		Author:      arg.Author,
		DeviceID:    arg.DeviceID,
		Status:      domain.UziStatusNew,
		Description: arg.Description,
		CreateAt:    time.Now(),
	}

	if err := s.dao.NewUziQuery(ctx).InsertUzi(uziEntity.Uzi{}.FromDomain(kt)); err != nil {
		return uuid.Nil, fmt.Errorf("insert uzi: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit transaction: %w", err)
	}

	return kt.Id, nil
}
