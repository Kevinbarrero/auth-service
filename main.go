package main

import (
	"database/sql"
	"log"

	"github.com/kevinbarrero/auth-service/api"
	db "github.com/kevinbarrero/auth-service/db/sqlc"

	"github.com/kevinbarrero/auth-service/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

}