package po

type Line struct {
	ID         string   `bson:"_id,omitempty"`
	Name       string   `bson:"name"`
	Type       string   `bson:"type"`
	Subnets    []string `bson:"subnets"`
	UserID     string   `bson:"user_id,omitempty"`
	NextNodeID string   `bson:"next_node_id,omitempty"`
	CreatedAt  int64    `bson:"created_at"`
	UpdatedAt  int64    `bson:"updated_at"`
	DeletedAt  int64    `bson:"deleted_at"`
}
