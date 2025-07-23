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

//
//package kafka
//
//import (
//"context"
//appErr "draft-zadania-1/errors"
//"github.com/twmb/franz-go/pkg/kgo"
//"log"
//)
//
//type KafkaProducer struct {
//	client   *kgo.Client
//	messages chan kafkaMsg
//	cancel   context.CancelFunc
//}
//
//type kafkaMsg struct {
//	topic string
//	value []byte
//}
//
//func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
//	client, err := kgo.NewClient(kgo.SeedBrokers(brokers...))
//	if err != nil {
//		return nil, appErr.ErrInternal
//	}
//
//	ctx, cancel := context.WithCancel(context.Background())
//
//	producer := &KafkaProducer{
//		client:   client,
//		messages: make(chan kafkaMsg, 100),
//		cancel:   cancel,
//	}
//
//	go producer.run(ctx)
//
//	return producer, nil
//}
//
//func (p *KafkaProducer) run(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		case msg := <-p.messages:
//			log.Printf("[PRODUCER] Sending message to Kafka topic '%s': %s\n", msg.topic, string(msg.value))
//
//			err := p.client.ProduceSync(ctx, &kgo.Record{
//				Topic: msg.topic,
//				Value: msg.value,
//			}).FirstErr()
//
//			if err != nil {
//				log.Printf("[PRODUCER] Kafka error: %v\n", err)
//			} else {
//				log.Println("[PRODUCER] Message successfully sent to Kafka")
//			}
//		}
//	}
//}
//
//func (p *KafkaProducer) Enqueue(topic string, value []byte) {
//	select {
//	case p.messages <- kafkaMsg{topic: topic, value: value}:
//		log.Printf("[ENQUEUE] Message queued for topic '%s': %s\n", topic, string(value))
//	default:
//		log.Println("[ENQUEUE] Channel full – message dropped")
//	}
//}
//
//func (p *KafkaProducer) Close() {
//	p.cancel()
//	p.client.Close()
//}
