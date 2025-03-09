package broker

import (
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	pb "gateway/internal/generated/broker/produce/uziupload"

	"google.golang.org/protobuf/proto"
)

const (
	mriuploadTopic = "mriupload"
)

type BrokerAdapter interface {
	SendMriUpload(msg *pb.MriUpload) error
}

// TODO: переписать библу/хотя бы в интерфейс обернуть продьюсера
func New(
	producer brokerlib.Producer,
) BrokerAdapter {
	return &adapter{
		producer: producer,
	}
}

type adapter struct {
	producer brokerlib.Producer
}

func (a *adapter) SendMriUpload(msg *pb.MriUpload) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal uziupload event: %w", err)
	}
	return a.producer.Send(mriuploadTopic, "52", payload)
}
