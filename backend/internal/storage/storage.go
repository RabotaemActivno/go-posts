package storage

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
)

type Post struct {
	ID int64
	Author string
	Text string
}