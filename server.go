package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/Anuolu-2020/hackernews-api-clone/graph"
	"github.com/Anuolu-2020/hackernews-api-clone/internal/auth"
	"github.com/Anuolu-2020/hackernews-api-clone/internal/db"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error occurred while loading env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize database
	db.InitDB()

	// Close database on shutdown
	defer db.Close()

	router := chi.NewRouter()

	router.Use(auth.Middleware())
	router.Use(middleware.Logger)

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
