package app

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/saufiroja/graphql-go-example/config"
	"github.com/saufiroja/graphql-go-example/internal/delivery/graph"
	"github.com/saufiroja/graphql-go-example/internal/repositories"
	"github.com/saufiroja/graphql-go-example/internal/services"
	"github.com/saufiroja/graphql-go-example/pkg/database"
	"github.com/saufiroja/graphql-go-example/pkg/logger"
	"github.com/sirupsen/logrus"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	// init config
	conf := config.NewAppConfig()

	// init logger
	logger.NewLogger()

	// init database
	db := database.NewPostgres(conf)

	// application
	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		TodoService: todoService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logrus.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.GraphQL.Port)
	logrus.Fatal(http.ListenAndServe(":"+conf.GraphQL.Port, nil))
}
