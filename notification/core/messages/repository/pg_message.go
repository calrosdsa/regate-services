package repository

import (
	"context"
	"database/sql"
	"notification/domain/repository"
)

type messageRepository struct {
	Conn *sql.DB
}

func (m messageRepository) GetLastMessagesFromGroup(ctx context.Context, id int) ([]repository.MessageGrupo, error) {
	//TODO implement me
	panic("implement me")
}

func (m messageRepository) GetUsersFromGroup(ctx context.Context, id int) ([]repository.ProfileUser, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(sql *sql.DB) repository.MessageRepository {
	return &messageRepository{
		Conn: sql,
	}
}
