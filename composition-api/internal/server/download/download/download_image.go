package download

import (
	"context"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) DownloadMriIDImageIDGet(ctx context.Context, params api.DownloadMriIDImageIDGetParams) (api.DownloadMriIDImageIDGetRes, error) {
	image, err := h.services.DownloadService.GetImage(ctx, params.MriID, params.ImageID)
	if err != nil {
		return nil, err
	}

	return &api.DownloadMriIDImageIDGetOK{
		Data: image,
	}, nil
}
