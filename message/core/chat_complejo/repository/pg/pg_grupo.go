package pg

import (
	"context"
	"database/sql"
	// "log"
	r "message/domain/repository"
)

type inboxRepo struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) r.InboxRepository {
	return &inboxRepo{
		Conn: conn,
	}
}


func (p *inboxRepo) SaveGrupoMessage(ctx context.Context, d *r.MessageGrupo) (err error) {
	query := `insert into conversation_message (conversation_id,sender_id,content,created_at,reply_to) 
	values($1,$2,$3,now(),$4) returning id,created_at`
	err = p.Conn.QueryRowContext(ctx, query, d.GrupoId, d.ProfileId, d.Content, d.ReplyTo).Scan(&d.Id, &d.CreatedAt)
	return
}





// func (m *inboxRepo) fetchMessagesGrupo(ctx context.Context, query string, args ...interface{}) (res []r.MessageGrupo, err error) {
// 	rows, err := m.Conn.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		errRow := rows.Close()
// 		if errRow != nil {
// 			log.Println(errRow)
// 		}
// 	}()
// 	res = make([]r.MessageGrupo, 0)
// 	for rows.Next() {
// 		t := r.MessageGrupo{}
// 		err = rows.Scan(
// 			&t.Id,
// 			&t.GrupoId,
// 			&t.ProfileId,
// 			&t.Content,
// 			&t.CreatedAt,
// 			&t.ReplyTo,
// 			&t.ReplyMessage.Id,
// 			&t.ReplyMessage.GrupoId,
// 			&t.ReplyMessage.ProfileId,
// 			&t.ReplyMessage.Content,
// 			&t.ReplyMessage.CreatedAt,
// 		)
// 		res = append(res, t)
// 	}
// 	return res, nil
// }
