package kafka

import (
	"context"
	appErr "draft-zadania-1/errors"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

func EnsureTopicExists(seeds []string, topicNames []string) error {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
	)
	if err != nil {
		return appErr.ErrInternal
	}
	defer client.Close()

	admin := kadm.NewClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	topicDetails, err := admin.ListTopics(ctx)
	if err != nil {
		return appErr.ErrInternal
	}

	for _, topic := range topicNames {
		if topicDetails.Has(topic) {
			continue
		}

		_, err := admin.CreateTopic(ctx, -1, -1, nil, topic)
		if err != nil {
			return appErr.ErrInternal
		}
	}

	return nil
}
