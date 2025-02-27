package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"os"

	"github.com/0ne290/todo-list/internal"
	createTask "github.com/0ne290/todo-list/internal/task/create"
	getAllTasks "github.com/0ne290/todo-list/internal/task/getAll"
	updateTask "github.com/0ne290/todo-list/internal/task/update"
	deleteTask "github.com/0ne290/todo-list/internal/task/delete"

	_ "github.com/0ne290/todo-list/docs"
)

// @title TODO list API
// @version 1.0
// @description Test task
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:80
// @BasePath /
func main() {
	app := fiber.New()

	registerMiddlewares(app)

	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		panic(err.Error())
	}
	internal.DatabaseConnectionPool = pool

	registerRoutes(app)

	app.Listen(":80")
}

func registerMiddlewares(app *fiber.App) {
	app.Use(
		logger.New(),
	)
}

func registerRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/tasks", createTask.Handle)
	app.Get("/tasks", getAllTasks.Handle)
	app.Put("/tasks/:id<int>", updateTask.Handle)
	app.Delete("/tasks/:id<int>", deleteTask.Handle)
}