package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type SqlDatabase struct {
	Database *sql.DB
	Name     string
}

// Database, a pointer to the SqlDatabase struct
var Database *SqlDatabase

// NewDatabase creates a new database connection
func NewDatabase(name string) {
	Database = newDatabase(name)
	Database.createTables() // create the tables in the database, if they don't already exist
}

// newDatabase creates a new database connection
// and returns a pointer to the Database struct
func newDatabase(name string) *SqlDatabase {
	database, err := sql.Open("sqlite3", fmt.Sprintf("./%s.db", name))
	if err != nil {
		panic(err)
	}
	return &SqlDatabase{Name: name, Database: database}
}

// CreateTables creates the person and relationship tables in the database
func (d *SqlDatabase) createTables() {
	d.CreatePersonTable()
	d.CreateRelationshipTable()
	d.CreateFamilyTreeTable()
}

// Close closes the database connection
func (d *SqlDatabase) Close() {
	err := d.Database.Close()
	if err != nil {
		return
	}
}
