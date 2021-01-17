package models

import (
	"time"
)

// Book Struct (Model)
type Book struct {
	ID        string    `json:"_id,omitempty"bson:"_id"`
	CreatedAt time.Time `json:"created_at,"bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,"bson:"updated_at"`
	Isbn      string    `json:"isbn,omitempty"bson:"isbn"`
	Title     string    `json:"title,omitempty"bson:"title"`
	Author    *Author   `json:"author,omitempty"bson:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname,omitempty"bson:"firstname"`
	Lastname  string `json:"lastname,omitempty"bson:"lastname"`
}
