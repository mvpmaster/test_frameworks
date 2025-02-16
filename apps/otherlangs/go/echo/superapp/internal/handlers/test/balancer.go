package handlers

import (
	// "github.com/gofiber/fiber/v2"
	//"fmt"
	"net/http"
	//"github.com/razonyang/fastrouter"
	"github.com/labstack/echo/v4"
	//"superapp/internal/service/redirect_url"
	"strings"
	"encoding/json"
)

var test=strings.Repeat("test", 5000)
var test2 []byte //[]byte
var test3 string

//func Init_Test(app *fiber.App) { //go service.SomeWorker(); 
func Init_Test(e *echo.Echo) { //go service.SomeWorker(); 
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("query users"))
	// })
	// https://pkg.go.dev/net/http

	test2,_=json.Marshal(map[string]string{"test":test})
	test3=string(test2)

	//go redirect_url.WorkerBlanaceForgotten() // start worker thread

	// app.Get("/test",// http.HandlerFunc(
	// 			Go_RootRedirectUrl,
	// 		) 
	e.GET("/test", 	func(c echo.Context) error {
		return c.String(http.StatusOK, test3)
	  })
	// app.Head("/", http.HandlerFunc(
	// 	Go_RootRedirectUrl,
	// )) 
	} // some events

	// func Greetings(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, HelloWorld{
	// 		Message: "Hello World",
	// 	})
	// }

	// func(c echo.Context) error {
	// 	return c.String(http.StatusOK, test2)
	//   }

// func Go_RootRedirectUrl(w http.ResponseWriter, r *http.Request) {

// 	//url := r.URL.Query().Get("video") //c.Query("video") ; // log.Println(url)
// 	// url, err := redirect_url.RootRedirect(url) ; if err != nil { 
// 	// 	errorHandler(w, r, http.StatusNotFound) ; return 
// 	// 	//return c.Status(404).SendString(err.Error())
// 	//  } // log.Println(url)
// 	w.Write(test2) //[]byte("query users"))

// 	//http.Redirect(w, r, url, 301) 
// 	//http.Redirect(w, r, url, http.StatusSeeOther)
// 	//c.Redirect(url, 301) ; 
// 	} // перенаправление 

// 	//w.Write([]byte("query users"))

// 	// return c.Status(fiber.StatusOK).JSON( &url )
// // interface
// // func SetSettings(cdn_host string, 
// // 	             percent float64, 
// // 				 history_balance float64){
// // 	redirect_url.SetSettings(cdn_host, 
// // 							 percent, 
// // 							 history_balance) }


// 	// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	// 	if r.URL.Path != "/" {
// 	// 		errorHandler(w, r, http.StatusNotFound)
// 	// 		return
// 	// 	}
// 	// 	fmt.Fprint(w, "welcome home")
// 	// }
	
// 	// func smthHandler(w http.ResponseWriter, r *http.Request) {
// 	// 	if r.URL.Path != "/smth/" {
// 	// 		errorHandler(w, r, http.StatusNotFound)
// 	// 		return
// 	// 	}
// 	// 	fmt.Fprint(w, "welcome smth")
// 	// }
	
// 	func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
// 		w.WriteHeader(status)
// 		if status == http.StatusNotFound {
// 			fmt.Fprint(w, "custom 404")
// 		}
// 	}