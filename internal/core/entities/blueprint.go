package entities

type Blueprint struct {
	ID          string            `bson:"_id" json:"id"`
	Name        string            `bson:"name" json:"name"`
	Description string            `bson:"description" json:"description"`
	Parameters  map[string]string `bson:"parameters" json:"parameters"`
	Category    string            `bson:"category" json:"category"` // frontend or backend
	Version     string            `bson:"version" json:"version"`
}
//crossplane blueprint entity