package repository

import (
	"context"
	"graphql_clean_practice/internal/entity"

	"gorm.io/gorm"
)

// Interface สำหรับ usecase
type AuthorRepository interface {
	FindAll(ctx context.Context) ([]*entity.Author, error)
}

// Implementation (ไม่ expose)
type authorRepositoryImpl struct {
	db *gorm.DB
}

// Constructor
func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepositoryImpl{db: db}
}

func (r *authorRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Author, error) {
	var authors []*entity.Author
	if err := r.db.WithContext(ctx).Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
