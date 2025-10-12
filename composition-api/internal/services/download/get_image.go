package download

import (
	"context"
	"io"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *service) GetImage(ctx context.Context, mriID uuid.UUID, imageID uuid.UUID) (io.ReadCloser, error) {
	return s.repo.NewFileRepo().GetFile(
		ctx,
		filepath.Join(
			mriID.String(),
			imageID.String(),
			imageID.String(),
		),
	)
}
