package database

import (
	"database/sql"
	"github.com/meanii/family-tree/model"
	"log"
)

func (d *SqlDatabase) CreateFamilyTreeTable() {
	_, err := d.Database.Exec(`CREATE TABLE IF NOT EXISTS family_tree (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		person1_id INTEGER NOT NULL,
		person2_id INTEGER NOT NULL,
		relationship_id INTEGER NOT NULL,
		FOREIGN KEY (person1_id) REFERENCES person(id),
    	FOREIGN KEY (person2_id) REFERENCES person(id),
    	FOREIGN KEY (relationship_id) REFERENCES relationship(id)
	)`)
	if err != nil {
		log.Fatalf("error creating family tree table: %v", err)
	}
}

// InsertFamilyTree inserts a new row into the family_tree table
func (d *SqlDatabase) InsertFamilyTree(name string, of string, relationshipType string) {

	person1Id := d.GetPerson(model.Person{Name: name}).ID
	if person1Id == 0 {
		log.Fatalf("person '%s' does not exist in the database!", name)
	}

	person2Id := d.GetPerson(model.Person{Name: of}).ID
	if person2Id == 0 {
		log.Fatalf("person '%s' does not exist in the database!", of)
	}

	relationshipId := d.GetRelationship(model.Relationship{Type: relationshipType}).ID
	if relationshipId == 0 {
		log.Fatalf("relationship type '%s' does not exist in the database!\n", relationshipType)
	}

	relationshipExists := d.GetFamilyTree(model.FamilyTree{Person1ID: person1Id, Person2ID: person2Id, RelationshipID: relationshipId}).ID != 0
	if relationshipExists {
		log.Fatalf("the relationship you provided already exists already!")
	}

	_, err := d.Database.Exec(`INSERT INTO family_tree (person1_id, person2_id, relationship_id) VALUES (?, ?, ?)`, person1Id, person2Id, relationshipId)
	if err != nil {
		log.Fatalf("error inserting family tree: %v", err)
	}
}

// GetFamilyTree returns a family tree from the database
func (d *SqlDatabase) GetFamilyTree(familyTreeArgs model.FamilyTree) model.FamilyTree {
	var familyTree model.FamilyTree

	row, err := d.Database.Query(
		`SELECT * FROM family_tree WHERE person1_id = ? AND person2_id = ? AND relationship_id = ?`,
		familyTreeArgs.Person1ID, familyTreeArgs.Person2ID, familyTreeArgs.RelationshipID,
	)
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
		err := row.Scan(&familyTree.ID, &familyTree.Person1ID, &familyTree.Person2ID, &familyTree.RelationshipID)
		if err != nil {
			return model.FamilyTree{}
		}
	}
	return familyTree
}

func (d *SqlDatabase) GetCountOf(relationshipType string, of string) int {
	relationshipId := d.GetRelationship(model.Relationship{Type: relationshipType}).ID
	if relationshipId == 0 {
		log.Fatalf("relationship type '%s' does not exist in the database!\n", relationshipType)
	}

	personId := d.GetPerson(model.Person{Name: of}).ID
	if personId == 0 {
		log.Fatalf("person '%s' does not exist in the database!", of)
	}

	var count int
	row, err := d.Database.Query(`SELECT COUNT(*) FROM family_tree WHERE relationship_id = ? AND person2_id = ?`, relationshipId, personId)
	if err != nil {
		log.Fatalf("error getting count of: %v", err)
	}

	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatalf("error closing rows: %v", err)
		}
	}(row)

	for row.Next() {
		err := row.Scan(&count)
		if err != nil {
			log.Fatalf("error scanning row: %v", err)
		}
	}

	return count
}

// GetName returns the name of a person in a relationship
func (d *SqlDatabase) GetName(relationship string, of string) string {
	relationshipId := d.GetRelationship(model.Relationship{Type: relationship}).ID
	if relationshipId == 0 {
		log.Fatalf("relationship type '%s' does not exist in the database!\n", relationship)
	}

	personId := d.GetPerson(model.Person{Name: of}).ID
	if personId == 0 {
		log.Fatalf("person '%s' does not exist in the database!", of)
	}

	var person1Id int
	row, err := d.Database.Query(`SELECT person1_id FROM family_tree WHERE relationship_id = ? AND person2_id = ?`, relationshipId, personId)
	if err != nil {
		log.Fatalf("error getting name: %v", err)
	}

	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatalf("error closing rows: %v", err)
		}
	}(row)

	for row.Next() {
		err := row.Scan(&person1Id)
		if err != nil {
			log.Fatalf("error scanning row: %v", err)
		}
	}

	return d.GetPerson(model.Person{ID: person1Id}).Name
}
