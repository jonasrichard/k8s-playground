package api

import (
	"encoding/json"
	"fileapi/repo"
	"log"
	"net/http"
	"strings"
)

type kvRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func KVHandler(rp *repo.Repo, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodPut:
		var req kvRequest

		// try JSON first, fall back to form/query values
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			// not JSON — try form values / query params
			if err := r.ParseForm(); err == nil {
				req.Key = r.FormValue("key")
				req.Value = r.FormValue("value")
			}
		}

		// also allow key/value in URL query (convenience)
		if req.Key == "" {
			req.Key = r.URL.Query().Get("key")
		}
		if req.Value == "" {
			req.Value = r.URL.Query().Get("value")
		}

		if strings.TrimSpace(req.Key) == "" {
			http.Error(w, "missing key", http.StatusBadRequest)
			return
		}

		log.Printf("setting key %q", req.Key)

		if err := rp.Set(r.Context(), req.Key, req.Value); err != nil {
			http.Error(w, "redis set error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]string{"result": "ok", "key": req.Key})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
