package data

import "gitlab.com/distributed_lab/kit/pgdb"

type UsersQ interface {
	New() UsersQ

	Create(user User) (int64, error)
	Delete(id int64) error
	Update(user User) error
	GetById(id int64) (*User, error)

	GetTotalCount() (int64, error)

	Select(selector UserSelector) ([]User, error)
}

type User struct {
	Id       int64  `db:"id" structs:"-"`
	Name     string `db:"name" structs:"name"`
	Surname  string `db:"surname" structs:"surname"`
	Email    string `db:"email" structs:"email"`
	Position string `db:"position" structs:"position"`
}

// UserSelector is a structure for all applicable filters and params
type UserSelector struct {
	OffsetParams *pgdb.OffsetPageParams
	Name         *string
	Surname      *string
	Position     *string
	Email        *string
	Search       *string
}
