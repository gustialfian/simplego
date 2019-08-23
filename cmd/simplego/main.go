package main

import (
	"log"
	"net/http"

	"github.com/gustialfian/simplego/pkg/app"
	"github.com/gustialfian/simplego/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.New()

	db, err := app.RegisterDB(cfg.ConnectionDB)
	defer db.Close()
	if err != nil {
		return
	}
	r := app.RegisterRouter(db)

	log.Printf("listen at port %v\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, r))
}
