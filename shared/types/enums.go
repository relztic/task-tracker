package types

type Enum[T ~string] interface {
	Values() []T
}

type CmdAction string

var _ Enum[CmdAction] = CmdAction("")

const (
	CmdActionCreate         = CmdAction("create")
	CmdActionRead           = CmdAction("read")
	CmdActionUpdate         = CmdAction("update")
	CmdActionDelete         = CmdAction("delete")
	CmdActionMarkTodo       = CmdAction("mark-todo")
	CmdActionMarkInProgress = CmdAction("mark-in-progress")
	CmdActionMarkDone       = CmdAction("mark-done")
	CmdActionList           = CmdAction("list")
	CmdActionListTodo       = CmdAction("list-todo")
	CmdActionListInProgress = CmdAction("list-in-progress")
	CmdActionListDone       = CmdAction("list-done")
	CmdActionListDeleted    = CmdAction("list-deleted")
	CmdActionListAll        = CmdAction("list-all")
	CmdActionClean          = CmdAction("clean")
)

func (e CmdAction) Values() []CmdAction {
	return []CmdAction{
		CmdActionCreate,
		CmdActionRead,
		CmdActionUpdate,
		CmdActionDelete,
		CmdActionMarkTodo,
		CmdActionMarkInProgress,
		CmdActionMarkDone,
		CmdActionList,
		CmdActionListTodo,
		CmdActionListInProgress,
		CmdActionListDone,
		CmdActionListDeleted,
		CmdActionListAll,
		CmdActionClean,
	}
}

type TaskStatus string

var _ Enum[TaskStatus] = TaskStatus("")

const (
	TaskStatusTodo       = TaskStatus("TODO")
	TaskStatusInProgress = TaskStatus("IN_PROGRESS")
	TaskStatusDone       = TaskStatus("DONE")
)

func (e TaskStatus) Values() []TaskStatus {
	return []TaskStatus{
		TaskStatusTodo,
		TaskStatusInProgress,
		TaskStatusDone,
	}
}
