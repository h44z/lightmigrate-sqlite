package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/h44z/lightmigrate"
	"github.com/h44z/lightmigrate-sqlite/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlClient, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		log.Fatalf("unable to setup sql client: %v", err)
	}
	defer sqlClient.Close()

	fsys := os.DirFS("examples")
	source, err := lightmigrate.NewFsSource(fsys, "migrations")
	if err != nil {
		log.Fatalf("unable to setup source: %v", err)
	}
	defer source.Close()

	driver, err := sqlite.NewDriver(sqlClient, "migration_test_db")
	if err != nil {
		log.Fatalf("unable to setup driver: %v", err)
	}
	defer driver.Close()

	migrator, err := lightmigrate.NewMigrator(source, driver, lightmigrate.WithVerboseLogging(true))
	if err != nil {
		log.Fatalf("unable to setup migrator: %v", err)
	}

	err = migrator.Migrate(1) // Migrate to schema version 1
	if err != nil {
		log.Fatalf("migration error: %v", err)
	}

	// check result
	rows, err := sqlClient.Query("SELECT sql FROM sqlite_master")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Table Schemas:")
	for rows.Next() {
		var sqlStr sql.NullString
		err = rows.Scan(&sqlStr)
		if err != nil {
			log.Fatal(err)
		}
		if sqlStr.Valid {
			fmt.Println(sqlStr.String)
		}
	}
}
