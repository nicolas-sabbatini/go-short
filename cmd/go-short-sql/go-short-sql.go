package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	sqliteUrlShort "github.com/nicolas-sabbatini/go-short/pkg/sqlite-url-short"
)

func main() {
	mux := defaultMux()

	sqlitlePath := flag.String("sqlitlePath", "inputs/urls.sqlite", "Path to the DB file")
	flag.Parse()

	db := sqliteUrlShort.OpenDb(*sqlitlePath)
	defer sqliteUrlShort.CloseDb(db)

	sqliteUrlShort.CreateUrlTable(db)
	sqliteUrlShort.UpsertUrl(db, "/g", "https://golang.org")
	sqliteUrlShort.UpsertUrl(db, "/y", "https://www.youtube.com")
	sqliteUrlShort.UpsertUrl(db, "/n", "https://nikcodes.xyz")
	sqliteUrlShort.UpsertUrl(db, "/g", "https://google.com")

	sqlHandler := sqliteUrlShort.SqlHandler(db, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", sqlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello, world! from /")
	fmt.Fprintln(w, "Hello, world!")
}
