package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"notification/domain/repository"
)

type MessageKafkaHandler struct {
	messageUcase repository.MessageUseCase
}

func NewKafkaHandler(messageUcase repository.MessageUseCase) MessageKafkaHandler {
	return MessageKafkaHandler{
		messageUcase: messageUcase,
	}
}

func (k *MessageKafkaHandler) MessageGroupConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9094"},
		Topic:     "notification-message-group",
		GroupID: "consumer-group-messages",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	// r.SetOffset(2)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		log.Println("RUNNN")
		fmt.Printf("message at offset %d: %s = %s\n %s", m.Offset, string(m.Key), string(m.Value),m.Time.Local().String())
		k.messageUcase.SendNotificationToUsersGroup(context.Background(),m.Value)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
