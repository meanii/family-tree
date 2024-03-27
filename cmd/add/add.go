package add

import (
	"github.com/spf13/cobra"

	"github.com/meanii/family-tree/cmd"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new person or relationship to the family tree",
	Long: `Add a new person or relationship to the family tree. For example:
family-tree add person --name="Someone Name"
family-tree add relationship --type="brother"
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(addCmd)
}
