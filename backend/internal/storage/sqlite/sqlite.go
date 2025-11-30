package sqlite

import (
	"database/sql"
	"fmt"
	"go-posts/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS posts(
			id INTEGER PRIMARY KEY,
			author TEXT NOT NULL,
			text TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db:db}, nil
}

func (s *Storage) SavePost(author string, text string) (int64, error) {
	const op = "storage.sqlite.SavePost"

	stmt, err := s.db.Prepare("INSERT INTO posts(author, text) VALUES(?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(author, text)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetAllPosts() ([]storage.Post, error) {
	const op = "storage.sqlite.GetAllPosts"

	var posts []storage.Post

	stmt, err := s.db.Prepare("SELECT * FROM posts")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(100)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var p storage.Post
		err := rows.Scan(&p.ID, &p.Author, &p.Text)
		if err != nil {	
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		posts = append(posts, p)
	}

	return posts, nil
} 