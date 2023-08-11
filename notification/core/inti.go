package core

import (
	"database/sql"
	"log"
	"os/signal"
	"syscall"
	"time"

	// "time"
	//Grupo
	_messageKafka "notification/core/grupo/delivery/kafka"
	_messageRepo "notification/core/grupo/repository"
	_messageUcase "notification/core/grupo/usecase"

	//Sala
	_salaKafka "notification/core/sala/delivery/kafka"
	_salaRepo "notification/core/sala/repository"
	_salaUcase "notification/core/sala/usecase"

	_utilUcase "notification/core/util/usecase"

	"os"

	_firebase "firebase.google.com/go"
)

func Init(db *sql.DB, firebase *_firebase.App) {
	timeout := time.Duration(5) * time.Second
	utilU := _utilUcase.NewUseCase()
	grupoRepo := _messageRepo.NewRepository(db)
	grupoUcase := _messageUcase.NewUseCase(grupoRepo, firebase, timeout)

	grupoKafka := _messageKafka.NewKafkaHandler(grupoUcase)

	salaRepo := _salaRepo.NewRepository(db)
	salaUseCase := _salaUcase.NewUseCase(salaRepo, firebase, timeout, utilU)
	salaKafka := _salaKafka.NewKafkaHandler(salaUseCase)

	go salaKafka.SalaReservationConflictConsumer()
	go grupoKafka.MessageGroupConsumer()
	go grupoKafka.SalaCreationConsumer()
	go salaKafka.SalaHasBennReservedConsumer()

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
