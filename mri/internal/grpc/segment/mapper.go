package segment

import (
	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

func DomainSegmentToPb(d *domain.Segment) *pb.Segment {
	return &pb.Segment{
		Id:        d.Id.String(),
		ImageId:   d.ImageID.String(),
		NodeId:    d.NodeID.String(),
		Contor:    d.Contor,
		Knosp_012: d.Knosp012,
		Knosp_3:   d.Knosp3,
		Knosp_4:   d.Knosp4,
	}
}
