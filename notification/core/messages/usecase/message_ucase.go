package usecase

import (
	"context"
	"encoding/json"
	"log"
	"notification/domain/repository"
	// domain "notification/domain"
	"time"


	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

type messageUcase struct {
	messageRepo repository.MessageRepository
	timeout     time.Duration
	firebase    *firebase.App
}

func NewUseCase(messageRepo repository.MessageRepository, firebase *firebase.App) repository.MessageUseCase {
	timeout := time.Duration(5) * time.Second
	return &messageUcase{messageRepo: messageRepo, firebase: firebase,timeout: timeout}
}


func (u *messageUcase) SendNotificationToUsersGroup(ctx context.Context,message []byte) (err error) {
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	var data repository.MessageGrupo
    err = json.Unmarshal(message,&data)
	if err != nil {
		return
	}
	// messages,err := u.messageRepo.GetLastMessagesFromGroup(ctx,data.GrupoId)
	// if err != nil{
	// 	log.Println("Message error",err)
	// }
	fcm_tokens ,err := u.messageRepo.GetUsersFromGroup(ctx ,data.GrupoId)
	if err != nil {
		log.Println("FCM_TOKENS",err)
		return
	}
	tokens := make([]string,len(fcm_tokens))
	for _,val := range fcm_tokens{
		tokens = append(tokens, val.FcmToken)
		u.sendNotifications(ctx,val.FcmToken,message)	
	}
	log.Println("TOKENS",tokens)
	return 
}
func (u *messageUcase) sendNotifications(ctx context.Context, tokens string,messages []byte) {

	client, err := u.firebase.Messaging(ctx)
	if err != nil {
		log.Println(err)
	}
	message := &messaging.Message{
			//Notification: &messaging.Notification{
				//	Title: "Notification Test",
		//	Body:  "Hello React!!",
		//},
		Token: tokens,
		Data: map[string]string{
			"title":    "Nuevo Mensaje",
			"entity":  string(messages),
			"type":     "999",
			"priority": "high",
		},
	}
	// arrayMessages := make([]*messaging.Message, 0,len(tokens))
	// for _,token := range tokens {
	// 	message := &messaging.Message{
	// 		//Notification: &messaging.Notification{
	// 			//	Title: "Notification Test",
	// 			//	Body:  "Hello React!!",
	// 			//},
	// 			Token: token,
	// 			Data: map[string]string{
	// 				"title":    "Nuevo Mensaje",
	// 				"subtext":  string(messages),
	// 				"type":     "999",
	// 				"priority": "high",
	// 			},
	// 			FCMOptions: &messaging.FCMOptions{},
	// 		}
	// 		arrayMessages = append(arrayMessages, message)
	// 	}
	// response, err := client.SendAll(ctx,arrayMessages)
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully sent message:", response)
}
