package mriupload

import (
	"context"
	"errors"
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	pb "mri/internal/generated/broker/consume/mriupload"
	"mri/internal/services/image"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

const (
	groupID = "mriupload"
	topic   = "mriupload"
)

type subscriber struct {
	imageSrv image.Service
}

func New(
	imageSrv image.Service,
) brokerlib.SubscriberStrategy {
	return &subscriber{
		imageSrv: imageSrv,
	}
}

func (h *subscriber) GetConfig() brokerlib.SubscriberConfig {
	return brokerlib.SubscriberConfig{
		GroupID: groupID,
		Topics:  []string{topic},
	}
}

func (h *subscriber) ProcessMessage(ctx context.Context, msg []byte) error {
	var event pb.MriUpload
	if err := proto.Unmarshal(msg, &event); err != nil {
		return errors.New("wrong msg type. mriupload required")
	}

	if err := h.imageSrv.SplitMri(ctx, uuid.MustParse(event.MriId)); err != nil {
		return fmt.Errorf("process mriupload: %w", err)
	}
	return nil
}
