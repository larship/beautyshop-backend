package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/larship/barbershop/config"
	"log"
)

type Database struct {
	config *config.Config
	conn   *pgx.Conn
}

var DB *Database

func Init(conf *config.Config) *Database {
	DB = &Database{
		config: conf,
	}

	return DB
}

func (db *Database) GetConnection() *pgx.Conn {
	var err error

	if db.conn == nil || db.conn.Ping(context.Background()) != nil {
		log.Print("Создаём новое подключение к БД")
		db.conn, err = pgx.Connect(context.Background(), db.config.DatabaseDsn)
		if err != nil {
			log.Fatalf("Невозможно подключиться к БД: %v\n", err)
		}
	}

	return db.conn
}

func (db *Database) CloseConnection() {
	if db.conn != nil {
		log.Println("Закрываем подключение к БД")
		db.conn.Close(context.Background())
	}
}
