package main

import "io"
import "github.com/gorilla/mux"
import "log"
import "net/http"
import "time"

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/_hc", HealthCheckHandler)

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}

// health check...
func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
