package repository

import (
	"context"
	"database/sql"
	"log"
	r "notification/domain/repository"

	// "github.com/lib/pq"
)

type messageRepository struct {
	Conn *sql.DB
}
func NewRepository(sql *sql.DB) r.MessageRepository {
	return &messageRepository{
		Conn: sql,
	}
}

func (p messageRepository) GetLastMessagesFromGroup(ctx context.Context, id int) (res []r.MessageGrupo,err error) {
	query := `select id,grupo_id,profile_id,content,created_at,reply_to from group_messages limit 3`
	res,err = p.fetchMessagesGrupo(ctx,query,id)
	return 
}

func (p messageRepository) GetUsersFromGroup(ctx context.Context, id int) (res []r.FcmToken,err error) {
	query := `select p.fcm_token from user_grupo as us 
	left join profiles as p on p.profile_id = us.profile_id where grupo_id = $1`
	log.Println("ID",id)
	res,err = p.fetchFcmTokens(ctx,query,id)
	if err != nil{
		log.Println("DEBUG_SQL",err)
	}
    return
}

func (p *messageRepository) fetchFcmTokens(ctx context.Context, query string, args ...interface{}) (res []r.FcmToken, err error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	res = make([]r.FcmToken, 0)
	for rows.Next() {
		t := r.FcmToken{}
		err = rows.Scan(
			&t.FcmToken,
		)
		res = append(res, t)
	}
	return res, nil
}

// func (p *messageRepository) fetchUserGrupo(ctx context.Context, query string, args ...interface{}) (res []r.ProfileUser, err error) {
// 	rows, err := p.Conn.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		errRow := rows.Close()
// 		if errRow != nil {
// 			log.Println(errRow)
// 		}
// 	}()
// 	res = make([]r.ProfileUser, 0)
// 	for rows.Next() {
// 		t := r.ProfileUser{}
// 		err = rows.Scan(
// 			&t.Nombre,
// 			&t.Apellido,
// 			&t.ProfilePhoto,
// 			&t.UserGrupoId,
// 			&t.FcmToken,
// 		)
// 		res = append(res, t)
// 	}
// 	return res, nil
// }

func (p *messageRepository) fetchMessagesGrupo(ctx context.Context, query string, args ...interface{}) (res []r.MessageGrupo, err error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	res = make([]r.MessageGrupo, 0)
	for rows.Next() {
		t := r.MessageGrupo{}
		err = rows.Scan(
			&t.Id,
			&t.GrupoId,
			&t.ProfileId,
			&t.Content,
			&t.CreatedAt,
			&t.ReplyTo,
		)
		res = append(res, t)
	}
	return res, nil
}

