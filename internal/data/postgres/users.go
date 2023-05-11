package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/acs-dl/identity-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	usersTable     = "users"
	idColumn       = "id"
	nameColumn     = "name"
	surnameColumn  = "surname"
	positionColumn = "position"
	emailColumn    = "email"
	telegramColumn = "telegram"
)

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

func (q *usersQ) GetTotalCount() (int64, error) {
	stmt := sq.Select("COUNT (*)").From(usersTable)

	var count int64
	err := q.db.Get(&count, stmt)

	return count, err
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

func (q *usersQ) Select(selector data.UserSelector) ([]data.User, error) {
	return q.selectByQuery(applyUserSelector(q.sql, selector))
}

func (q *usersQ) selectByQuery(query sq.Sqlizer) ([]data.User, error) {
	var result []data.User

	err := q.db.Select(&result, query)

	return result, err
}

func applyUserSelector(sql sq.SelectBuilder, selector data.UserSelector) sq.SelectBuilder {
	if selector.Name != nil {
		sql = sql.Where(sq.ILike{nameColumn: "%" + *selector.Name + "%"})
	}
	if selector.Surname != nil {
		sql = sql.Where(sq.ILike{surnameColumn: "%" + *selector.Surname + "%"})
	}
	if selector.Position != nil {
		sql = sql.Where(sq.ILike{positionColumn: "%" + *selector.Position + "%"})
	}
	if selector.OffsetParams != nil {
		sql = selector.OffsetParams.ApplyTo(sql, fmt.Sprintf("CONCAT(%s, ' ', %s)", nameColumn, surnameColumn))
	}
	if selector.Email != nil {
		sql = sql.Where(sq.ILike{emailColumn: "%" + *selector.Email + "%"})
	}
	if selector.Search != nil {
		searchModified := strings.Replace(*selector.Search, " ", "%", -1)
		searchModified = fmt.Sprint("%", searchModified, "%")

		sql = sql.Where(
			sq.Or{
				sq.ILike{nameColumn: searchModified},
				sq.ILike{surnameColumn: searchModified},
				sq.ILike{surnameColumn + " || " + nameColumn: searchModified},
				sq.ILike{nameColumn + " || " + surnameColumn: searchModified},
			})
	}

	return sql
}

func (q *usersQ) GetById(id int64) (*data.User, error) {
	var result data.User

	err := q.db.Get(&result, q.sql.Where(sq.Eq{idColumn: id}))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, nil
}

func (q *usersQ) Delete(id int64) error {
	stmt := sq.Delete(usersTable).Where(sq.Eq{idColumn: id})
	err := q.db.Exec(stmt)

	return err
}

func (q *usersQ) Update(user data.User) error {
	clauses := structs.Map(user)

	stmt := sq.Update(usersTable).SetMap(clauses).Where(sq.Eq{idColumn: user.Id})
	err := q.db.Exec(stmt)

	return err
}

func (q *usersQ) UpdateTelegram(user data.User) error {
	stmt := sq.Update(usersTable)

	if user.Telegram == nil {
		stmt = stmt.Set(telegramColumn, user.Telegram)
	} else {
		stmt = stmt.Set(telegramColumn, *user.Telegram)
	}

	stmt = stmt.Where(sq.Eq{idColumn: user.Id})

	err := q.db.Exec(stmt)

	return err
}
