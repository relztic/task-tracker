package domain

import "github.com/relztic/task-tracker/shared/types"

type CLIRouter interface {
	Route(cmd types.Command) (string, error)
}

type CLIHandler interface {
	Create(description string) (string, error)
	Read(id uint) (string, error)
	Update(id uint, description string) (string, error)
	Delete(id uint) (string, error)
	MarkAsTodo(id uint) (string, error)
	MarkAsInProgress(id uint) (string, error)
	MarkAsDone(id uint) (string, error)
	List() (string, error)
	ListTodo() (string, error)
	ListInProgress() (string, error)
	ListDone() (string, error)
	ListDeleted() (string, error)
	ListAll() (string, error)
	Clean() (string, error)
}

type CLIInteractor interface {
	Create(description string) (types.Task, error)
	Read(id uint) (types.Task, error)
	Update(id uint, description string, status types.TaskStatus) (types.Task, error)
	Delete(id uint) (types.Task, error)
	List() ([]types.Task, error)
	ListTodo() ([]types.Task, error)
	ListInProgress() ([]types.Task, error)
	ListDone() ([]types.Task, error)
	ListDeleted() ([]types.Task, error)
	ListAll() ([]types.Task, error)
	Clean() error
}

type CLIRepository interface {
	Create(task types.Task) (types.Task, error)
	Read(id uint) (types.Task, error)
	Update(id uint, task types.Task) (types.Task, error)
	Delete(id uint) (types.Task, error)
	List(predicate func(value types.Task, index uint, slice []types.Task) bool) ([]types.Task, error)
	Clean() error
}

type CLIDatabase interface {
	Read() ([]types.Task, error)
	Write(tasks []types.Task) error
	Drop() error
}
