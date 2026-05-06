package cli

import (
	"fmt"

	"github.com/relztic/task-tracker/shared/domain"
	"github.com/relztic/task-tracker/shared/types"
	"github.com/relztic/task-tracker/shared/utils"
)

type Handler struct {
	cliInteractor domain.CLIInteractor
}

var _ domain.CLIHandler = &Handler{}

func NewHandler(cliInteractor domain.CLIInteractor) *Handler {
	return &Handler{
		cliInteractor: cliInteractor,
	}
}

func (h *Handler) Create(description string) (string, error) {
	if description == "" {
		return "", fmt.Errorf("-description cannot be empty")
	}

	task, err := h.cliInteractor.Create(description)
	if err != nil {
		return "", fmt.Errorf("creating task: %v", err)
	}

	return fmt.Sprintf("task created (-id=%d)", task.ID), nil
}

func (h *Handler) Read(id uint) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("-id cannot be empty")
	}

	task, err := h.cliInteractor.Read(id)
	if err != nil {
		return "", fmt.Errorf("reading task: %v", err)
	}

	writer := utils.ListTasks([]types.Task{task})
	writer.Flush()

	return "", nil
}

func (h *Handler) Update(id uint, description string) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("-id cannot be empty")
	}

	if description == "" {
		return "", fmt.Errorf("-description cannot be empty")
	}

	task, err := h.cliInteractor.Update(id, description, "")
	if err != nil {
		return "", fmt.Errorf("updating task: %v", err)
	}

	return fmt.Sprintf("task updated (-id=%d)", task.ID), nil
}

func (h *Handler) Delete(id uint) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("-id cannot be empty")
	}

	task, err := h.cliInteractor.Delete(id)
	if err != nil {
		return "", fmt.Errorf("deleting task: %v", err)
	}

	return fmt.Sprintf("task deleted (-id=%d)", task.ID), nil
}

func (h *Handler) MarkAsTodo(id uint) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("-id cannot be empty")
	}

	task, err := h.cliInteractor.Update(id, "", types.TaskStatusTodo)
	if err != nil {
		return "", fmt.Errorf("marking task as %s: %v", types.TaskStatusTodo, err)
	}

	return fmt.Sprintf("task marked as %s (-id=%d)", types.TaskStatusTodo, task.ID), nil
}

func (h *Handler) MarkAsInProgress(id uint) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("-id cannot be empty")
	}

	task, err := h.cliInteractor.Update(id, "", types.TaskStatusInProgress)
	if err != nil {
		return "", fmt.Errorf("marking task as %s: %v", types.TaskStatusInProgress, err)
	}

	return fmt.Sprintf("task marked as %s (-id=%d)", types.TaskStatusInProgress, task.ID), nil
}

func (h *Handler) MarkAsDone(id uint) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("-id cannot be empty")
	}

	task, err := h.cliInteractor.Update(id, "", types.TaskStatusDone)
	if err != nil {
		return "", fmt.Errorf("marking task as %s: %v", types.TaskStatusDone, err)
	}

	return fmt.Sprintf("task marked as %s (-id=%d)", types.TaskStatusDone, task.ID), nil
}

func (h *Handler) List() (string, error) {
	tasks, err := h.cliInteractor.List()
	if err != nil {
		return "", fmt.Errorf("listing tasks: %v", err)
	}

	writer := utils.ListTasks(tasks)
	writer.Flush()

	return "", nil
}

func (h *Handler) ListTodo() (string, error) {
	tasks, err := h.cliInteractor.ListTodo()
	if err != nil {
		return "", fmt.Errorf("listing %s tasks: %v", types.TaskStatusTodo, err)
	}

	writer := utils.ListTasks(tasks)
	writer.Flush()

	return "", nil
}

func (h *Handler) ListInProgress() (string, error) {
	tasks, err := h.cliInteractor.ListInProgress()
	if err != nil {
		return "", fmt.Errorf("listing %s tasks: %v", types.TaskStatusInProgress, err)
	}

	writer := utils.ListTasks(tasks)
	writer.Flush()

	return "", nil
}

func (h *Handler) ListDone() (string, error) {
	tasks, err := h.cliInteractor.ListDone()
	if err != nil {
		return "", fmt.Errorf("listing %s tasks: %v", types.TaskStatusDone, err)
	}

	writer := utils.ListTasks(tasks)
	writer.Flush()

	return "", nil
}

func (h *Handler) ListDeleted() (string, error) {
	tasks, err := h.cliInteractor.ListDeleted()
	if err != nil {
		return "", fmt.Errorf("listing deleted tasks: %v", err)
	}

	writer := utils.ListTasks(tasks)
	writer.Flush()

	return "", nil
}

func (h *Handler) ListAll() (string, error) {
	tasks, err := h.cliInteractor.ListAll()
	if err != nil {
		return "", fmt.Errorf("listing all tasks: %v", err)
	}

	writer := utils.ListTasks(tasks)
	writer.Flush()

	return "", nil
}

func (h *Handler) Clean() (string, error) {
	if err := h.cliInteractor.Clean(); err != nil {
		return "", fmt.Errorf("cleaning tasks: %v", err)
	}

	return "tasks cleaned", nil
}
