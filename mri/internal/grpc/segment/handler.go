package segment

import (
	"context"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services/segment"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SegmentHandler interface {
	CreateSegment(ctx context.Context, in *pb.CreateSegmentIn) (*pb.CreateSegmentOut, error)
	DeleteSegment(ctx context.Context, in *pb.DeleteSegmentIn) (*empty.Empty, error)
	UpdateSegment(ctx context.Context, in *pb.UpdateSegmentIn) (*pb.UpdateSegmentOut, error)
}

type handler struct {
	segmentSrv segment.Service
}

func New(
	segmentSrv segment.Service,
) SegmentHandler {
	return &handler{
		segmentSrv: segmentSrv,
	}
}

func (h *handler) CreateSegment(ctx context.Context, in *pb.CreateSegmentIn) (*pb.CreateSegmentOut, error) {
	id, err := h.segmentSrv.CreateSegment(ctx, domain.Segment{
		ImageID:  uuid.MustParse(in.ImageId),
		NodeID:   uuid.MustParse(in.NodeId),
		Contor:   in.Contor,
		Knosp012: in.Knosp_012,
		Knosp3:   in.Knosp_3,
		Knosp4:   in.Knosp_4,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateSegmentOut{
		Id: id.String(),
	}, nil
}

func (h *handler) DeleteSegment(ctx context.Context, in *pb.DeleteSegmentIn) (*empty.Empty, error) {
	if err := h.segmentSrv.DeleteSegment(ctx, uuid.MustParse(in.Id)); err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &empty.Empty{}, nil
}

func (h *handler) UpdateSegment(ctx context.Context, in *pb.UpdateSegmentIn) (*pb.UpdateSegmentOut, error) {
	segment, err := h.segmentSrv.UpdateSegment(
		ctx,
		uuid.MustParse(in.Id),
		segment.UpdateSegment{
			Knosp012: in.Knosp_012,
			Knosp3:   in.Knosp_3,
			Knosp4:   in.Knosp_4,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateSegmentOut{
		Segment: DomainSegmentToPb(&segment),
	}, nil
}
