package add

import (
	"fmt"

	"github.com/meanii/family-tree/database"
	"github.com/spf13/cobra"
)

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "a new person to the family tree",
	Long: `Add a new person to the family tree. For example:
family-tree add person --name="Someone Name"
`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flags().Lookup("name").Value.String()
		gender := cmd.Flags().Lookup("gender").Value.String()
		familyRoot, _ := cmd.Flags().GetBool("family_root")

		// Add the person to the database
		database.Database.InsertPerson(name, gender, familyRoot)
		fmt.Printf("New person added: '%s'\n", name)
	},
}

func init() {
	addCmd.AddCommand(personCmd)

	// Here you will define your flags and configuration settings.
	personCmd.Flags().StringP("name", "n", "", "Name of the person")
	_ = personCmd.MarkFlagRequired("name")
	personCmd.Flags().StringP("gender", "g", "M", "Gender of the person")
	personCmd.Flags().BoolP("family_root", "r", false, "Is the person the root of the family tree")
}
