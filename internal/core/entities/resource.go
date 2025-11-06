package entities

type Resource struct {
		 ID          string                 `bson:"_id" json:"id"`
		 Name        string                 `bson:"name" json:"name"`
		 BlueprintID string                 `bson:"blueprint_id" json:"blueprintId"`
		 Description string                 `bson:"description" json:"description"`
		 Status      string                 `bson:"status" json:"status"`
		 Metadata    map[string]interface{} `bson:"metadata" json:"metadata"`
		 CreatedAt   int64                  `bson:"created_at" json:"createdAt"`
		 UpdatedAt   int64                  `bson:"updated_at" json:"updatedAt"`
}