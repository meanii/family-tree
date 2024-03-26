package add

import (
	"fmt"
	"github.com/meanii/family-tree/database"
	"github.com/spf13/cobra"
)

// relationshipCmd represents the relationship command
var relationshipCmd = &cobra.Command{
	Use:   "relationship",
	Short: "create new relationship in the family tree",
	Long: `Create a new relationship in the family tree. For example:
family-tree add relationship --type=father
`,
	Run: func(cmd *cobra.Command, args []string) {
		relationshipType, _ := cmd.Flags().GetString("type")
		// Add the relationship to the database
		database.Database.InsertRelationship(relationshipType)
		fmt.Printf("New relationship added: '%s'\n", relationshipType)
	},
}

func init() {
	addCmd.AddCommand(relationshipCmd)
	relationshipCmd.Flags().StringP("type", "t", "", "Type of the relationship")
	_ = relationshipCmd.MarkFlagRequired("type")
}
