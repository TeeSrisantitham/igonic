package main

import (
	"fmt"
	"os"

	"github.com/openware/igonic/config"
)

var cfg config.Config

func create() {
	db := config.ConnectDatabase()
	db = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`;", cfg.Database.Name))
}

func migrate() {
	db := config.ConnectDatabase()
	config.RunMigrations(db)
}

func seed() {
	db := config.ConnectDatabase()
	config.LoadSeeds(db)
}

func usage() {
	fmt.Println(`
Usage: operator

db:create		Create database
db:migrate		Migrate database
db:seed			Seed database`)
	os.Exit(1)
}

func main() {

	config.Parse(&cfg)
	if len(os.Args) < 2 {
		fmt.Println("Expected CLI subcommands")
		usage()
	}

	switch os.Args[1] {

	case "db:create":
		create()
	case "db:migrate":
		fmt.Println("Database migrate")
		migrate()

	case "db:seed":
		fmt.Println("Database seed")
		seed()

	default:
		fmt.Println("Invalid subcommands")
		usage()
	}
}
