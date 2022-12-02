package data

import "gitlab.com/distributed_lab/kit/pgdb"

type UsersQ interface {
	New() UsersQ

	Create(user User) (int64, error)
	Select() ([]User, error)
	Get() (*User, error)
	Delete(id int64) error
	Update(user User) error

	FilterById(id int64) UsersQ
	FilterByName(name string) UsersQ
	FilterBySurname(surname string) UsersQ
	Page(pageParams pgdb.OffsetPageParams) UsersQ
}

type User struct {
	Id       int64  `db:"id" structs:"-"`
	Name     string `db:"name" structs:"name"`
	Surname  string `db:"surname" structs:"surname"`
	Position string `db:"position" structs:"position"`
}
