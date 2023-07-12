package usecase

import (
	"context"
	"encoding/json"
	"log"
	r "message/domain/repository"
	"time"

	"github.com/segmentio/kafka-go"
)

type grupoUcase struct {
	timeout   time.Duration
	grupoRepo r.GrupoRepository
	kafkaW    *kafka.Writer
}

func NewUseCase(timeout time.Duration, grupoRepo r.GrupoRepository) r.GrupoUseCase {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9094"),
		Topic:   "notification-message-group",
		Balancer: &kafka.LeastBytes{},
	}
	return &grupoUcase{
		timeout:   timeout,
		grupoRepo: grupoRepo,
		kafkaW: w,
	}
}

func (u *grupoUcase) SaveGrupoMessage(ctx context.Context, d *r.MessageGrupo) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer func() {
		cancel()
	} ()
	err = u.grupoRepo.SaveGrupoMessage(ctx, d)
	go func() {
		json,err := json.Marshal(d)
		if err != nil {
		log.Println("Fail to parse",err)
		}
		err = u.kafkaW.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Message"),
				Value: json,
			},
		)
		if err != nil {
			log.Println("failed to write messages:", err)
		}
	}()
	
	return

}
