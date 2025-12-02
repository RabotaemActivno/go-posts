package storage

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
)

type Post struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}
