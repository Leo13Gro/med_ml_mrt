package adapters

import (
	billingPB "composition-api/internal/generated/grpc/clients/billing"

	"google.golang.org/grpc"

	"composition-api/internal/adapters/auth"
	"composition-api/internal/adapters/billing"
	"composition-api/internal/adapters/exam"
	"composition-api/internal/adapters/med"
	authPB "composition-api/internal/generated/grpc/clients/auth"
	examPB "composition-api/internal/generated/grpc/clients/exam"
	medPB "composition-api/internal/generated/grpc/clients/med"
)

type Adapters struct {
	Exam    exam.Adapter
	Auth    auth.Adapter
	Med     med.Adapter
	Billing billing.Adapter
}

func NewAdapters(
	examConn *grpc.ClientConn,
	authConn *grpc.ClientConn,
	medConn *grpc.ClientConn,
	billingConn *grpc.ClientConn,
) *Adapters {
	examClient := examPB.NewExamSrvClient(examConn)
	examAdapter := exam.NewAdapter(examClient)

	authClient := authPB.NewAuthSrvClient(authConn)
	authAdapter := auth.NewAdapter(authClient)

	medClient := medPB.NewMedSrvClient(medConn)
	medAdapter := med.NewAdapter(medClient)

	billingClient := billingPB.NewBillingServiceClient(billingConn)
	billingAdapter := billing.NewAdapter(billingClient)

	return &Adapters{
		Exam:    examAdapter,
		Auth:    authAdapter,
		Med:     medAdapter,
		Billing: billingAdapter,
	}
}
