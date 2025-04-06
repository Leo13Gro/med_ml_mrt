package mri

import (
	pb "gateway/internal/generated/grpc/client/mri"

	"google.golang.org/grpc"
)

type MriAdapter interface {
	pb.MriSrvClient
}

type adapter struct {
	pb.MriSrvClient
}

func New(
	conn *grpc.ClientConn,
) MriAdapter {
	client := pb.NewMriSrvClient(conn)

	return &adapter{
		MriSrvClient: client,
	}
}
