package config

import (
	"io/ioutil"
	"reflect"

	"github.com/openware/igonic/models"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// Seed structure
type seed struct {
	Pages []models.Page `yaml:"pages"`
	Articles []models.Article `yaml:"articles"`
}

func readENVFile(path string, seedInfo *seed) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(raw, &seedInfo)
}

func insertToDB(db *gorm.DB, seedInfo seed) error {
	v := reflect.ValueOf(seedInfo)

	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsNil()  {
			tx := db.Create(v.Field(i).Interface())
			if tx.Error != nil {
				return tx.Error
			}
		}
	}
	return nil
}

// Seeds import seeds files into database by env yml file
func Seeds(db *gorm.DB, env string) error {
	var seedInfo seed
	
	err := readENVFile("db/seeds/" + env, &seedInfo)
	if err != nil {
		return err
	}
	
	return insertToDB(db, seedInfo)
}
