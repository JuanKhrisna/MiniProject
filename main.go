package main

import (
	"miniproject/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)
func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&{})
	DB.AutoMigrate(&{})
	DB.AutoMigrate(&{})
	DB.AutoMigrate(&{})
	DB.AutoMigrate(&{})
}

func main() {
	ConfigDB := config.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}
	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	e := echo.New()
}
