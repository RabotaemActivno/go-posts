package getall

import (
	"go-posts/internal/storage"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type PostsGetter interface {
	GetAllPosts() ([]storage.Post, error)
}

func New(log *slog.Logger, postsGetter PostsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.getall.New"
		log := log.With("op", op)

		posts, err := postsGetter.GetAllPosts()
		if err != nil {

			log.Error("failed to get posts")

			render.JSON(w, r, struct {
				Status string `json:"status"`
				Text   string `json:"text"`
			}{
				Status: "Error",
				Text:   "Failed to get posts",
			})

			return
		}

		log.Info("posts get")

		render.JSON(w, r, struct {
			Status string         `json:"status"`
			Posts  []storage.Post `json:"posts"`
		}{
			Status: "OK",
			Posts:  posts,
		})
	}
}
