package kt

import (
	"context"
	"fmt"
	"path/filepath"

	ktuploadpb "composition-api/internal/generated/dbus/produce/ktupload"

	"github.com/google/uuid"
)

func (s *service) Create(ctx context.Context, arg CreateKtArg) (uuid.UUID, error) {
	ktID := uuid.New()

	err := s.dao.NewFileRepo().LoadFile(ctx, filepath.Join(ktID.String(), ktID.String()), arg.File)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("load uzi file to s3: %w", err)
	}

	err = s.dbus.SendKtUpload(ctx, &ktuploadpb.KtUpload{KtId: ktID.String()})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("send kt upload to dbus: %w", err)
	}

	return ktID, nil
}
