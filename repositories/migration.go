package repositories

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"

	"github.com/pressly/goose"
)

const (
	driver       = "postgres"
	migrationDir = "./migrations/sql"
)

// command helper
func usage() {
	const (
		usageRun      = `goose [OPTIONS] COMMAND`
		usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version`
	)
	fmt.Println(usageRun)
	flag.PrintDefaults()
	fmt.Println(usageCommands)
}

// Migrate the changes to database
func Migrate() error {
	flag.Usage = usage

	flag.Parse()
	args := flag.Args()

	dbString := NewConnectionString()
	db, err := sql.Open(driver, dbString)
	if err != nil {
		return errors.New("unable to open database")
	}

	defer db.Close()

	if err = goose.SetDialect(driver); err != nil {
		return fmt.Errorf("failed to set goose dialect: %v", err)
	}

	if len(args) == 0 {
		return errors.New("expected at least one arg")
	}

	command := args[0]

	if err = goose.Run(command, db, migrationDir, args[1:]...); err != nil {
		return fmt.Errorf("goose run: %v", err)
	}
	return db.Close()
}
