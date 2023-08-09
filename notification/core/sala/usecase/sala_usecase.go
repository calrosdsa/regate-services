package usecase

import (
	"context"
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

func NewUseCase(salaRepo r.SalaRepository, firebase *firebase.App, timeout time.Duration,utilU r.UtilUseCase) r.SalaUseCase {
	return &salaUseCase{
		salaRepo: salaRepo,
		timeout:  timeout,
		firebase: firebase,
		utilU: utilU,
	}
}

func (u *salaUseCase) HasBeenReserved(ctx context.Context, d r.SalaHasBeenReserved) (err error) {
	return
}
func (u *salaUseCase) SendNotificationUsersSala(ctx context.Context, salaId int) (err error) {

	return
}
