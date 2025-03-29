package main

import (
	"database/sql"
	"log"

	"github.com/husky_dusky/simplebank/api"
	db "github.com/husky_dusky/simplebank/db/sqlc"
	"github.com/husky_dusky/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(cfg, store)

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
