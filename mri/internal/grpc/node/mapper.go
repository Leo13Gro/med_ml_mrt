package node

import (
	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
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
