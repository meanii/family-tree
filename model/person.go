package model

type Person struct {
	ID         int    // unique identifier, primary key, autoincrement
	Name       string // "John Doe", contains first and last name
	Gender     string // "M" or "F"
	FamilyRoot bool   // true if this person is the root of the family tree
}
