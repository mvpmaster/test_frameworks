// package main

// import (
// 	"log"
// 	"net/http"
// 	//"strings"
// 	"github.com/razonyang/fastrouter"
// )


package main

import (
	//"log"
	//"superapp/internal"
	"superapp/internal/handlers/test"
	//msql "superapp/internal/storage/mysql"
	//msql "superapp/internal/storage/mysql_gorm"
	//"superapp/config"
	//"superapp/internal/models"
	// "runtime"
	// "fmt"
)

func main() {
	//app_settings := internal.InitConfig()
	//Config := app_settings["config"].(*config.Cfg)
	// cores := runtime.NumCPU()
	// fmt.Printf("This machine has %d CPU cores. \n", cores)
	// runtime.GOMAXPROCS(cores)
	handlers.InitServer( map[string]interface{}{},  map[string]interface{}{})//app_settings, map[string]interface{}{})
}

// func main(){

// 		r := fastrouter.New()
	
// 		// TrailingSlashesPolicy
// 		//r.TrailingSlashesPolicy = fastrouter.StrictTrailingSlashes
// 		// PanicHandler
// 		// r.PanicHandler = func(w http.ResponseWriter, req *http.Request, rcv interface{}) {
// 		// 	log.Printf("received a panic: %#v\n", rcv)
// 		// }
// 		// // OptionsHandler
// 		// r.OptionsHandler = func(w http.ResponseWriter, req *http.Request, methods []string) {
// 		// 	w.Header().Set("Allow", strings.Join(methods, ", "))
// 		// 	w.Write([]byte("user-defined OptionsHandler"))
// 		// }
// 		// // MethodNotAllowedHandler
// 		// r.MethodNotAllowedHandler = func(w http.ResponseWriter, req *http.Request, methods []string) {
// 		// 	w.WriteHeader(http.StatusMethodNotAllowed)
// 		// 	w.Header().Set("Allow", strings.Join(methods, ", "))
// 		// 	w.Write([]byte("user-defined MethodNotAllowedHandler"))
// 		// }
// 		// // NotFoundHandler
// 		// r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.WriteHeader(http.StatusNotFound)
// 		// 	w.Write([]byte("user-defined NotFoundHandler"))
// 		// })
	
// 		// // homepage handler
// 		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.Write([]byte("hello world"))
// 		// })
// 		// // panic handler
// 		// r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
// 		// 	panic("panic handler")
// 		// })
// 		// // hello handler
// 		// r.Get("/hello/<name>", func(w http.ResponseWriter, r *http.Request) {
// 		// 	var name string
	
// 		// 	name = fastrouter.Params(r)["name"]
// 		// 	/* The following code snippet is equivalent.
// 		// 	if params, ok := r.Context().Value(fastrouter.ParamsKey{}).(map[string]string); ok {
// 		// 		name = params["name"]
// 		// 	}
// 		// 	*/
	
// 		// 	w.Write([]byte("hello " + name))
// 		// })
	
// 		// RESTful API
// 		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.Write([]byte("query users"))
// 		// })
// 		// r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.Write([]byte("create an user"))
// 		// })
// 		// r.Get("/users/<name>", func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.Write([]byte(fastrouter.Params(r)["name"] + " profile"))
// 		// })
// 		// r.Delete("/users/<name>", func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.Write([]byte("deleted user " + fastrouter.Params(r)["name"]))
// 		// })
// 		// r.Put("/users/<name>", func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.Write([]byte("updated " + fastrouter.Params(r)["name"] + " profile"))
// 		// })
// 		// r.Get("/users/<name>/posts", func(w http.ResponseWriter, r *http.Request) {
// 		// 	w.Write([]byte("the posts created by " + fastrouter.Params(r)["name"]))
// 		// })
	
// 		// Make preparations before handling incoming request.
// 		// Note that, this method MUST be invoked before handling incoming request,
// 		// otherwise the router can not works as expected.
// 		r.Prepare()
	
// 		log.Fatal(http.ListenAndServe(":3000", r))
	
	
// }