package server

import (
	"context"
	"fmt"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/auth"
	"composition-api/internal/server/download"
	exam "composition-api/internal/server/exam"
	"composition-api/internal/server/med"
	"composition-api/internal/server/register"
	services "composition-api/internal/services"
)

type server struct {
	auth.AuthRoute
	exam.ExamRoute
	med.MedRoute
	register.RegisterRoute
	download.DownloadRoute
}

func New(services *services.Services) api.Handler {
	examRoute := exam.NewExamRoute(services)
	authRoute := auth.NewAuthRoute(services)
	medRoute := med.NewMedRoute(services)
	registerRoute := register.NewRegisterRoute(services)
	downloadRoute := download.NewDownloadRoute(services)

	return &server{
		ExamRoute:     examRoute,
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
