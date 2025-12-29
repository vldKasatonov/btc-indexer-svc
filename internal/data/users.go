package data

import "time"

type UsersQ interface {
	New() UsersQ

	Insert(data User) (User, error)
	Get() (*User, error)

	FilterByUsername(username string) UsersQ
}

type User struct {
	ID           int64     `db:"id" structs:"-"`
	Username     string    `db:"username" structs:"username"`
	PasswordHash string    `db:"password_hash" structs:"password_hash"`
	CreatedAt    time.Time `db:"created_at" structs:"-"`
}
