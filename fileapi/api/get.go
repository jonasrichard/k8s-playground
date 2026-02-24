package api

import (
	"context"
	"encoding/json"
	"fileapi/repo"
	"log"
	"net/http"
)

func KeysHandler(rp *repo.Repo, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("getting keys: %s", r.URL.Path)

	keys, err := rp.Keys(context.Background())
	if err != nil {
		http.Error(w, "redis keys error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(keys)
}
