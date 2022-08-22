package database

import (
	"fmt"
	"time"

	"github.com/joaopedsa/dock-challenge/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseCli struct {
	db           *gorm.DB
	readTimeout  time.Duration
	writeTimeout time.Duration
}

var Conn *DatabaseCli

func GetReadTimeout() time.Duration {
	if Conn == nil {
		return 0
	}

	return Conn.readTimeout
}

func GetWriteTimeout() time.Duration {
	if Conn == nil {
		return 0
	}

	return Conn.writeTimeout
}

func Ping() error {
	if Conn == nil || Conn.db == nil {
		return ErrDBNil
	}

	db, err := Conn.db.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

func Close() error {
	if Conn == nil || Conn.db == nil {
		return ErrDBNil
	}

	db, err := Conn.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func GetDB() (*gorm.DB, error) {
	if Conn == nil || Conn.db == nil {
		return nil, ErrDBNil
	}

	return Conn.db, nil
}

func Connect(disableLogs bool) (*DatabaseCli, error) {
	var err error
	host := utils.GetEnv("DB_HOST", "0.0.0.0")
	port := utils.GetEnv("DB_PORT", "5432")
	dbName := utils.GetEnv("DB_NAME", "docks")
	username := utils.GetEnv("DB_USERNAME", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "123")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, username, password, dbName, port)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	_, err = gormDB.DB()
	if err != nil {
		return nil, err
	}
	Conn = &DatabaseCli{
		db:           gormDB,
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
	}
	return Conn, nil
}
