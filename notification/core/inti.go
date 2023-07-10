package core

import (
	"database/sql"
	_firebase "firebase.google.com/go"
	_messageKafka "notification/core/messages/delivery/kafka"
	_messageRepo "notification/core/messages/repository"
	_messageUcase "notification/core/messages/usecase"
)

func Init(db *sql.DB, firebase *_firebase.App) {

	
	messageRepo := _messageRepo.NewRepository(db)
	messageUcase := _messageUcase.NewUseCase(messageRepo, firebase)

	messageKafka := _messageKafka.NewKafkaHandler(messageUcase)
	messageKafka.MessageGroupConsumer()

}
