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

			failedResp := struct {
				status string
				text string
			} {
				status: "Error",
				text: "Failed to get posts",
			}

			log.Error("failed to get posts")
			
			render.JSON(w, r, failedResp)

			return
		}

		log.Info("posts get")

		successResp := struct {
			status string
			posts []storage.Post
		} {
			status: "OK",
			posts: posts,
		}

		render.JSON(w, r, successResp)
	}
}