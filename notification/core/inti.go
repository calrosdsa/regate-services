package core

import (
	"database/sql"
	"log"
	"os/signal"
	"syscall"
	// "time"

	_messageKafka "notification/core/grupo/delivery/kafka"
	_messageRepo "notification/core/grupo/repository"
	_messageUcase "notification/core/grupo/usecase"
	"os"

	_firebase "firebase.google.com/go"
)

func Init(db *sql.DB, firebase *_firebase.App) {
	grupoRepo := _messageRepo.NewRepository(db)
	grupoUcase := _messageUcase.NewUseCase(grupoRepo, firebase)

	grupoKafka := _messageKafka.NewKafkaHandler(grupoUcase)

	go grupoKafka.MessageGroupConsumer()
	go grupoKafka.SalaCreationConsumer()	

    quitChannel := make(chan os.Signal, 1)
    signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
    <-quitChannel
    //time for cleanup before exit
    log.Println("Adios!")
}


// func forever() {
//     for {
//         log.Printf("%v+\n", time.Now())
//         time.Sleep(time.Second)
//     }
// }

