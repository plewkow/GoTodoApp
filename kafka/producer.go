package kafka

import (
	"context"
	appErr "draft-zadania-1/errors"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaProducer struct {
	client *kgo.Client
	topic  string
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, appErr.ErrInternal
	}
	return &KafkaProducer{
		client: cl,
	}, nil
}

func (p *KafkaProducer) Close() {
	p.client.Close()
}

func (p *KafkaProducer) Produce(ctx context.Context, topic string, value []byte) error {
	record := &kgo.Record{
		Topic: topic,
		Value: value,
	}

	err := p.client.ProduceSync(ctx, record).FirstErr()
	if err != nil {
		return appErr.ErrInternal
	}
	return nil
}
