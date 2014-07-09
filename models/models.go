package models

import (
	"github.com/rebelnz/rego/db"
	"log"
	"os"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

// page meta
type Page struct {
	UserID int64
	Title string
	Account interface{}
	CSS []string
	JS []string
	Message []string
}

func init() {
	db.DB.SingularTable(true)
	db.DB.AutoMigrate(Account{})
	log.Println("Tables created")
}
