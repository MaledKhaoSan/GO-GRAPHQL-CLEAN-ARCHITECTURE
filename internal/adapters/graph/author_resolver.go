package graph

import (
	"context"
	"fmt"
	"graphql_clean_practice/internal/adapters/graph/model"
	"strconv"
)

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	authors, err := r.AuthorUC.GetAuthors(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.Author
	for _, a := range authors {
		result = append(result, &model.Author{
			ID:      strconv.Itoa(a.ID),
			Name:    a.Name,
			Country: &a.Country,
		})
	}
	return result, nil
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}
