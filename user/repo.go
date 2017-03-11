package user

import (
	"database/sql"
	"fmt"
	"github.com/Sirupsen/logrus"
)

type UserRepo interface {
	Create(*User) error
}

type User struct {
	UserID   string
	Provider string
	Uid      int
	Token    string
	Ttl      int
}

type userRepo struct {
	db: *sql.DB
}

func (r *userRepo) Create(u *User) error {
	tx := r.db.Begin()
	defer tx.Commit()
	
	_, err := tx.Exec(
		"INSERT kv_user (user_id, provider, uid, token, ttl) value (?, ?, ?, ?, ?)",
		u.UserID,
		u.Provider,
		u.Uid,
		u.Token,
		u.Ttl,
	)
	if err != nil {
		fmt.Errorf("Could not create user %s", err)
		tx.Rollback
		return err
	}

	return nil
}
