package kafka_cli

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"io"
)

type ConsumerCmd struct {
	Group string
	Topic string
}

func (c ConsumerCmd) Run(out io.Writer) error {
	client := kafka.NewClient("localhost:9092")
	ctx := context.Background()
	res, err := client.ConsumerOffsets(ctx, kafka.TopicAndGroup{
		Topic:   c.Topic,
		GroupId: c.Group,
	})
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
