package user

import (
	"database/sql"
	"fmt"
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
	db *sql.DB
}

func (r *userRepo) Create(u *User) error {
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Errorf("Tx user fail %s", err)
		return err
	}
	defer tx.Commit()

	_, err = tx.Exec(
		"INSERT kv_user (user_id, provider, uid, token, ttl) value (?, ?, ?, ?, ?)",
		u.UserID,
		u.Provider,
		u.Uid,
		u.Token,
		u.Ttl,
	)
	if err != nil {
		fmt.Errorf("Could not create user %s", err)
		if err := tx.Rollback(); err != nil {
			fmt.Errorf("Rollback user fail %s", err)
		}
		return err
	}

	return nil
}

func (r *userRepo) Update(u *User) error {
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Errorf("Tx fail %s", err)
		return err
	}
	defer tx.Commit()
	_, err = tx.Exec(
		"UPDATE kv_user SET token = ?, ttl = ? WHERE user_id = ?",
		u.Token,
		u.Ttl,
		u.UserID,
	)
	if err != nil {
		fmt.Errorf("Could not update user %s", err)
		if err := tx.Rollback(); err != nil {
			fmt.Errorf("Rollback user fail %s", err)
		}
		return err
	}
	return nil
}
