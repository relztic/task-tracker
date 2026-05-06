package main

import (
	"flag"
	"fmt"
	"os"

	sliceutils "github.com/relztic/go-utils/v2/slice"
	"github.com/relztic/task-tracker/modules/cli"
	"github.com/relztic/task-tracker/shared/database/json"
	"github.com/relztic/task-tracker/shared/types"
)

func main() {
	var cmd types.Command

	validActions := sliceutils.Join(types.CmdAction("").Values(), "|")
	flag.StringVar((*string)(&cmd.Action), "action", "", fmt.Sprintf("action to perform [%s]", validActions))
	flag.UintVar(&cmd.ID, "id", 0, "id of the task to read/update/delete")
	flag.StringVar(&cmd.Description, "description", "", "description of the task")
	flag.Parse()

	jsonDatabase := json.NewDatabase()
	cliRepository := cli.NewRepository(jsonDatabase)
	cliInteractor := cli.NewInteractor(cliRepository)
	cliHandler := cli.NewHandler(cliInteractor)
	cliRouter := cli.NewRouter(cliHandler)

	msg, err := cliRouter.Route(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	} else if msg != "" {
		fmt.Fprintf(os.Stdout, "success: %s\n", msg)
	}

	os.Exit(0)
}
