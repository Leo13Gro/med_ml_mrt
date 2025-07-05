package kt

import (
	"context"

	dbus "uzi/internal/dbus/producers"
	"uzi/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateUzi(ctx context.Context, arg ) (uuid.UUID, error)
	PrepareKt(ctx context.Context, id uuid.UUID) error
}

type service struct {
	dao  repository.DAO
	dbus dbus.Producer
}

func New(
	dao repository.DAO,
	dbus dbus.Producer,
) Service {
	return &service{
		dao:  dao,
		dbus: dbus,
	}
}
