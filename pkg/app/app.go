package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustialfian/simplego/pkg/postgresql"
	"github.com/gustialfian/simplego/pkg/users"
)

// RegisterRouter list of route
func RegisterRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "safe\n")
		return
	})

	apiRoute := r.PathPrefix("/api").Subrouter()
	amw := authenticationMiddleware{}
	amw.Populate()
	apiRoute.Use(amw.Middleware)
	apiRoute.Handle("/user", users.FindAllHandler(db)).Methods("GET")
	apiRoute.Handle("/user", users.InsertHandler(db)).Methods("POST")
	apiRoute.Handle("/user/{id}", users.FindByIDHandler(db)).Methods("GET")
	apiRoute.Handle("/user/{id}", users.UpdateHandler(db)).Methods("PUT")
	apiRoute.Handle("/user/{id}", users.DeleteHandler(db)).Methods("DELETE")

	return r
}

// RegisterDB register database service
func RegisterDB(con string) (*sql.DB, error) {
	db, err := postgresql.New(con)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := postgresql.Migration(db); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := postgresql.Seed(db); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return db, nil
}
