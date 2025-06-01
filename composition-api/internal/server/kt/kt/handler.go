package kt

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type KTHandler interface {
	KtIDGet(ctx context.Context, params api.KtIDGetParams) (api.KtIDGetRes, error)
	KtPost(ctx context.Context, req *api.KtPostReq) (api.KtPostRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) KTHandler {
	return &handler{
		services: services,
	}
}
