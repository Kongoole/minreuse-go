package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Model struct {
	Master *sql.DB
	Slave  *sql.DB
}

var conn *sql.DB

func init() {
	// connect to database
	dsn := os.Getenv("DB_USERNAME")+":" + os.Getenv("DB_PASSWORD") + "@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_DATABASE")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("fail to connect to mysql")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	conn = db
}

// with a Model pointer
func (m *Model) InitMaster() {
	m.Master = conn
}

func (m *Model) InitSlave() {
	// TODO: build slave connection
	m.Slave = conn
}
