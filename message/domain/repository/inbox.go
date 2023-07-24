package repository

import "context"

type InboxRepository interface {
	SaveGrupoMessage(ctx context.Context, d *MessageGrupo) (err error)
	GetMessages(ctx context.Context, id int) ([]Inbox, error)
}

type InboxUseCase interface {
	SaveGrupoMessage(ctx context.Context, d *MessageGrupo) (err error)
	GetMessages(ctx context.Context, id int) ([]Inbox, error)
}

type Inbox struct {
	Id                int        `json:"id"`
	SenderId          int        `json:"sender_id"`
	EstablecimientoId int        `json:"establecimiento_id"`
	Content           string     `json:"content"`
	CreatedAt         string     `json:"created_at,omitempty"`
	ReplyTo           *int       `json:"reply_to"`
	Reply             ReplyInbox `json:"reply"`
}

type ReplyInbox struct {
	Id                int    `json:"id"`
	SenderId          int    `json:"sender_id"`
	EstablecimientoId int    `json:"establecimiento_id"`
	Content           string `json:"content"`
	CreatedAt         string `json:"created_at,omitempty"`
	ReplyTo           *int   `json:"reply_to"`
}
