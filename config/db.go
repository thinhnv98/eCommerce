package config

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
}

func (_self Database) InitDatabase() *sql.DB {
	var (
		driver   = os.Getenv("POSTGRES_DRIVER")
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PWD")
		dbname   = os.Getenv("POSTGRES_DBNAME")
	)

	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)
	fmt.Println(psqlInfo)
	// Open db connection
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}