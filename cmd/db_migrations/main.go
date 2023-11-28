package main

// This is custom goose binary with sqlite3 support only.

import (
	"flag"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

var (
    flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir = "./cmd/db_migrations/migrations"
)

func main() {
    flags.Parse(os.Args[1:])
	args := flags.Args()
    command := "up"
    if len(args) == 1 {
        command = args[0]
    }
    
	dbString := os.Getenv("GOOSE_DBSTRING")
    logrus.Info(dbString)
	db, err := goose.OpenDBWithDriver("pgx", dbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.Run(command, db, dir); err != nil {
		log.Fatalf("goose: %v", err)
	}
}
