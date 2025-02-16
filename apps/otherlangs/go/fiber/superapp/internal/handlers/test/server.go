package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	//"runtime"
)

var Funcs map[string]interface{}

func InitServer(settings map[string]interface{}, funcs map[string]interface{}){ 
	//PreSettings(settings) ;	Funcs = funcs // исполняемые функции
	//runtime.GOMAXPROCS(3)
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		//CaseSensitive: true,
		//StrictRouting: true,
		//ServerHeader:  "Fiber",
		//AppName: "Test App v1.0.1"
	}) // инициализация
	
	// инициализация роутов
	Init_Test(app) // handlers
	//Init_SettingsRoutes(app)

	log.Fatal(app.Listen(":3000"))  // запуск 
}