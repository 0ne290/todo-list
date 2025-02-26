package core

import (
	"context"
	"github.com/0ne290/todo-list/internal/core/task"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IUnitOfWork interface {
	TaskRepository() task.ITaskRepository
	Save(ctx context.Context)
	Rollback(ctx context.Context)
}

type UnitOfWork struct {
	repository *task.TaskRepository
}

func NewUnitOfWork(ctx context.Context, pool *pgxpool.Pool) *UnitOfWork {
	transaction, err := pool.Begin(ctx)

	if err != nil {
		panic(err.Error())
	}

	return &UnitOfWork{task.NewTaskRepository(transaction)}
}

func (unitOfWork *UnitOfWork) TaskRepository() task.ITaskRepository {
	return unitOfWork.repository
}

func (unitOfWork *UnitOfWork) Save(ctx context.Context) {
	if err := unitOfWork.repository.Transaction.Commit(ctx); err != nil {
		panic(err.Error())
	}
}

func (unitOfWork *UnitOfWork) Rollback(ctx context.Context) {
	if err := unitOfWork.repository.Transaction.Rollback(ctx); err != nil {
		panic(err.Error())
	}
}