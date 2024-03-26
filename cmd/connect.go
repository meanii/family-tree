package cmd

import (
	"fmt"
	"github.com/meanii/family-tree/database"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to the family tree database",
	Long: `Connect to the family tree database. For example:
family-tree connect --name="Name 1" --relationship="father" --of="Name 2"`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flags().Lookup("name").Value.String()
		relationshipType := cmd.Flags().Lookup("relationship").Value.String()
		of := cmd.Flags().Lookup("of").Value.String()

		// Add the relationship to the database
		database.Database.InsertFamilyTree(name, of, relationshipType)
		fmt.Printf("successfully connected '%s' as '%s' of '%s'.\n", name, relationshipType, of)
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringP("name", "n", "", "Name of the person")
	_ = connectCmd.MarkFlagRequired("name")
	connectCmd.Flags().StringP("relationship", "r", "", "Relationship type")
	_ = connectCmd.MarkFlagRequired("relationship")
	connectCmd.Flags().StringP("of", "o", "", "Name of the other person")
	_ = connectCmd.MarkFlagRequired("of")
}
