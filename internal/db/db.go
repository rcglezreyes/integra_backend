package db

import (
	"database/sql"
	"errors"
	"fmt"

	c "integra_backend/internal/constant"

	_ "github.com/lib/pq"
)

type dbConnection struct {
	conn *sql.DB
}

type DbConnection interface {
	GetConnection() *sql.DB
	CloseConnection() error
}

func NewDbConnection() (DbConnection, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)
	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, errors.New("error connecting to Postgres DB")
	}
	return &dbConnection{conn: conn}, nil
}

func (db *dbConnection) GetConnection() *sql.DB {
	return db.conn
}

func (db *dbConnection) CloseConnection() error {
	return db.conn.Close()
}
