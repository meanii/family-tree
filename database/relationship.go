package database

import (
	"database/sql"
	"log"

	"github.com/meanii/family-tree/model"
)

func (d *SqlDatabase) CreateRelationshipTable() {
	_, err := d.Database.Exec(`CREATE TABLE IF NOT EXISTS relationship (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT NOT NULL
	)`)
	if err != nil {
		panic(err)
	}
}

// InsertRelationship inserts a new row into the relationship table
func (d *SqlDatabase) InsertRelationship(relationshipType string) {
	// check if the relationship type already exists in the database, if it does, exit the program
	relationshipTypeExists := d.GetRelationship(
		model.Relationship{Type: relationshipType, ID: 0},
	).ID == 0
	if relationshipTypeExists != true {
		log.Fatalf("the relationship type you provided already exists in the database!")
	}

	_, err := d.Database.Exec(`INSERT INTO relationship (type) VALUES (?)`, relationshipType)
	if err != nil {
		log.Fatalf("error inserting relationship: %v", err)
	}
}

// GetRelationship returns a relationship from the database
func (d *SqlDatabase) GetRelationship(relationshipArgs model.Relationship) model.Relationship {
	var relationship model.Relationship

	row, err := d.getRelationshipQuery(relationshipArgs)
	if err != nil {
		panic(err)
	}

	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatalf("error closing rows: %v", err)
		}
	}(row)

	for row.Next() {
		err := row.Scan(&relationship.ID, &relationship.Type)
		if err != nil {
			return model.Relationship{}
		}
	}
	return relationship
}

// GetRelationships returns all relationships from the database
func (d *SqlDatabase) GetRelationships() []model.Relationship {
	var relationships []model.Relationship
	rows, err := d.Database.Query(`SELECT id, type FROM relationship`)
	if err != nil {
		log.Fatalf("error getting relationships: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("error closing rows: %v", err)
		}
	}(rows)
	for rows.Next() {
		var relationship model.Relationship
		err := rows.Scan(&relationship.ID, &relationship.Type)
		if err != nil {
			log.Fatalf("error scanning relationships: %v", err)
		}
		relationships = append(relationships, relationship)
	}
	return relationships
}

// getQuery returns a sql.Rows object
func (d *SqlDatabase) getRelationshipQuery(relationshipArgs model.Relationship) (*sql.Rows, error) {
	if relationshipArgs.ID != 0 {
		return d.Database.Query(
			`SELECT id, type FROM relationship WHERE id = ?`,
			relationshipArgs.ID,
		)
	}

	if relationshipArgs.Type != "" {
		return d.Database.Query(
			`SELECT id, type FROM relationship WHERE type = ?`,
			relationshipArgs.Type,
		)
	}

	return nil, nil
}
