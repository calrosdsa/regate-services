package repository

import "context"

type MessageRepository interface {
	GetLastMessagesFromGroup(ctx context.Context, id int) ([]MessageGrupo, error)
	GetUsersFromGroup(ctx context.Context, id int) ([]ProfileUser, error)
}

type MessageUseCase interface {
	GetLastMessagesFromGroup(ctx context.Context, id int) ([]MessageGrupo, error)
	GetUsersFromGroup(ctx context.Context, id int) ([]ProfileUser, error)
}
type ProfileUser struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	// Genero   string     `json:"genero"`
	// BirthDay *time.Time `json:"birthDay"`
	ProfilePhoto *string `json:"profile_photo"`
	//only for user grupo table
	UserGrupoId int `json:"user_grupo_id,omitempty"`
}

type MessageGrupo struct {
	Id        int    `json:"id"`
	GrupoId   int    `json:"grupo_id"`
	ProfileId int    `json:"profile_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	ReplyTo   *int   `json:"reply_to"`
}
