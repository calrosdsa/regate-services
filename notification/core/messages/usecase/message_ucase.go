package usecase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"log"
	"notification/domain/repository"
	"time"
)

type messageUcase struct {
	messageRepo repository.MessageRepository
	timeout     time.Duration
	firebase    *firebase.App
}

func (m *messageUcase) GetLastMessagesFromGroup(ctx context.Context, id int) ([]repository.MessageGrupo, error) {
	//TODO implement me
	panic("implement me")
}

func (m *messageUcase) GetUsersFromGroup(ctx context.Context, id int) ([]repository.ProfileUser, error) {
	//TODO implement me
	panic("implement me")
}

func NewUseCase(messageRepo repository.MessageRepository, firebase *firebase.App) repository.MessageUseCase {
	return &messageUcase{messageRepo: messageRepo, firebase: firebase}
}

func (u *messageUcase) sendNotification(ctx context.Context, tokens []string) {

	client, err := u.firebase.Messaging(ctx)
	if err != nil {
		log.Println(err)
	}

	// registrationToken := "dlRvPgmLQyaMVqFUuoJRCZ:APA91bHTm16p5Vw87ftsimJ0DyIBD00hd3RdGSJWXS8-vQyO1Mn-ntV8XlaTbpGExrBg9Tqil2FkZTvQN-QmRTJvNRN52mx1jy_OuTyigsM2LG2Q9ThyPMR5T6o0ah9eezDZITGWMNzK"
	message := &messaging.MulticastMessage{
		//Notification: &messaging.Notification{
		//	Title: "Notification Test",
		//	Body:  "Hello React!!",
		//},
		Tokens: tokens,
		Data: map[string]string{
			"title":    "FCM Notification Title ",
			"subtext":  "FCM Notification Sub Title",
			"type":     "999",
			"priority": "high",
		},
	}

	response, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully sent message:", response)
}
