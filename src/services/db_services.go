package services

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func readConfig() string {
	DbUsername := os.Getenv("POSTGRES_USER")
	DbPassword := os.Getenv("POSTGRES_PASSWORD")
	DbName := os.Getenv("POSTGRES_DB")

	config := fmt.Sprintf("host=test_avito_postgres port=5432 user=%s password=%s dbname=%s sslmode=disable",
		                  DbUsername, DbPassword, DbName)

	return config
}

func OpenDB() (*sql.DB, error) {
	conf := readConfig()
	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Println("Open database error: ", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Check database connection error: ", err)
		return nil, err
	}
	log.Println("Successful connection to database.")
	return db, nil
}

func Setup(filename string, db *sql.DB) error {
	query, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Read .sql file error: ", err)
		return err
	}
	if _, err = db.Exec(string(query)); err != nil {
		log.Println("Setup database error: ", err)
		return err
	}
	return nil
}

