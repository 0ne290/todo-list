package task

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type ITaskRepository interface {
	Add(ctx context.Context, task *Task)
	GetAll(ctx context.Context) []*Task
	Update(ctx context.Context, task *Task)
	Remove(ctx context.Context, id int)
}

type TaskRepository struct {
	Transaction pgx.Tx
}

func NewTaskRepository(transaction pgx.Tx) *TaskRepository {
	return &TaskRepository{transaction}
}

func (repository *TaskRepository) Add(ctx context.Context, task *Task) {
	if _, err := repository.Transaction.Exec(ctx, "INSERT INTO tasks VALUES ($1, $2, $3, $4, $5)", task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt); err != nil {
		panic(err.Error())
	}
}

func (repository *TaskRepository) GetAll(ctx context.Context) []*Task {
	rows, err := repository.Transaction.Query(ctx, "SELECT * FROM tasks FOR UPDATE")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		task := &Task{}

		err := rows.Scan(&task.Id, &task.Title, task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return tasks
}

func (repository *TaskRepository) Update(ctx context.Context, task *Task) {
	if _, err := repository.Transaction.Exec(ctx, "UPDATE tasks SET status = $1, updated_at = $2 WHERE id = $3", task.Status, task.UpdatedAt, task.Id); err != nil {
		panic(err.Error())
	}
}

func (repository *TaskRepository) Remove(ctx context.Context, id int) {
	if _, err := repository.Transaction.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id); err != nil {
		panic(err.Error())
	}
}