package remove

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PostRemover interface {
	RemovePost(id int64) (int64, error)
}

type BadResponse struct {
	Status string `json:"status"`
	Text string `json:"text"`
}

func New(log *slog.Logger, postRemover PostRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.remove.New"
		log := log.With("op", op)

		strID := chi.URLParam(r, "id")
		if strID == "" {
			log.Info("invalid id")
			render.JSON(w, r, BadResponse{
				Status: "Error",
				Text: "Invalid id",
			})
		}

		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			log.Info("failed to convert string to int64")
			render.JSON(w, r, BadResponse{
				Status: "Error",
				Text: "Bad response",
			})
		}

		id, err = postRemover.RemovePost(id) 
		if err != nil {
			log.Info("failed to remove post")
			render.JSON(w, r, BadResponse{
				Status: "Error",
				Text: "Bad response",
			})
		}

		log.Info("post removed")

		render.JSON(w, r, struct {
			Status string `json:"status"`
			ID int64 `json:"id"`
		}{
			Status: "OK",
			ID: id,
		})
	}
}