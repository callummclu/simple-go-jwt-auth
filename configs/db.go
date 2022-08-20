package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	host, err := getEnvByName("DB_HOST")

	if err != nil {
		panic(err)
	}

	port, err := getEnvByName("DB_PORT")

	if err != nil {
		panic(err)
	}

	user, err := getEnvByName("DB_USER")

	if err != nil {
		panic(err)
	}

	password, err := getEnvByName("DB_PASS")

	if err != nil {
		panic(err)
	}

	dbname, err := getEnvByName("DB_NAME")

	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	return db, err
}
