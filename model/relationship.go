package model

type Relationship struct {
	ID   int    // unique identifier, primary key, autoincrement
	Type string // father, mother, son, daughter, wife, husband, etc. (relationship type)
}
