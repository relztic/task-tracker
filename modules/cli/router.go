package cli

import (
	"fmt"

	sliceutils "github.com/relztic/go-utils/v2/slice"
	streamutils "github.com/relztic/go-utils/v2/stream"
	"github.com/relztic/task-tracker/shared/domain"
	"github.com/relztic/task-tracker/shared/types"
)

type Router struct {
	cliHandler domain.CLIHandler
}

var _ domain.CLIRouter = &Router{}

func NewRouter(cliHandler domain.CLIHandler) Router {
	return Router{
		cliHandler: cliHandler,
	}
}

func (r *Router) Route(cmd types.Command) (string, error) {
	var msg string
	var err error

	switch cmd.Action {
	case types.CmdActionCreate:
		msg, err = r.cliHandler.Create(cmd.Description)
	case types.CmdActionRead:
		msg, err = r.cliHandler.Read(cmd.ID)
	case types.CmdActionUpdate:
		msg, err = r.cliHandler.Update(cmd.ID, cmd.Description)
	case types.CmdActionDelete:
		if streamutils.Confirm("are you sure?") {
			msg, err = r.cliHandler.Delete(cmd.ID)
		} else {
			msg, err = "", fmt.Errorf("action canceled")
		}
	case types.CmdActionMarkTodo:
		msg, err = r.cliHandler.MarkAsTodo(cmd.ID)
	case types.CmdActionMarkInProgress:
		msg, err = r.cliHandler.MarkAsInProgress(cmd.ID)
	case types.CmdActionMarkDone:
		msg, err = r.cliHandler.MarkAsDone(cmd.ID)
	case types.CmdActionList:
		msg, err = r.cliHandler.List()
	case types.CmdActionListTodo:
		msg, err = r.cliHandler.ListTodo()
	case types.CmdActionListInProgress:
		msg, err = r.cliHandler.ListInProgress()
	case types.CmdActionListDone:
		msg, err = r.cliHandler.ListDone()
	case types.CmdActionListDeleted:
		msg, err = r.cliHandler.ListDeleted()
	case types.CmdActionListAll:
		msg, err = r.cliHandler.ListAll()
	case types.CmdActionClean:
		if streamutils.Confirm("are you sure?") {
			msg, err = r.cliHandler.Clean()
		} else {
			msg, err = "", fmt.Errorf("action canceled")
		}
	default:
		validActions := sliceutils.Join(types.CmdAction("").Values(), "|")
		msg, err = "", fmt.Errorf("invalid action, please use: [%s]", validActions)
	}

	return msg, err
}
