// Entry point: sets up HTTP server and routes

package web

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ascii-art", withRecover(asciiArtHandler))
	mux.HandleFunc("/ascii-table", withRecover(asciiTableHandler))
	mux.HandleFunc("/error", withRecover(errorPageHandler))
	mux.HandleFunc("/", withRecover(indexHandler))
	mux.HandleFunc("/export", withRecover(handleExport))

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
