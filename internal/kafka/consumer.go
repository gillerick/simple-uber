package kafka

import (
	"context"
	"time"
)
import "github.com/segmentio/kafka-go"
import "log"

// MessageHandler provides message processing capabilities
type MessageHandler interface {
	HandleMessage(ctx context.Context, msg []byte) error
}

// Consumer defines a Kafka messages consumer
type Consumer struct {
	reader         *kafka.Reader
	messageHandler MessageHandler
}

// Run starts the Kafka consumer
func (p *Consumer) Run() error {
	var err error
	for {
		ctx := context.Background()
		var m kafka.Message

		// Get the next message to consume from the broker
		m, err = p.reader.FetchMessage(ctx)
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}

		//	Processing loop
		err = p.messageHandler.HandleMessage(ctx, m.Value)
		if err != nil {
			//	ToDO: Add error logic
		}

		// Mark the processed messages as committed on the broker
		err = p.reader.CommitMessages(ctx, m)
		if err != nil {
			log.Println("commit failed with error %w", err)
			break
		}
	}
	return err

}

// NewConsumer sets up a new Kafka consumer for the given topic using the provided handler
func NewConsumer(brokerAddrs []string, consumerGroupId string, topic string, messageHandler MessageHandler) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokerAddrs,
		GroupID:        consumerGroupId,
		Topic:          topic,
		CommitInterval: time.Second,
	})

	proc := &Consumer{
		messageHandler: messageHandler,
		reader:         reader,
	}

	return proc
}
