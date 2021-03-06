package config

import (
	"fmt"
	"log"

	"github.com/openware/igonic/models"
	"github.com/openware/igonic/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDatabase : connect to database MySQL using gorm
// gorm (GO ORM for SQL): http://gorm.io/docs/connecting_to_the_database.html
func ConnectDatabase() (db *gorm.DB) {
	dbDriver := utils.GetEnv("DATABASE_DRIVER", "mysql")
	dbHost := utils.GetEnv("DATABASE_HOST", "localhost")
	dbPort := utils.GetEnv("DATABASE_PORT", "3306")
	dbName := utils.GetEnv("DATABASE_NAME", "opendax_development")
	dbUser := utils.GetEnv("DATABASE_USER", "root")
	dbPass := utils.GetEnv("DATABASE_PASS", "")

	var err error
	var dial gorm.Dialector

	switch dbDriver {
	case "memory":
		dial = sqlite.Open(":memory:")

	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName,
		)
		dial = mysql.Open(dsn)

	default:
		panic("Unsupported DB_DRIVER: " + dbDriver)

	}
	db, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Pass db connection to package controllers and models
	models.SetUpDBConnection(db)
	// controllers.SetUpDBConnection(db)

	return db
}

// RunMigrations create and modify database tables according to the models
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Page{})
}

// LoadSeeds import seed files into database
func LoadSeeds(db *gorm.DB) {
	err := models.SeedArticles(db)
	if err != nil {
		log.Fatal(err)
	}
	err = models.SeedPages(db)
	if err != nil {
		log.Fatal(err)
	}
}
