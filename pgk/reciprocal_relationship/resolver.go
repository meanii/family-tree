package reciprocal_relationship

import (
	"github.com/meanii/family-tree/model"
)

// ReciprocalRelationships is a map of relationship types and their reciprocal relationships
func GetReciprocalRelationship(
	person2 model.Person,
	relationship model.Relationship,
) model.Relationship {
	if relMap, ok := ReciprocalRelationships[relationship.Type]; ok {
		if rel, ok := relMap[person2.Gender]; ok {
			return model.Relationship{Type: rel}
		}
	}

	return model.Relationship{Type: ""}
}
