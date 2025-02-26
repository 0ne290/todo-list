package create

import (
	"github.com/gofiber/fiber/v2"
	"github.com/0ne290/todo-list/internal/task"
	"github.com/0ne290/todo-list/internal"
	"context"
)

// @Description Create a new task.
// @Summary create a new task
// @Tags Task
// @Accept json
// @Produce json
// @Param requestDto body RequestDto true "RequestDto"
// @Success 200 {object} task.Task
// @Router /tasks [post]
func Handle(c *fiber.Ctx) error {
	request := &RequestDto{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ctx := c.Context()
	task := createTask(request, ctx, internal.NewUnitOfWork(ctx), internal.NewTimeProvider())

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"task":  task,
	})
}

func createTask(request *RequestDto, ctx context.Context, unitOfWork internal.IUnitOfWork, timeProvider internal.ITimeProvider) *task.Task {
	defer unitOfWork.Save(ctx)

	task := task.NewTask(request.Title, request.Description, timeProvider.Now())

	unitOfWork.TaskRepository().Add(ctx, task)

	return task
}