package reciprocal_relationship

// Define a map for relationships and their reciprocal counterparts
var ReciprocalRelationships = map[string]map[string]string{
	"father": {
		"M": "son",
		"F": "daughter",
	},
	"mother": {
		"M": "son",
		"F": "daughter",
	},
	"son": {
		"M": "father",
		"F": "mother",
	},
	"daughter": {
		"M": "father",
		"F": "mother",
	},
	"brother": {
		"M": "bother",
		"F": "sister",
	},
	"sister": {
		"M": "brother",
		"F": "sister",
	},
	"husband": {
		"F": "wife",
	},
	"wife": {
		"M": "husband",
	},
	"uncle": {
		"M": "niece",
		"F": "nephew",
	},
	"aunt": {
		"M": "nephew",
		"F": "niece",
	},
	"nephew": {
		"M": "uncle",
		"F": "aunt",
	},
	"niece": {
		"M": "uncle",
		"F": "aunt",
	},
	"grandfather": {
		"M": "grandson",
		"F": "granddaughter",
	},
	"grandmother": {
		"M": "grandson",
		"F": "granddaughter",
	},
	"grandson": {
		"M": "grandfather",
		"F": "grandmother",
	},
	"granddaughter": {
		"M": "grandfather",
		"F": "grandmother",
	},
}
