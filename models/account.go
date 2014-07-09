package models

import (
	"time"
	"github.com/rebelnz/rego/db"
	"errors"
	"database/sql"
)

var ErrNoUsername = errors.New("Username doesnt exist")

type Account struct {
	Id        int64
	Username  string `sql:"size:255;not null"`
	Password  string `sql:"size:255;not null"`
	Firstname string `sql:"size:255"`
	Lastname  string `sql:"size:255"`
	Email     string `sql:"size:255;not null"`
	Priv      int64  `sql:"not null; default=1"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}


func GetAccountByUsername(username string) (Account, error) {
	a := Account{}
	err := db.DB.Where("username = ?", username).First(&a).Error
	if err != sql.ErrNoRows {
		return a, ErrNoUsername
	} else if err != nil {
		return a, err
	}
	return a, nil
}
