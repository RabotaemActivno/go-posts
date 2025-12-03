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

	return &Storage{db: db}, nil
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

	stmt, err := s.db.Prepare("SELECT id, author, text FROM posts")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

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

func (s *Storage) RemovePost(id int64) (int64, error) {
	const op = "storage.sqlite.RemovePost"
	
	stmt, err := s.db.Prepare("DELETE FROM posts WHERE id = (?)")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement %w", op, err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, fmt.Errorf("%s: exec: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("%s: rows affected: %w", op, err)
	}
	if affected == 0 {
		return 0, storage.ErrPostNotFound
	}
	return id, nil
}

func (s *Storage) UpdatePost(id int64, author string, text string) (storage.Post, error) {
	const op = "storage.sqlite.UpdatePost"

	stmt, err := s.db.Prepare("UPDATE posts SET author = (?), text = (?) WHERE id = (?)")
	if err != nil {
		return storage.Post{}, fmt.Errorf("%s: prepare statement %w", op, err)
	}

	res, err := stmt.Exec(author, text, id)
	if err != nil {
		return storage.Post{}, fmt.Errorf("%s: exec: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return storage.Post{}, fmt.Errorf("%s: rows affected: %w", op, err)
	}
	if affected == 0 {
		return storage.Post{}, storage.ErrPostNotFound
	}

	return storage.Post{
		ID: id,
		Author: author,
		Text: text,
	}, nil
}
