package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type wordRequest struct {
	Word string `json:"word"`
}

var db *sql.DB

func initDB(source string) (*sql.DB, error) {
	d, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}

	if _, err := d.Exec(`PRAGMA journal_mode=WAL; PRAGMA foreign_keys=ON; PRAGMA busy_timeout=5000;`); err != nil {
		d.Close()
		return nil, err
	}

	if _, err := d.Exec(`CREATE TABLE IF NOT EXISTS words (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                word TEXT NOT NULL,
                created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`); err != nil {
		d.Close()
		return nil, err
	}
	return d, nil
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func weekHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string][]string{"letters": []string{"a", "b", "c", "d", "e", "f", "g"}})
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	var req wordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Word == "" {
		http.Error(w, "word cannot be empty", http.StatusBadRequest)
		return
	}

	if _, err := db.Exec("INSERT INTO words(word) VALUES(?)", req.Word); err != nil {
		http.Error(w, "could not store word", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Received word " + req.Word})
}

func main() {
	var err error
	db, err = initDB("wordgame.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/api/status", statusHandler)
	http.HandleFunc("/api/game/week", weekHandler)
	http.HandleFunc("/api/game/word", wordHandler)

	log.Println("Server listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
