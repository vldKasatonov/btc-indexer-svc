package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/vldKasatonov/btc-indexer-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const usersTableName = "users"

var UsersSelector = sq.Select(usersTableName + ".*").From(usersTableName)

func NewUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db:  db,
		sql: UsersSelector,
	}
}

type usersQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *usersQ) New() data.UsersQ {
	return NewUsersQ(q.db.Clone())
}

func (q *usersQ) Insert(value data.User) (data.User, error) {
	var result data.User
	clauses := structs.Map(value)
	stmt := sq.Insert(usersTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	return result, err
}

func (q *usersQ) Get() (*data.User, error) {
	var result data.User
	err := q.db.Get(&result, q.sql)
	return &result, err
}

func (q *usersQ) FilterByUsername(username string) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{usersTableName + ".username": username})
	return q
}
