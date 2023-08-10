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
		err = u.SendNotificationUsersSala(ctx, val.Id)
		if err != nil {
			log.Println("ERROR",err)
			return
		}
	}
	return
}
func (u *salaUseCase) SendNotificationUsersSala(ctx context.Context, salaId int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, err := u.salaRepo.GetFcmTokensUserSalasSala(ctx, salaId)
	if err != nil {
		log.Println("FAILT TO FETCG TOKENS",err)
		return
	}
	message := r.MessageNotification{
		Message:  "Alguien mas reservo la cancha que se eligio para la sala.",
		EntityId: salaId,
	}
	data, err := json.Marshal(message)
	for _, val := range res {
		log.Println("FCM_TOKENS", val.FcmToken)
		u.utilU.SendNotification(ctx, val.FcmToken, data, r.NotificationSalaReservationConflict, u.firebase)
	}
	return
}
