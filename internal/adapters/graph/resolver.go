// internal/adapters/graph/resolver.go
package graph

import (
	"graphql_clean_practice/internal/adapters/graph/generated"
	"graphql_clean_practice/internal/usecase"
)

type Resolver struct {
	BookUC   *usecase.BookUsecase
	AuthorUC *usecase.AuthorUsecase
}

// gqlgen expects this function to know what struct handles Query resolvers
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// struct ที่ใช้ implement Books() ให้ gqlgen รู้จัก
type queryResolver struct {
	*Resolver
}
