package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// FindAllHandler return handlefunc
func FindAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := findAll(db)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}

		json.NewEncoder(w).Encode(users)
		return
	}
}

// FindByIDHandler find specific user that match this id
func FindByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user, err := findByID(db, vars["id"])
		if err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}
}

// InsertHandler insert new user
func InsertHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(b, &u); err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}

		user, err := insert(db, u)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}
}

// UpdateHandler update specific user from id
func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(b, &u); err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}

		vars := mux.Vars(r)
		u.ID = vars["id"]

		user, err := update(db, u)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}
}

// DeleteHandler delete spcific user by id
func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user, err := delete(db, vars["id"])
		if err != nil {
			fmt.Fprintf(w, "%v\n", err.Error())
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}
}
