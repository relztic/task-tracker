package json

import (
	"encoding/json"
	"os"

	"github.com/relztic/task-tracker/shared"
	"github.com/relztic/task-tracker/shared/domain"
	"github.com/relztic/task-tracker/shared/types"
)

type Database struct{}

var _ domain.CLIDatabase = &Database{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Read() ([]types.Task, error) {
	var tasks []types.Task

	file, err := os.Open(shared.Config.TasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []types.Task{}, nil
		}

		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (d *Database) Write(tasks []types.Task) error {
	file, err := os.Create(shared.Config.TasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(&tasks); err != nil {
		return err
	}

	return nil
}

func (d *Database) Drop() error {
	if _, err := os.Stat(shared.Config.TasksFile); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(shared.Config.TasksFile)
}
