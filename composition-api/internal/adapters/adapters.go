package adapters

import (
	"google.golang.org/grpc"

	"composition-api/internal/adapters/auth"
	"composition-api/internal/adapters/exam"
	"composition-api/internal/adapters/med"
	authPB "composition-api/internal/generated/grpc/clients/auth"
	examPB "composition-api/internal/generated/grpc/clients/exam"
	medPB "composition-api/internal/generated/grpc/clients/med"
)

type Adapters struct {
	Exam exam.Adapter
	Auth auth.Adapter
	Med  med.Adapter
}

func NewAdapters(
	examConn *grpc.ClientConn,
	authConn *grpc.ClientConn,
	medConn *grpc.ClientConn,
) *Adapters {
	examClient := examPB.NewExamSrvClient(examConn)
	examAdapter := exam.NewAdapter(examClient)

	authClient := authPB.NewAuthSrvClient(authConn)
	authAdapter := auth.NewAdapter(authClient)

	medClient := medPB.NewMedSrvClient(medConn)
	medAdapter := med.NewAdapter(medClient)

	return &Adapters{
		Exam: examAdapter,
		Auth: authAdapter,
		Med:  medAdapter,
	}
}
