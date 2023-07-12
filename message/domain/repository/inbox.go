package repository

import "context"

type InboxRepository interface {
	SaveGrupoMessage(ctx context.Context, d *MessageGrupo) (err error)
}

type InboxUseCase interface {
	SaveGrupoMessage(ctx context.Context, d *MessageGrupo) (err error)
}


type Inbox struct {
	Id           int          `json:"id"`
	ProfileId    int          `json:"profile_id"`
	EstablecimientoId int `json:"establecimiento_id"`
	Content      string       `json:"content"`
	CreatedAt    string       `json:"created_at,omitempty"`
	ReplyTo      *int    `json:"reply_to"`
	ReplyMessage ReplyMessage `json:"reply_message"`
}