package handlers

import (
	//"github.com/gofiber/fiber/v2"
	//"net/http"
	// "github.com/razonyang/fastrouter"
	"github.com/labstack/echo/v4"
	//"log"
	//"strings"
)

//var Funcs map[string]interface{}

func InitServer(settings map[string]interface{}, funcs map[string]interface{}){ 
	e := echo.New()
    Init_Test(e)
	//e.GET("/hello", Greetings)
    e.Logger.Fatal(e.Start(":3000"))
	 //app) // handlers
	//Init_SettingsRoutes(r) //app)

	// r.TrailingSlashesPolicy = fastrouter.StrictTrailingSlashes
	//PanicHandler
	// r.PanicHandler = func(w http.ResponseWriter, req *http.Request, rcv interface{}) {
	// 	log.Printf("received a panic: %#v\n", rcv)
	// }
	// // OptionsHandler
	// r.OptionsHandler = func(w http.ResponseWriter, req *http.Request, methods []string) {
	// 	w.Header().Set("Allow", strings.Join(methods, ", "))
	// 	w.Write([]byte("user-defined OptionsHandler"))
	// }
	// r.MethodNotAllowedHandler = func(w http.ResponseWriter, req *http.Request, methods []string) {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	w.Header().Set("Allow", strings.Join(methods, ", "))
	// 	w.Write([]byte("user-defined MethodNotAllowedHandler"))
	// }
	// // NotFoundHandler
	// r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("user-defined NotFoundHandler"))
	// })
	//log.Fatal(app.Listen(":3000"))  // запуск 
	// r.Prepare()
	// log.Fatal(http.ListenAndServe(":3000", r))
}