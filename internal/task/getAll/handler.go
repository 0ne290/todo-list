package getAll

import (
	"github.com/gofiber/fiber/v2"
	"github.com/0ne290/todo-list/internal/task"
	"github.com/0ne290/todo-list/internal"
	"context"
)

// @Description Get all tasks.
// @Summary get all tasks
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} []task.Task
// @Router /tasks [get]
func Handle(c *fiber.Ctx) error {
	ctx := c.Context()
	tasks := getAllTasks(ctx, internal.NewUnitOfWork(ctx))

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tasks":  tasks,
	})
}

func getAllTasks(ctx context.Context, unitOfWork internal.IUnitOfWork) []*task.Task {
	defer unitOfWork.Rollback(ctx)

	return unitOfWork.TaskRepository().GetAll(ctx)
}