package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Declare all variables here
var (
	port string
	cwd  string
	err  error
)

// Log the request & return the handler
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	// Different port can be assigned with command line variable -port
	flag.StringVar(&port, "port", "8080", "Assign a port for the server to listen to. Defaults to 8080.")
	flag.StringVar(&cwd, "root", "", "Set a root directory for server. Defaults to current working directory.")
	flag.Parse()

	// Print the assigned port
	log.Println("Listening on port", port)

	// Get current working directory if not set with command line argument
	if len(cwd) == 0 {
		if cwd, err = os.Getwd(); err != nil {
			log.Fatal(err) // Panic if something goes wrong
		}
	}

	// Log a fatal error if a HTTP server fails
	// Server listens to %port% and servers the contents of the current working directory
	// If index.html (.htm not OK) is present, serve that on root request instead of directory listing
	log.Fatal(http.ListenAndServe(":" + port, logger(http.FileServer(http.Dir(cwd)))))
}
