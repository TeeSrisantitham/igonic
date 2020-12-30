package config

import (
	"fmt"
	"io/ioutil"

	"github.com/openware/igonic/models"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// Seed structure
type seed struct {
	Pages []models.Page `yaml:"pages"`
	Articles []models.Article `yaml:"articles"`
}

// Seeds import seeds files into database by env yml file
func Seeds(db *gorm.DB, env string) error {
	raw, err := ioutil.ReadFile("db/seeds/" + env)
	if err != nil {
		return err
	}

	Seed := seed{}
	err = yaml.Unmarshal(raw, &Seed)
	if err != nil {
		return err
	}
	fmt.Println("Seeding pages")
	tx := db.Create(&Seed.Pages)
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("Seeding articles")
	tx = db.Create(&Seed.Articles)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}