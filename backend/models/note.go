package models

import "time"

type Note struct {
	ID        int        `db:"id" json:"id"`
	UserID    int        `db:"user_id" json:"user_id"`
	Title     string     `db:"title" json:"title"`
	Content   string     `db:"content" json:"content"`
	ImageURL  string     `db:"image_url" json:"image_url"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
