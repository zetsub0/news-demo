package main

import (
	"context"
	"log"

	"github.com/denchick/news-aggregator/repo"
)

func main() {
	ctx := context.Background()
	repos, err := repo.NewArticlesRepository(ctx)
	if err != nil {
		log.Fatal(err)
	}

	art := &repo.Article{
		ID:          92929,
		URL:         "https://www.example.com",
		Title:       "do electric sheep dream of androids quotes",
		Description: "No, they do not",
	}

	err = repos.Create(ctx, art)
	if err != nil {
		log.Fatal(err)
	}
}
