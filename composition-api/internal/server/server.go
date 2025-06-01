package server

import (
	"context"
	"fmt"

	"composition-api/internal/server/kt"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/auth"
	"composition-api/internal/server/download"
	"composition-api/internal/server/med"
	"composition-api/internal/server/register"
	"composition-api/internal/server/uzi"
	services "composition-api/internal/services"
)

type server struct {
	auth.AuthRoute
	uzi.UziRoute
	kt.KTRoute
	med.MedRoute
	register.RegisterRoute
	download.DownloadRoute
}

func New(services *services.Services) api.Handler {
	uziRoute := uzi.NewUziRoute(services)
	ktRoute := kt.NewKTRoute(services)
	authRoute := auth.NewAuthRoute(services)
	medRoute := med.NewMedRoute(services)
	registerRoute := register.NewRegisterRoute(services)
	downloadRoute := download.NewDownloadRoute(services)

	return &server{
		UziRoute:      uziRoute,
		KTRoute:       ktRoute,
		AuthRoute:     authRoute,
		MedRoute:      medRoute,
		RegisterRoute: registerRoute,
		DownloadRoute: downloadRoute,
	}
}

func (s *server) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Code:    500,
			Message: fmt.Sprint("Необработанная ошибка сервера: ", err.Error()),
		},
	}
}
