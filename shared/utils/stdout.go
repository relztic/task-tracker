package utils

import (
	"fmt"
	"os"
	"text/tabwriter"

	formatutils "github.com/relztic/go-utils/v2/format"
	"github.com/relztic/task-tracker/shared/types"
)

func ListTasks(tasks []types.Task) *tabwriter.Writer {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(tabWriter, " ID\tCREATED AT\tUPDATED AT\tDELETED AT\tDESCRIPTION\tSTATUS\t")
	fmt.Fprintln(tabWriter, "---\t----------\t----------\t----------\t-----------\t------\t")

	for _, task := range tasks {
		formattedCreatedAt := formatutils.Datetime(&task.CreatedAt)
		formattedUpdatedAt := formatutils.Datetime(task.UpdatedAt)
		formattedDeletedAt := formatutils.Datetime(task.DeletedAt)
		fmt.Fprintf(tabWriter, "%d\t%s\t%s\t%s\t%s\t%s\t\n", task.ID, formattedCreatedAt, formattedUpdatedAt, formattedDeletedAt, task.Description, task.Status)
	}

	return tabWriter
}
