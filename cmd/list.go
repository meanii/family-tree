package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/meanii/family-tree/database"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all relationships or persons",
	Long: `list all relationships or persons. For example:
family-tree list --relationship
family-tree list --person`,
	Run: func(cmd *cobra.Command, args []string) {
		relationship, _ := cmd.Flags().GetBool("relationship")
		person, _ := cmd.Flags().GetBool("person")

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		if relationship {
			tbl := table.New("ID", "Type")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, relationship := range database.Database.GetRelationships() {
				tbl.AddRow(relationship.ID, relationship.Type)
			}
			tbl.Print()
			return
		}

		if person {
			tbl := table.New("ID", "Name", "Gender", "Family Root")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, person := range database.Database.GetPeople() {
				tbl.AddRow(person.ID, person.Name, person.Gender, person.FamilyRoot)
			}
			tbl.Print()
			return
		}

		fmt.Println("Please specify either --relationship or --person")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("relationship", "r", false, "List all relationships")
	listCmd.Flags().BoolP("person", "p", false, "List all persons")
}
