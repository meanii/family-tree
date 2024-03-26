package cmd

import (
	"fmt"
	"github.com/meanii/family-tree/database"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "count the number of people in the family tree",
	Long: `Count the number of people in the family tree. For example:
family-tree count --type=father --of="Name 2"`,
	Run: func(cmd *cobra.Command, args []string) {
		relationshipType, _ := cmd.Flags().GetString("type")
		of, _ := cmd.Flags().GetString("of")

		// Count the number of people in the family tree
		count := database.Database.GetCountOf(relationshipType, of)
		if count == 0 {
			fmt.Printf("No %s are associated with %s.\n", relationshipType, of)
			return
		}
		fmt.Printf("Found %d people with the relationship type '%s' for '%s'.\n", count, relationshipType, of)
	},
}

func init() {
	RootCmd.AddCommand(countCmd)
	countCmd.Flags().StringP("type", "t", "", "Type of the relationship")
	_ = countCmd.MarkFlagRequired("type")
	countCmd.Flags().StringP("of", "o", "", "Name of the other person")
	_ = countCmd.MarkFlagRequired("of")
}
