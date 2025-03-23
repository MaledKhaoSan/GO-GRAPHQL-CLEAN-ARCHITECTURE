package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"graphql_clean_practice/internal/adapters/graph"
	"graphql_clean_practice/internal/adapters/graph/generated"
	"graphql_clean_practice/internal/infrastructure/postgresql"
	"graphql_clean_practice/internal/infrastructure/repository"
	"graphql_clean_practice/internal/usecase"
)

func main() {
	_ = godotenv.Load()

	// 🔌 Connect DB via GORM
	db, err := postgresql.NewPostgresGorm()
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	// DI: Repository -> Usecase -> Resolver
	bookRepo := repository.NewBookRepository(db)
	bookUC := usecase.NewBookUsecase(bookRepo)

	authorRepo := repository.NewAuthorRepository(db) // <— เพิ่ม
	authorUC := usecase.NewAuthorUsecase(authorRepo) // <— เพิ่ม

	resolver := &graph.Resolver{
		BookUC:   bookUC,
		AuthorUC: authorUC,
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)

	// 🚀 Fiber App
	app := fiber.New()
	app.Use(logger.New())

	// 📍 gqlgen ผ่าน Fiber
	app.All("/query", adaptor.HTTPHandler(srv))

	// 🧪 Playground
	app.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL", "/query")))

	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
