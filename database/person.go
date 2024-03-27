package database

import (
	"database/sql"
	"log"
	"slices"

	"github.com/meanii/family-tree/model"
)

// CreatePersonTable creates the person table in the database
func (d *SqlDatabase) CreatePersonTable() {
	_, err := d.Database.Exec(`CREATE TABLE IF NOT EXISTS person (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		gender TEXT DEFAULT "M",
		family_root BOOLEAN DEFAULT FALSE
	)`)
	if err != nil {
		panic(err)
	}
}

// InsertPerson inserts a new row into the person table
func (d *SqlDatabase) InsertPerson(name string, gender string, family_root bool) {
	validGenders := []string{"M", "F"}
	if !slices.Contains(validGenders, gender) {
		log.Fatalf("gender must be one of %v", validGenders)
	}

	nameExists := d.GetPerson(model.Person{Name: name}).ID != 0
	if nameExists {
		log.Fatalf(
			"the name you provided already exists in the database! We cannot have two people with the same name.",
		)
	}

	_, err := d.Database.Exec(
		`INSERT INTO person (name, gender, family_root) VALUES (?, ?, ?)`,
		name,
		gender,
		family_root,
	)
	if err != nil {
		panic(err)
	}
}

// GetPerson returns a person from the database
func (d *SqlDatabase) GetPerson(personArgs model.Person) model.Person {
	var person model.Person
	row, err := d.getPersonQuery(personArgs)
	if err != nil {
		log.Fatalf("error getting person: %v", err)
	}

	if row == nil {
		return model.Person{ID: 0}
	}

	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatalf("error closing rows: %v", err)
		}
	}(row)

	for row.Next() {
		err := row.Scan(&person.ID, &person.Name, &person.Gender, &person.FamilyRoot)
		if err != nil {
			return model.Person{}
		}
	}
	return person
}

// GetPeople returns all people from the database
func (d *SqlDatabase) GetPeople() []model.Person {
	var people []model.Person
	rows, err := d.Database.Query(`SELECT * FROM person`)
	if err != nil {
		log.Fatalf("error getting people: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("error closing rows: %v", err)
		}
	}(rows)
	for rows.Next() {
		var person model.Person
		err := rows.Scan(&person.ID, &person.Name, &person.Gender, &person.FamilyRoot)
		if err != nil {
			log.Fatalf("error scanning people: %v", err)
		}
		people = append(people, person)
	}
	return people
}

// getPersonQuery returns a sql.Rows object, or an error
func (d *SqlDatabase) getPersonQuery(personArgs model.Person) (*sql.Rows, error) {
	if personArgs.ID != 0 {
		return d.Database.Query(`SELECT * FROM person WHERE id = ?`, personArgs.ID)
	}
	if personArgs.Name != "" {
		return d.Database.Query(`SELECT * FROM person WHERE name = ?`, personArgs.Name)
	}
	return nil, nil
}
