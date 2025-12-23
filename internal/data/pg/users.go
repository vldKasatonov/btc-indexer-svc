package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/vldKasatonov/btc-indexer-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const usersTableName = "users"

func NewUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db: db,
	}
}

type usersQ struct {
	db *pgdb.DB
}

func (q *usersQ) New() data.UsersQ {
	return NewUsersQ(q.db)
}

func (q *usersQ) Insert(value data.User) (data.User, error) {
	var result data.User
	clauses := structs.Map(value)
	stmt := sq.Insert(usersTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	return result, err
}
