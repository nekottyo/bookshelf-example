package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'nickname'"`
}

var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:hogehoge@localhost/bookshelf")
	if err != nil {
		panic(err)
	}
}

func NewUser(id int, username string) User {
	return User{
		ID:       id,
		Username: username,
	}
}

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (m UserRepository) GetByID(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}
	return nil
}
