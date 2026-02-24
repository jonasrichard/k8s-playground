package repo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Repo struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Client *redis.Client
}

func New() *Repo {
	repository := &Repo{
		Host: "filecache-service",
		Port: 6379,
		User: "",
		Pass: "",
	}

	repository.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", repository.Host, repository.Port),
		Password: repository.Pass,
		DB:       0,
	})

	if err := repository.Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return repository
}

func (r *Repo) Set(ctx context.Context, filename, size string) error {
	return r.Client.Set(ctx, filename, size, 0).Err()
}

func (r *Repo) Get(ctx context.Context, filename string) (string, error) {
	return r.Client.Get(ctx, filename).Result()
}

func (r *Repo) Keys(ctx context.Context) ([]string, error) {
	var cursor uint64
	keys := make([]string, 0)
	for {
		ks, cur, err := r.Client.Scan(ctx, cursor, "*", 100).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, ks...)
		cursor = cur
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}
