package repository

import "context"

type SalaPayload struct {
	Id      int    `json:"id"`
	Titulo  string `json:"titulo"`
	GrupoId int    `json:"grupo_id"`
}

type MessageNotification struct {
	// Title    string `json:"title"`
	Message  string `json:"message"`
	EntityId int    `json:"id"`
}

type SalaConflictData struct {
	SalaIds []Ids `json:"sala_ids"`
}

type Sala struct {
	Id int `json:"id"`
}

type Ids struct {
	Id int `json:"id"`
}

type SalaUseCase interface {
	//Enviar noticaciones a todos los usuarios de las salas donde se ha echo la reserva
	//Para notificar que la cancha que querian notificar ya fue reservada por otro
	SalaReservationConflict(ctx context.Context, d []byte) (err error)
	SendNotificationUsersSala(ctx context.Context, message MessageNotification, notification NotificationType) (err error)
	SalaHasBennReserved(ctx context.Context, d []byte) (err error)
}

type SalaRepository interface {
	GetFcmTokensUserSalasSala(ctx context.Context, salaId int) ([]UserSalaFcmToken, error)
	DeleteSala(ctx context.Context, salaId int) (err error)
}

type UserSalaFcmToken struct {
	FcmToken  string
	ProfileId int
	Amount    float64
}
