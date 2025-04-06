package adapters

import (
	"gateway/internal/adapters/broker"
	"gateway/internal/adapters/grpc/auth"
	"gateway/internal/adapters/grpc/med"
	"gateway/internal/adapters/grpc/mri"
)

type Adapter struct {
	AuthAdapter   auth.AuthAdapter
	MedAdapter    med.MedAdapter
	MriAdapter    mri.MriAdapter
	BrokerAdapter broker.BrokerAdapter
}

func New(
	AuthAdapter auth.AuthAdapter,
	MedAdapter med.MedAdapter,
	MriAdapter mri.MriAdapter,
	BrokerAdapter broker.BrokerAdapter,
) Adapter {
	return Adapter{
		AuthAdapter:   AuthAdapter,
		MedAdapter:    MedAdapter,
		MriAdapter:    MriAdapter,
		BrokerAdapter: BrokerAdapter,
	}
}
