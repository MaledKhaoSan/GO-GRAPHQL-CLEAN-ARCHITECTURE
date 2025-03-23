// internal/entity/entity_model.go
package entity

import "time"

type Author struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Books   []*Book `gorm:"many2many:book_authors;" json:"books"`
}

type Book struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title"`
	PublishedYear int       `json:"published_year"`
	Authors       []*Author `gorm:"many2many:book_authors;" json:"authors"`
}

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type BooksBorrow struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
}
