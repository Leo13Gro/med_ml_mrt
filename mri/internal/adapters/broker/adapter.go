package broker

import (
	"fmt"

	"github.com/WantBeASleep/goooool/brokerlib"

	uzicompletepb "uzi/internal/generated/broker/produce/uzicomplete"
	uzisplittedpb "uzi/internal/generated/broker/produce/uzisplitted"

	"google.golang.org/protobuf/proto"
)

const (
	mrisplittedTopic = "mrisplitted"
	mricompleteTopic = "mricomplete"
)

type BrokerAdapter interface {
	SendMriSplitted(msg *uzisplittedpb.MriSplitted) error
	SendMriComplete(msg *uzicompletepb.MriComplete) error
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

func (a *adapter) SendMriSplitted(msg *uzisplittedpb.MriSplitted) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal uzisplitted event: %w", err)
	}
	return a.producer.Send(mrisplittedTopic, "52", payload)
}

func (a *adapter) SendMriComplete(msg *uzicompletepb.MriComplete) error {
	// TODO: когда будем делать партицированние пробрасывать сюда ключи
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal uzicomplete event: %w", err)
	}
	return a.producer.Send(mricompleteTopic, "52", payload)
}
