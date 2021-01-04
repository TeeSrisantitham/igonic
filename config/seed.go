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

// Seeds import seeds files into database by env yml file
func Seeds(db *gorm.DB, env string) error {
	raw, err := ioutil.ReadFile("db/seeds/" + env)
	if err != nil {
		return err
	}

	seedInfo := seed{}
	err = yaml.Unmarshal(raw, &seedInfo)
	if err != nil {
		return err
	}

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
