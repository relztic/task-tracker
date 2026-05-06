package cli

import (
	"fmt"
	"time"

	sliceutils "github.com/relztic/go-utils/v2/slice"
	"github.com/relztic/task-tracker/shared/domain"
	"github.com/relztic/task-tracker/shared/types"
)

type Repository struct {
	cliDatabase domain.CLIDatabase
}

var _ domain.CLIRepository = &Repository{}

func NewRepository(cliDatabase domain.CLIDatabase) *Repository {
	return &Repository{
		cliDatabase: cliDatabase,
	}
}

func (r *Repository) Create(input types.Task) (types.Task, error) {
	tasks, err := r.cliDatabase.Read()
	if err != nil {
		return types.Task{}, err
	}

	id := uint(1)
	if len(tasks) > 0 {
		lastTask, _ := sliceutils.Pop(tasks)
		id = lastTask.ID + 1
	}

	task := types.Task{
		ID: id,

		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,

		Description: input.Description,
		Status:      input.Status,
	}

	err = r.cliDatabase.Write(sliceutils.Push(tasks, task))
	if err != nil {
		return types.Task{}, err
	}

	return task, nil
}

func (r *Repository) Read(id uint) (types.Task, error) {
	tasks, err := r.cliDatabase.Read()
	if err != nil {
		return types.Task{}, err
	}

	task, found := sliceutils.Find(tasks, func(task types.Task, _ uint, _ []types.Task) bool {
		return task.ID == id
	})
	if !found {
		return types.Task{}, fmt.Errorf("task not found (-id=%d)", id)
	}

	return task, nil
}

func (r *Repository) Update(id uint, input types.Task) (types.Task, error) {
	tasks, err := r.cliDatabase.Read()
	if err != nil {
		return types.Task{}, err
	}

	task, err := r.Read(id)
	if err != nil {
		return types.Task{}, err
	} else if task.DeletedAt != nil {
		return types.Task{}, fmt.Errorf("deleted task cannot be updated (-id=%d)", id)
	}

	if input.Description != "" {
		task.Description = input.Description
	}

	if input.Status != "" {
		task.Status = input.Status
	}

	now := time.Now()
	task.UpdatedAt = &now

	return task, r.cliDatabase.Write(sliceutils.Splice(tasks, id-1, 1, task))
}

func (r *Repository) Delete(id uint) (types.Task, error) {
	tasks, err := r.cliDatabase.Read()
	if err != nil {
		return types.Task{}, err
	}

	task, err := r.Read(id)
	if err != nil {
		return types.Task{}, err
	} else if task.DeletedAt != nil {
		return types.Task{}, fmt.Errorf("task already deleted (-id=%d)", id)
	}

	now := time.Now()
	task.DeletedAt = &now

	return task, r.cliDatabase.Write(sliceutils.Splice(tasks, id-1, 1, task))
}

func (r *Repository) List(
	predicate func(value types.Task, index uint, slice []types.Task) bool,
) ([]types.Task, error) {
	tasks, err := r.cliDatabase.Read()
	if err != nil {
		return nil, err
	}

	if predicate != nil {
		return sliceutils.Filter(tasks, predicate), nil
	}

	return tasks, nil
}

func (r *Repository) Clean() error {
	return r.cliDatabase.Drop()
}
