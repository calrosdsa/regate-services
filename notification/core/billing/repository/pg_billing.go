package repository

import (
	"context"
	"database/sql"
	"log"
	r "notification/domain/repository"
)

type salaRepo struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.SalaRepository {
	return &salaRepo{
		Conn: conn,
	}
}

func (p *salaRepo) GetFcmTokensUserSalasSala(ctx context.Context,salaId int)(res []r.FcmToken,err error){
	query := `select p.fcm_token from users_sala as us
	inner join profiles as p on p.profile_id = us.profile_id
	where sala_id = $1`
	res,err = p.fetchFcmTokens(ctx,query,salaId)
	return
}

func (p *salaRepo) DeleteSala(ctx context.Context,salaId int)(err error){
	query := `delete from salas where estado = $1 and sala_id = $2`
	_,err =  p.Conn.ExecContext(ctx,query,r.SalaUnAvailable,salaId)
	return
}


func (p *salaRepo) fetchFcmTokens(ctx context.Context, query string, args ...interface{}) (res []r.FcmToken, err error) {
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