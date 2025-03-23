package usecase

import (
	"context"
	"graphql_clean_practice/internal/entity"
	"graphql_clean_practice/internal/infrastructure/repository"
)

type BookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) *BookUsecase {
	return &BookUsecase{bookRepo: bookRepo}
}

func (uc *BookUsecase) GetBooks(ctx context.Context) ([]*entity.Book, error) {
	return uc.bookRepo.FindAll(ctx)
}

func (uc *BookUsecase) GetBooksByAuthorID(ctx context.Context, authorID int) ([]*entity.Book, error) {
	return uc.bookRepo.FindByAuthorID(ctx, authorID)
}
