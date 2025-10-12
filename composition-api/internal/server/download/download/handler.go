package download

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type DownloadHandler interface {
	DownloadMriIDImageIDGet(ctx context.Context, params api.DownloadMriIDImageIDGetParams) (api.DownloadMriIDImageIDGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) DownloadHandler {
	return &handler{
		services: services,
	}
}
