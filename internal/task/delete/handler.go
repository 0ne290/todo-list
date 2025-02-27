package delete

import (
	"context"
	"errors"

	"github.com/0ne290/todo-list/internal"
	"github.com/gofiber/fiber/v2"
)

// @Description Delete the task.
// @Summary delete the task
// @Tags Task
// @Accept json
// @Produce json
// @Param id query int true "Id"
// @Router /tasks/{id:int} [delete]
func Handle(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	ctx := c.Context()
	err := deleteTask(id, ctx, internal.NewUnitOfWork(ctx))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "task is deleted",
	})
}

func deleteTask(id int, ctx context.Context, unitOfWork internal.IUnitOfWork) error {
	defer unitOfWork.Save(ctx)
	repository := unitOfWork.TaskRepository()

	task := repository.GetById(ctx, id)
	if task == nil {
		return errors.New("task does not exists")
	}

	repository.Remove(ctx, id)

	return nil
}
