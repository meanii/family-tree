package main

import (
	"github.com/meanii/family-tree/cmd"
	_ "github.com/meanii/family-tree/cmd/add"
	"github.com/meanii/family-tree/database"
)

func main() {
	// prepare database connection, before executing the command
	database.NewDatabase("family-tree")

	// Close the database connection after the command is executed
	defer database.Database.Close()

	// Execute the command
	cmd.Execute()
}
