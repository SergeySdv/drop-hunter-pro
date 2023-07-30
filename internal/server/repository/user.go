package repository

import (
	"context"
	"database/sql"

	"github.com/hardstylez72/cry/internal/server/repository/pg"
)

type User struct {
	Id     string `db:"id"`
	Email  string `db:"email"`
	Access bool   `db:"access"`
}

func (r *pgRepository) GetOrCreateUser(ctx context.Context, user *User) (*User, bool, error) {

	q := `select id, email, access from users where email = $1`

	var u User
	if err := r.conn.GetContext(ctx, &u, q, user.Email); err != nil {
		if err == sql.ErrNoRows {
		} else {
			return nil, false, pg.PgError(err)
		}
	} else {
		return &u, true, nil
	}

	q = `insert into users (id, email, access) values (:id, :email, :access) on conflict (email) DO NOTHING`

	if _, err := r.conn.NamedExecContext(ctx, q, user); err != nil {
		return nil, false, pg.PgError(err)
	}

	q = `select id, email, access from users where email = $1`

	if err := r.conn.GetContext(ctx, &u, q, user.Email); err != nil {
		return nil, false, pg.PgError(err)
	}

	return &u, false, nil
}

func (r *pgRepository) GetUserById(ctx context.Context, id string) (*User, error) {
	q := `select id, email, access from users where id = $1`

	var u User
	if err := r.conn.GetContext(ctx, &u, q, id); err != nil {
		return nil, pg.PgError(err)
	}

	return &u, nil
}

func (r *pgRepository) SubscribeAlerts(ctx context.Context, email, chatId string) error {
	q := `select id from users where email = $1`
	var userId string
	if err := r.conn.GetContext(ctx, &userId, q, email); err != nil {
		return pg.PgError(err)
	}
	q = `insert into alert_subscriptions (user_id, telegram_chat_id) values ($1, $2) on conflict (user_id) do nothing `
	if _, err := r.conn.ExecContext(ctx, q, userId, chatId); err != nil {
		return err
	}
	return nil
}
func (r *pgRepository) GetUserTelegramChatId(ctx context.Context, userId string) (*string, error) {
	q := `select telegram_chat_id from alert_subscriptions where user_id = $1`
	var chatId string
	if err := r.conn.GetContext(ctx, &chatId, q, userId); err != nil {
		return nil, pg.PgError(err)
	}
	return &chatId, nil
}
