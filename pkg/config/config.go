package config

import (
	"flag"
	"fmt"
	"log"
)

// Config list configuration for this app
type Config struct {
	Port         string
	ConnectionDB string
}

// New setup config
func New() *Config {
	cfg := Config{}
	log.Println("RegisterConfig")

	defaultConnectionDB := `
	user=postgres
	password=mysecretpassword
	host=localhost
	port=8080
	dbname=sandbox
	sslmode=disable`
	flag.StringVar(&cfg.ConnectionDB, "db", defaultConnectionDB, "database connection string")

	var port string
	defaultPort := ":8000"
	flag.StringVar(&port, "port", defaultPort, "port for web server")
	flag.Parse()

	cfg.Port = fmt.Sprintf(":%v", port)

	return &cfg
}
