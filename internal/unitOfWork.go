package internal

import (
	"context"
	"github.com/0ne290/todo-list/internal/task"
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

var DatabaseConnectionPool *pgxpool.Pool

func NewUnitOfWork(ctx context.Context) *UnitOfWork {
	transaction, err := DatabaseConnectionPool.Begin(ctx)

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