package main

import (
	"fmt"
	"log"
	"net/http"

	"docshub/service"
	"github.com/julienschmidt/httprouter"
)

//Health check status 200 and content is ["OK"]
// func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprint(w, "[\"OK\"]")
// }
//Health check status 200 and content is ["OK"]
func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "[\"OK\"]")
}

func Render(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, *service.RenderMd())
}

func main() {
	fmt.Println("hello")
	router := httprouter.New()

	// mux := http.NewServeMux()
	// router := httprouter.New()
	router.GET("/api/health", Health)
	router.GET("/api/html", Render)

	// mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./app"))))
	// mux.HandleFunc("/api/health", Health)
	log.Fatal(http.ListenAndServe(":9000", router))
}
