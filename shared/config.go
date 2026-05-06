package shared

import configutils "github.com/relztic/go-utils/v2/config"

type EnvConfig struct {
	TasksFile string
}

var Config *EnvConfig

func init() {
	Config = &EnvConfig{
		TasksFile: configutils.GetString("TASKS_FILE", "/tmp/tasks.json"),
	}
}
