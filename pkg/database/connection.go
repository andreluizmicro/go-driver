package database

import (
	"database/sql"
	"fmt"

	"github.com/andreluizmicro/go-driver-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(cfg *configs.AppConfig) (*sql.DB, error) {
	strConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open(cfg.DBDriver, strConnection)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
