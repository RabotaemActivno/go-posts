package save

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type Request struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type FailedResponse struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

type PostSaver interface {
	SavePost(author string, text string) (int64, error)
}

func New(log *slog.Logger, postSaver PostSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log := log.With("op", op)

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request")

			failedResp := struct {
				status string
				text   string
			}{
				status: "Error",
				text:   "Bad request",
			}

			render.JSON(w, r, failedResp)

			return
		}

		log.Info("request body decoded")

		id, err := postSaver.SavePost(req.Author, req.Text)
		if err != nil {
			log.Error("failed to save post")

			render.JSON(w, r, FailedResponse{
				Status: "Error",
				Text:   "Failed to save url",
			})

			return
		}

		log.Info("post saved")

		successResp := struct {
			Status string `json:"status"`
			PostID int64  `json:"postID"`
		}{
			Status: "OK",
			PostID: id,
		}

		render.JSON(w, r, successResp)
	}
}
