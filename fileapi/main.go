package main

import (
	"context"
	"fileapi/api"
	"fileapi/repo"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func main() {
	// TODO use ENV vars for config
	repository := repo.New()

	log.Println("connected to redis at filecache:6379")

	http.HandleFunc("/api/kv", func(w http.ResponseWriter, r *http.Request) {
		api.KVHandler(repository, w, r)
	})
	http.HandleFunc("/api/keys", func(w http.ResponseWriter, r *http.Request) {
		api.KeysHandler(repository, w, r)
	})
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("fileapi is running\n"))
	})

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
