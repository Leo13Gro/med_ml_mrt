package uziprocessed

import (
	"context"
	"errors"
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	"uzi/internal/domain"
	pb "uzi/internal/generated/broker/consume/uziprocessed"
	"uzi/internal/services/node"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

const (
	groupID = "mriprocessed"
	topic   = "mriprocessed"
)

type subscriber struct {
	nodeSrv node.Service
}

func New(
	nodeSrv node.Service,
) brokerlib.SubscriberStrategy {
	return &subscriber{
		nodeSrv: nodeSrv,
	}
}

func (h *subscriber) GetConfig() brokerlib.SubscriberConfig {
	return brokerlib.SubscriberConfig{
		GroupID: groupID,
		Topics:  []string{topic},
	}
}

func (h *subscriber) ProcessMessage(ctx context.Context, msg []byte) error {
	var event pb.MriProcessed
	if err := proto.Unmarshal(msg, &event); err != nil {
		return errors.New("wrong msg type. uziprocessed required")
	}

	nodes := make([]domain.Node, 0, len(event.Nodes))
	segments := make([]domain.Segment, 0, len(event.Segments))

	for _, v := range event.Nodes {
		nodes = append(nodes, domain.Node{
			Id:       uuid.MustParse(v.Id),
			MriID:    uuid.MustParse(v.MriId),
			Knosp012: v.Knosp_012,
			Knosp3:   v.Knosp_3,
			Knosp4:   v.Knosp_4,
		})
	}

	for _, v := range event.Segments {
		segments = append(segments, domain.Segment{
			Id:       uuid.MustParse(v.Id),
			ImageID:  uuid.MustParse(v.ImageId),
			NodeID:   uuid.MustParse(v.NodeId),
			Contor:   v.Contor,
			Knosp012: v.Knosp_012,
			Knosp3:   v.Knosp_3,
			Knosp4:   v.Knosp_4,
		})
	}

	if err := h.nodeSrv.InsertAiNodeWithSegments(ctx, nodes, segments); err != nil {
		return fmt.Errorf("isert ai nodes && segments: %w", err)
	}
	return nil
}
