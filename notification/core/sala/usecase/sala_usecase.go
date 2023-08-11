package usecase

import (
	"context"
	"encoding/json"
	// "fmt"
	"log"
	r "notification/domain/repository"
	"time"

	firebase "firebase.google.com/go"
)

type salaUseCase struct {
	salaRepo r.SalaRepository
	firebase *firebase.App
	timeout  time.Duration
	utilU    r.UtilUseCase
}

func NewUseCase(salaRepo r.SalaRepository, firebase *firebase.App, timeout time.Duration, utilU r.UtilUseCase) r.SalaUseCase {
	return &salaUseCase{
		salaRepo: salaRepo,
		timeout:  timeout,
		firebase: firebase,
		utilU:    utilU,
	}
}

func (u *salaUseCase) SalaHasBennReserved(ctx context.Context,d []byte) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	var data r.Sala
	err = json.Unmarshal(d, &data)
	if err != nil {
		log.Println(err)
		return
	}
	message := r.MessageNotification{
		Message:  "La reserva para la sala se ha completado. ¡Prepárate para jugar!",
		EntityId: data.Id,
	}
	log.Println(message)
	err = u.SendNotificationUsersSala(ctx,message,r.NotificationSalaHasBeenReserved)
	if err != nil {
		log.Println("ERROR",err)
	}
	return
}

func (u *salaUseCase) SalaReservationConflict(ctx context.Context,d []byte) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	var data r.SalaConflictData
	err = json.Unmarshal(d, &data)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("IDDDDDD",data.SalaIds)
	for _, val := range data.SalaIds {
		log.Println("IDDDDDD",val.Id)
		message := r.MessageNotification{
			Message:  "Lamentamos informarte que alguien más ha reservado la cancha que habías seleccionado para la sala.",
			EntityId: val.Id,
		}
		err = u.SendNotificationUsersSala(ctx,message,r.NotificationSalaReservationConflict)
		if err != nil {
			log.Println("ERROR",err)
		}
		err = u.salaRepo.DeleteSala(ctx,val.Id)
		if err != nil{
			log.Println()
			return 
		}
	}
	return
}
func (u *salaUseCase) SendNotificationUsersSala(ctx context.Context,message r.MessageNotification,
	notification  r.NotificationType) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, err := u.salaRepo.GetFcmTokensUserSalasSala(ctx, message.EntityId)
	if err != nil {
		log.Println("FAILT TO FETCG TOKENS",err)
		return
	}
	data, err := json.Marshal(message)
	for _, val := range res {
		log.Println("FCM_TOKENS", val.FcmToken)
		u.utilU.SendNotification(ctx, val.FcmToken, data,notification, u.firebase)
	}
	return
}
