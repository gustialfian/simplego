package postgresql

import (
	"database/sql"
	"log"
)

// New return new *sql.DB instance
func New(con string) (*sql.DB, error) {
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("psql new: %v", err.Error())
		log.Println("db down")
		return nil, err
	}
	log.Println("db up")

	return db, nil
}

// Migration make table and seed the data
func Migration(db *sql.DB) error {
	log.Println("Migrating...")
	query := `
	create table if not exists users (
		id serial,
		username varchar(250),
		password varchar(250),
		age integer
	)`

	_, err := db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

// Seed the dummy data
func Seed(db *sql.DB) error {
	log.Println("Seeding...")
	query := `
	insert into users (username, password, age) 
	values ('foo', 'bar', '23')`

	_, err := db.Query(query)
	if err != nil {
		return err
	}

	return nil
}
