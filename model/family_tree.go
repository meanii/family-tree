package model

type FamilyTree struct {
	ID             int // unique identifier, primary key, autoincrement
	Person1ID      int // foreign key to person
	Person2ID      int // foreign key to person
	RelationshipID int // foreign key to relationship
}
