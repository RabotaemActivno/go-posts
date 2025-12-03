package update

import (
	"go-posts/internal/storage"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Request struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type BadResponse struct {
	Status string `json:"status"`
	Text string `json:"text"`
}

type PostUpdater interface {
	UpdatePost(id int64, author string, text string) (storage.Post, error)
}

func New(log *slog.Logger, postUpdater PostUpdater) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.update.New"

		log := log.With("op", op)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request")

			render.JSON(w, r, BadResponse{
				Status: "Error",
				Text: "Failed to decode request",
			})

			return 
		}

		strID := chi.URLParam(r, "id")
		if strID == "" {
			log.Info("invalid id")
			render.JSON(w, r, BadResponse{
				Status: "Error",
				Text: "Invalid id",
			})
			return 
		}

		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			log.Info("failed to convert string to int64")
			render.JSON(w, r, BadResponse{
				Status: "Error",
				Text: "Bad response",
			})
			return 
		}

		log.Info("Request decoded")

		post, err := postUpdater.UpdatePost(id, req.Author, req.Text)
		if err != nil {
			log.Error("failed to update post")
			render.JSON(w, r, BadResponse{
				Status: "Error",
				Text: "Failed to update post",
			})
		}

		log.Info("Post updated")

		render.JSON(w, r, post)
	}
}