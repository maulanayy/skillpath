package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"skillpath/model"
	"skillpath/pkg"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	filePath := "./config.yaml"

	var cfg model.Config
	err := pkg.ReadFile(&cfg, filePath)
	if err != nil {
		fmt.Println("error read file")
		log.Fatal(err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DB.Address, cfg.DB.Port),
		Database: cfg.DB.Name,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
	})
	defer db.Close()
	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
