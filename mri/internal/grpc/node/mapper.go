package node

import (
	"mri/internal/domain"
	pb "mri/internal/generated/grpc/service"
)

func DomainNodeToPb(d *domain.Node) *pb.Node {
	return &pb.Node{
		Id:        d.Id.String(),
		Ai:        d.Ai,
		Knosp_012: d.Knosp012,
		Knosp_3:   d.Knosp3,
		Knosp_4:   d.Knosp4,
	}
}
