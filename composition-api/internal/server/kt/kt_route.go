package kt

import (
	"composition-api/internal/server/kt/kt"
	"composition-api/internal/services"
)

type KTRoute interface {
	kt.KTHandler
}

type ktRoute struct {
	kt.KTHandler
}

func NewKTRoute(services *services.Services) KTRoute {
	ktHandler := kt.NewHandler(services)
	return &ktRoute{
		KTHandler: ktHandler,
	}
}
