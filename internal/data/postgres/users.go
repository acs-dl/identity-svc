package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const usersTable = "action"

var selectUsersTable = sq.Select("*").From(usersTable)

type usersQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func NewUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db:  db.Clone(),
		sql: selectUsersTable,
	}
}

func (q *usersQ) New() data.UsersQ {
	return NewUsersQ(q.db)
}

func (q *usersQ) Create(user data.User) (int64, error) {
	clauses := structs.Map(user)

	var id int64

	stmt := sq.Insert(usersTable).SetMap(clauses).Suffix("returning id")
	err := q.db.Get(&id, stmt)

	return id, err
}

func (q *usersQ) Select() ([]data.User, error) {
	var result []data.User
	err := q.db.Select(&result, q.sql)

	return result, err
}

func (q *usersQ) Get() (*data.User, error) {
	var result data.User

	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, nil
}

func (q *usersQ) Delete(id int64) error {
	stmt := sq.Delete(usersTable).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)

	return err
}

func (q *usersQ) Update(user data.User) error {
	clauses := structs.Map(user)

	stmt := sq.Update(usersTable).SetMap(clauses).Where(sq.Eq{"id": user.Id})
	err := q.db.Exec(stmt)

	return err
}

func (q *usersQ) FilterById(id int64) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{"id": id})
	return q
}

func (q *usersQ) FilterByName(name string) data.UsersQ {
	q.sql = q.sql.Where(sq.ILike{"name": "%" + name + "%"})
	return q
}

func (q *usersQ) FilterBySurname(surname string) data.UsersQ {
	q.sql = q.sql.Where(sq.ILike{"surname": "%" + surname + "%"})
	return q
}

func (q *usersQ) Page(pageParams pgdb.OffsetPageParams) data.UsersQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}
