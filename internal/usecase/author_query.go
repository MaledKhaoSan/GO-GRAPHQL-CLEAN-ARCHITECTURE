package usecase

import (
	"context"
	"graphql_clean_practice/internal/entity"
	"graphql_clean_practice/internal/infrastructure/repository"
)

type AuthorUsecase struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorUsecase(authorRepo repository.AuthorRepository) *AuthorUsecase {
	return &AuthorUsecase{authorRepo: authorRepo}
}

func (uc *AuthorUsecase) GetAuthors(ctx context.Context) ([]*entity.Author, error) {
	return uc.authorRepo.FindAll(ctx)
}
