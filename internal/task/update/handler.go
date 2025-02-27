package update

import (
	"context"
	"errors"

	"github.com/0ne290/todo-list/internal"
	"github.com/0ne290/todo-list/internal/task"
	"github.com/gofiber/fiber/v2"
)

// @Description Change the task status.
// @Summary Possible transitions between statuses are: new -> in_progress -> done. If this endpoint is called for a task that is in the final status (done), a 400 code will be returned to the client.
// @Tags Task
// @Accept json
// @Produce json
// @Param id query int true "Id"
// @Success 200 {object} task.Task
// @Router /tasks/{id:int} [put]
func Handle(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	ctx := c.Context()
	task, err := updateTask(id, ctx, internal.NewUnitOfWork(ctx), internal.NewTimeProvider())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "task is updated",
		"task":  task,
	})
}

func updateTask(id int, ctx context.Context, unitOfWork internal.IUnitOfWork, timeProvider internal.ITimeProvider) (*task.Task, error) {
	defer unitOfWork.Save(ctx)
	repository := unitOfWork.TaskRepository()

	task := repository.GetById(ctx, id)
	if task == nil {
		return nil, errors.New("task does not exists")
	}

	err := task.Update(timeProvider.Now())
	if err != nil {
		return nil, err
	}
	repository.Update(ctx, task)

	return task, nil
}
