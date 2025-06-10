package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type wordRequest struct {
	Word string `json:"word"`
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Received word " + req.Word})
}

func main() {
	http.HandleFunc("/api/status", statusHandler)
	http.HandleFunc("/api/game/week", weekHandler)
	http.HandleFunc("/api/game/word", wordHandler)

	log.Println("Server listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
