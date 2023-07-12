package repository

import "context"

type GrupoRepository interface {
	SaveGrupoMessage(ctx context.Context, d *MessageGrupo) (err error)
}


type GrupoUseCase interface {
	SaveGrupoMessage(ctx context.Context, d *MessageGrupo) (err error)
}

type MessageGrupo struct {
	Id           int          `json:"id"`
	GrupoId      int          `json:"grupo_id"`
	ProfileId    int          `json:"profile_id"`
	Content      string       `json:"content"`
	CreatedAt    string       `json:"created_at,omitempty"`
	ReplyTo      *int    `json:"reply_to"`
	ReplyMessage ReplyMessage `json:"reply_message"`
}
type ReplyMessage struct {
	Id        int    `json:"id"`
	GrupoId   int    `json:"grupo_id"`
	ProfileId int    `json:"profile_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	ReplyTo   *int   `json:"reply_to"`
}


type GrupoEvent struct {
	Type    string       `json:"type"`
	Message MessageGrupo `json:"message"`
	// Sala    SalaData     `json:"sala,omitempty"`
}
