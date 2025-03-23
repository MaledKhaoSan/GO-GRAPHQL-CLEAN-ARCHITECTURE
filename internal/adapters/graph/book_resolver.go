package graph

import (
	"context"
	"fmt"
	"graphql_clean_practice/internal/adapters/graph/model"
	"strconv"
)

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	// เรียก Usecase
	books, err := r.BookUC.GetBooks(ctx)
	if err != nil {
		return nil, err
	}

	// Map ไปยัง GraphQL Model
	var result []*model.Book
	for _, b := range books {
		result = append(result, &model.Book{
			ID:            strconv.Itoa(b.ID),
			Title:         b.Title,
			PublishedYear: &b.PublishedYear,
		})
	}

	return result, nil
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: Book - book"))
}

// BooksByAuthor is the resolver for the booksByAuthor field.
func (r *queryResolver) BooksByAuthor(ctx context.Context, authorID string) ([]*model.Book, error) {
	id, err := strconv.Atoi(authorID)
	if err != nil {
		return nil, err
	}

	books, err := r.BookUC.GetBooksByAuthorID(ctx, id)
	if err != nil {
		return nil, err
	}

	var result []*model.Book
	for _, b := range books {
		var authors []*model.Author
		for _, a := range b.Authors {
			authors = append(authors, &model.Author{
				ID:      strconv.Itoa(a.ID),
				Name:    a.Name,
				Country: &a.Country,
			})
		}

		result = append(result, &model.Book{
			ID:            strconv.Itoa(b.ID),
			Title:         b.Title,
			PublishedYear: &b.PublishedYear,
			Authors:       authors,
		})
	}
	return result, nil
}

// BooksBorrow is the resolver for the booksBorrow field.
func (r *queryResolver) BooksBorrow(ctx context.Context) ([]*model.BooksBorrow, error) {
	panic(fmt.Errorf("not implemented: BooksBorrow - booksBorrow"))
}

// BooksBorrowHistoryByUser is the resolver for the booksBorrowHistoryByUser field.
func (r *queryResolver) BooksBorrowHistoryByUser(ctx context.Context, userID string) ([]*model.BooksBorrow, error) {
	panic(fmt.Errorf("not implemented: BooksBorrowHistoryByUser - booksBorrowHistoryByUser"))
}

// BooksBorrowStatByAuthor is the resolver for the booksBorrowStatByAuthor field.
func (r *queryResolver) BooksBorrowStatByAuthor(ctx context.Context, authorID string) ([]*model.BorrowStat, error) {
	panic(fmt.Errorf("not implemented: BooksBorrowStatByAuthor - booksBorrowStatByAuthor"))
}
