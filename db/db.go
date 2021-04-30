package db

import (
	"2gis/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error
	if db == nil {
		var configuration = config.GetConfig()
		db, err = sql.Open(
			"postgres",
			fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
				configuration.Database.User,
				configuration.Database.Password,
				configuration.Database.Host,
				configuration.Database.Port,
				configuration.Database.Name,
			),
		)
		if err != nil {
			panic(err)
		}
	}

	return db
}