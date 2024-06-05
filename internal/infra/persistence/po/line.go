package po

import "go.mongodb.org/mongo-driver/bson/primitive"

type Line struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Type       string             `bson:"type"`
	Subnets    []string           `bson:"subnets"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty"`
	NextNodeID primitive.ObjectID `bson:"next_node_id,omitempty"`
	CreatedAt  int64              `bson:"created_at"`
	UpdatedAt  int64              `bson:"updated_at"`
	DeletedAt  int64              `bson:"deleted_at"`
}
