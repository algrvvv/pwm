package main

import (
	"database/sql"
	"errors"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"github.com/algrvvv/pwm/utils"
)

var (
	down = flag.Bool("down", false, "down migrations")
	up   = flag.Bool("up", false, "down migrations")
)

func main() {
	flag.Parse()
	if (*up && *down) || (!*up && !*down) {
		log.Fatalf("invalid flags")
	}

	path, err := utils.GetDBPath()
	if err != nil {
		log.Fatalf("failed to get database path: %v", err)
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}
	defer db.Close()

	log.Println("successfully connected to database")

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("failed to initialize driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrate/migrations/",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to initialize migrations: %s", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			log.Printf("failed to close src migrations: %v\n", srcErr)
		}
		if dbErr != nil {
			log.Printf("failed to close db migrations: %v\n", dbErr)
		}

		if srcErr == nil && dbErr == nil {
			log.Println("successfully closed migrations")
		}
	}()

	if v, dirty, err := m.Version(); err == nil && dirty {
		log.Printf("last version: %v is dirty\n", v)

		if err = m.Force(int(v)); err != nil {
			log.Fatalf("failed to force: %v", err)
		}

		if err = m.Steps(-1); err != nil {
			log.Fatalf("failed to rollback migrations: %v", err)
		}

		log.Println("successfully cleaning dirty migrations")
	} else if err != nil {
		if !errors.Is(err, migrate.ErrNilVersion) {
			log.Fatalf("failed to get mirgations version: %v", err)
		}
	}

	if *up {
		log.Println("start up migrations")
		if err = m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("Migrations have not changes")
			} else {
				log.Fatalf("failed to run migrations: %v", err)
			}
		} else {
			log.Println("successfully applied migrations")
		}
	} else if *down {
		log.Println("start down migrations")
		if err = m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("Migrations have not changes")
			} else {
				log.Fatalf("failed to run migrations: %v", err)
			}
		} else {
			log.Println("successfully applied migrations")
		}
	}
}
