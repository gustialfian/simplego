package users

import (
	"database/sql"
	"log"
)

func findAll(db *sql.DB) ([]User, error) {
	query := `select * from users order by id asc`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var cur User

		if err := rows.Scan(&cur.ID, &cur.Username, &cur.Password, &cur.Age); err != nil {
			log.Fatal(err.Error())
			return nil, err
		}

		users = append(users, cur)
	}

	return users, nil
}

func findByID(db *sql.DB, id string) (User, error) {
	query := `select * from users where users.id = $1;`

	var user User
	err := db.QueryRow(query, id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Age)
	if err != nil {
		log.Fatal(err.Error())
		return User{}, err
	}

	return user, nil
}

func insert(db *sql.DB, u User) (User, error) {
	query := `insert into users (username, password, age)
		values ($1, $2, $3) returning *;`

	var user User
	err := db.QueryRow(query, u.Username, u.Password, u.Age).
		Scan(&user.ID, &user.Username, &user.Password, &user.Age)
	if err != nil {
		log.Fatal(err.Error())
		return User{}, err
	}

	return user, nil
}

func update(db *sql.DB, u User) (User, error) {
	query := `
	update users
	set username=$1, 
			password=$2, 
			age=$3
	where id=$4
	returning *;`

	var user User
	err := db.QueryRow(query, u.Username, u.Password, u.Age, u.ID).
		Scan(&user.ID, &user.Username, &user.Password, &user.Age)
	if err != nil {
		log.Fatal(err.Error())
		return User{}, err
	}

	return user, nil
}

func delete(db *sql.DB, id string) (User, error) {
	query := `delete from users where id=$1 returning *;`

	var user User
	err := db.QueryRow(query, id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Age)
	if err != nil {
		log.Fatal(err.Error())
		return User{}, err
	}

	return user, nil
}
