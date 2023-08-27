package main

import (
	"database/sql"
	"log"

	"github.com/MohammadZeyaAhmad/Bank-App/api"
	db "github.com/MohammadZeyaAhmad/Bank-App/db/sqlc"
	"github.com/MohammadZeyaAhmad/Bank-App/util"
	_ "github.com/lib/pq"
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

	err = server.Start(config.Port)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
  
}