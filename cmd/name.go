package cmd

import (
	"fmt"

	"github.com/meanii/family-tree/database"
	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "name of the person",
	Long: `Name of the person. For example:
family-tree name --relationship="father" --of="Name 2"`,
	Run: func(cmd *cobra.Command, args []string) {
		relationship := cmd.Flags().Lookup("relationship").Value.String()
		of := cmd.Flags().Lookup("of").Value.String()

		name := database.Database.GetName(relationship, of)
		fmt.Printf("Name of the person: %s\n", name)
	},
}

func init() {
	RootCmd.AddCommand(nameCmd)
	nameCmd.Flags().StringP("relationship", "r", "", "Relationship type")
	_ = nameCmd.MarkFlagRequired("relationship")
	nameCmd.Flags().StringP("of", "o", "", "Name of the other person")
	_ = nameCmd.MarkFlagRequired("of")
}
