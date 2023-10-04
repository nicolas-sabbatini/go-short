package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	urlShort "github.com/nicolas-sabbatini/go-short/pkg/url-short"
)

func main() {
	mux := defaultMux()

	yamlPath := flag.String("yamlPath", "inputs/urls.yaml", "Path to the YAML url file")
	jsonPath := flag.String("jsonPath", "inputs/urls.json", "Path to the JSON url file")
	flag.Parse()

	yamlFile, err := os.ReadFile(*yamlPath)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlShort.YAMLHandler(yamlFile, mux)
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.ReadFile(*jsonPath)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlShort.JSONHandler(jsonFile, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
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
