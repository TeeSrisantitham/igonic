package helper

import (
	"github.com/openware/igonic/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetTestContext() *gorm.DB {
	driver := utils.GetEnv("DB_TEST_DRIVER", "sqlite3")

	switch driver {
	case "sqlite3":
		url := ":memory:?parseTime=True"
		ctx := Connect(driver, url)

		return ctx
	}
	panic("Driver not supported yet in test")
}

// Connect returns a SQL client
func Connect(driver, url string) *gorm.DB {
	var dial gorm.Dialector

	switch driver {
	case "sqlite3":
		dial = sqlite.Open(url)

	case "mysql":
		dial = mysql.Open(url)

	default:
		panic("Unsupported driver " + driver)

	}

	db, err := gorm.Open(dial, &gorm.Config{})

	if err != nil {
		panic("Connection to database failed: " + err.Error())
	}

	return db
}