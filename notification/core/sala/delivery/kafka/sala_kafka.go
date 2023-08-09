package kafka

import (
	"context"
	"fmt"
	"log"
	r "notification/domain/repository"

	"github.com/segmentio/kafka-go"
)



type SalaKafkaHander struct {
	salaU r.SalaUseCase
}

func NewKafkaHandler(salaU r.SalaUseCase) SalaKafkaHander {
	return SalaKafkaHander{
		salaU: salaU,
	}
}

func (k *SalaKafkaHander) SalaHasBeenReserved(){
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9094"},
		Topic:     "sala-creation-group",
		GroupID:   "consumer-group-sala-has-been-reserved",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
}