package repository

import (
	"context"
	"graphql_clean_practice/internal/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(ctx context.Context) ([]*entity.Book, error)
	FindByAuthorID(ctx context.Context, authorID int) ([]*entity.Book, error)
}

type bookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepositoryImpl{db: db}
}

func (r *bookRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Book, error) {
	var books []*entity.Book
	err := r.db.WithContext(ctx).
		Preload("Authors").
		Find(&books).Error
	return books, err
}

func (r *bookRepositoryImpl) FindByAuthorID(ctx context.Context, authorID int) ([]*entity.Book, error) {
	var books []*entity.Book
	err := r.db.WithContext(ctx).
		Joins("JOIN book_authors ON books.id = book_authors.book_id").
		Where("book_authors.author_id = ?", authorID).
		Preload("Authors").
		Find(&books).Error
	return books, err
}
