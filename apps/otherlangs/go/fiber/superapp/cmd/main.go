package main


import (
	//"log"
	//"superapp/internal"
	"superapp/internal/handlers/test"
	//msql "superapp/internal/storage/mysql"
	//msql "superapp/internal/storage/mysql_gorm"
	//"superapp/config"
	//"superapp/internal/models"
)

// https://go.libhunt.com/fiber-alternatives

func main() {
	//app_settings := internal.InitConfig()
	//Config := app_settings["config"].(*config.Cfg)
	handlers.InitServer(map[string]interface{}{}, map[string]interface{}{})//app_settings, map[string]interface{}{})
}