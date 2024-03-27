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

		people := database.Database.GetNames(relationship, of)

		// This is the case where we don't have any person with the relationship
		if len(people) == 0 {
			fmt.Printf("No one found with relationship %s to %s\n", relationship, of)
			return
		}

		// This is the case where we have only one person with the relationship
		if len(people) == 1 {
			fmt.Printf("Name of the person: %s\n", people[0].Name)
			return
		}

		// This is the case where we have multiple people with the same relationship
		fmt.Printf("Names of the people:\n")
		for _, person := range people {
			fmt.Printf("- %s\n", person.Name)
		}
		return

	},
}

func init() {
	RootCmd.AddCommand(nameCmd)
	nameCmd.Flags().StringP("relationship", "r", "", "Relationship type")
	_ = nameCmd.MarkFlagRequired("relationship")
	nameCmd.Flags().StringP("of", "o", "", "Name of the other person")
	_ = nameCmd.MarkFlagRequired("of")
}
