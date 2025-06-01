package kt

import (
	"context"

	"composition-api/internal/adapters"
	dbus "composition-api/internal/dbus/producers"
	domain "composition-api/internal/domain/kt"
	"composition-api/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, arg CreateKtArg) (uuid.UUID, error)

	GetByID(ctx context.Context, id uuid.UUID) (domain.KT, error)
}

type service struct {
	adapters *adapters.Adapters
	dao      repository.DAO
	dbus     dbus.Producer
}

func New(
	adapters *adapters.Adapters,
	dao repository.DAO,
	dbus dbus.Producer,
) Service {
	return &service{
		adapters: adapters,
		dao:      dao,
		dbus:     dbus,
	}
}
