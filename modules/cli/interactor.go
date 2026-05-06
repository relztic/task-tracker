package cli

import (
	"github.com/relztic/task-tracker/shared/domain"
	"github.com/relztic/task-tracker/shared/types"
)

type Interactor struct {
	cliRepository domain.CLIRepository
}

var _ domain.CLIInteractor = &Interactor{}

func NewInteractor(cliRepository domain.CLIRepository) *Interactor {
	return &Interactor{
		cliRepository: cliRepository,
	}
}

func (i *Interactor) Create(description string) (types.Task, error) {
	input := types.Task{
		Description: description,
		Status:      types.TaskStatusTodo,
	}

	task, err := i.cliRepository.Create(input)
	if err != nil {
		return types.Task{}, err
	}

	return task, nil
}

func (i *Interactor) Read(id uint) (types.Task, error) {
	return i.cliRepository.Read(id)
}

func (i *Interactor) Update(id uint, description string, status types.TaskStatus) (types.Task, error) {
	input := types.Task{
		Description: description,
		Status:      status,
	}

	task, err := i.cliRepository.Update(id, input)
	if err != nil {
		return types.Task{}, err
	}

	return task, nil
}

func (i *Interactor) Delete(id uint) (types.Task, error) {
	task, err := i.cliRepository.Delete(id)
	if err != nil {
		return types.Task{}, err
	}

	return task, nil
}

func (i *Interactor) List() ([]types.Task, error) {
	return i.cliRepository.List(func(task types.Task, _ uint, _ []types.Task) bool {
		return task.DeletedAt == nil
	})
}

func (i *Interactor) ListTodo() ([]types.Task, error) {
	return i.cliRepository.List(func(task types.Task, _ uint, _ []types.Task) bool {
		return task.Status == types.TaskStatusTodo
	})
}

func (i *Interactor) ListInProgress() ([]types.Task, error) {
	return i.cliRepository.List(func(task types.Task, _ uint, _ []types.Task) bool {
		return task.Status == types.TaskStatusInProgress
	})
}

func (i *Interactor) ListDone() ([]types.Task, error) {
	return i.cliRepository.List(func(task types.Task, _ uint, _ []types.Task) bool {
		return task.Status == types.TaskStatusDone
	})
}

func (i *Interactor) ListDeleted() ([]types.Task, error) {
	return i.cliRepository.List(func(task types.Task, _ uint, _ []types.Task) bool {
		return task.DeletedAt != nil
	})
}

func (i *Interactor) ListAll() ([]types.Task, error) {
	return i.cliRepository.List(nil)
}

func (i *Interactor) Clean() error {
	return i.cliRepository.Clean()
}
