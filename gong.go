package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	port   string
	dir    string
	err    error
	lop    string
	output *log.Logger = log.New(os.Stdout, "", log.Ldate | log.Ltime)
)

func init() {
	parseFlags()

	// Get current working directory if not set with command line argument.
	if dir == "" {
		if dir, err = os.Getwd(); err != nil {
			log.Fatal(err) // Exit if something goes wrong.
		}
	}
}

func main() {
	// If a log file was set, open it and set it as the output.
	if lop != "" {
		// Open the file write only. Use append and create it if nonexistant.
		f, err := os.OpenFile(lop, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err) // Something went wrong while opening the log file.
		}
		defer f.Close() // Close the log file when main() exits.

		output = log.New(f, "", log.Ldate | log.Ltime) // Set logger.
		output.Println("Server started on port", port) // Log the start.
		output.Println("Document root:", dir) // Log the document root.
	}
	
	// Print the assigned port on screen.
	log.Println("Listening on port", port)

	// Print the document root path on screen.
	log.Println("Document root:", dir)
	
	// Log a fatal error if the HTTP server fails.
	// Server listens on %port% and serves the contents of %dir%.
	// If index.html (.htm not OK) is present, serve that on root request instead of directory listing.
	output.Fatal(http.ListenAndServe(":" + port, logger(http.FileServer(http.Dir(dir)))))
}

// Log the request & return the handler.
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output.Println(r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func parseFlags() {
	flag.StringVar(&port, "port", "8080", "Assign a port for the server to listen to. Defaults to 8080.")
	flag.StringVar(&dir, "root", "", "Set a root directory for server. Defaults to current working directory.")
	flag.StringVar(&lop, "log", "", "Optional log output file.")
	flag.Parse()
}
