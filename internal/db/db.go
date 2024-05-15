package db

import (
	"database/sql"
	"errors"
	"fmt"
	"integra_backend/internal/message"

	_ "github.com/lib/pq"
)

type dbConnection struct {
	conn *sql.DB
}

type DbConnection interface {
	GetConnection() *sql.DB
	CloseConnection() error
}

func NewDbConnection(host string, port int, user string, pwd string, dbname string) (DbConnection, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbname)
	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, errors.New(message.MsgConnectingPgDatabaseError)
	}
	return &dbConnection{conn: conn}, nil
}

func (db *dbConnection) GetConnection() *sql.DB {
	return db.conn
}

func (db *dbConnection) CloseConnection() error {
	return db.conn.Close()
}
