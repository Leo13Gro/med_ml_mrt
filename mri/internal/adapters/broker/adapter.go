package broker

import (
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	mricompletepb "mri/internal/generated/broker/produce/mricomplete"
	mrisplittedpb "mri/internal/generated/broker/produce/mrisplitted"

	"google.golang.org/protobuf/proto"
)

const (
	mrisplittedTopic = "mrisplitted"
	mricompleteTopic = "mricomplete"
)

type BrokerAdapter interface {
	SendMriSplitted(msg *mrisplittedpb.MriSplitted) error
	SendMriComplete(msg *mricompletepb.MriComplete) error
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

func (a *adapter) SendMriSplitted(msg *mrisplittedpb.MriSplitted) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal mrisplitted event: %w", err)
	}
	return a.producer.Send(mrisplittedTopic, "52", payload)
}

func (a *adapter) SendMriComplete(msg *mricompletepb.MriComplete) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal mricomplete event: %w", err)
	}
	return a.producer.Send(mricompleteTopic, "52", payload)
}
