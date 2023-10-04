package sqliteUrlShort

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb(filePath string) *sql.DB {
	log.Println("Opening database in", filePath, "...")
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}
	log.Println("Database opened successfully")
	return db
}

func CloseDb(db *sql.DB) {
	log.Println("Closing database ...")
	err := db.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Database closed successfully")
}

func CreateUrlTable(db *sql.DB) {
	log.Println("Creating table URL ...")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS URL (
		path TEXT PRIMARY KEY NOT NULL,
		url TEXT NOT NULL
	)`)
	if err != nil {
		panic(err)
	}
	log.Println("Table URL created successfully")
}

func UpsertUrl(db *sql.DB, path string, url string) {
	log.Println("Upserting URL ...")
	_, err := db.Exec(`INSERT INTO URL (path, url) VALUES (?, ?)
	ON CONFLICT(path) DO UPDATE SET url = excluded.url`, path, url)
	if err != nil {
		panic(err)
	}
	log.Println("URL upserted successfully")
}

func SelectUrl(db *sql.DB, path string) string {
	log.Println("Selecting URL ...")
	var url string
	err := db.QueryRow("SELECT url FROM URL WHERE path = ?", path).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("URL not found", err)
			return ""
		}
		panic(err)
	}
	log.Println("URL selected successfully", url)
	return url
}

func SqlHandler(db *sql.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("SqlHandler: ", r.URL.Path)
		url := SelectUrl(db, r.URL.Path)
		if url != "" {
			log.Println("SqlHandler: redirecting to ", url)
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
