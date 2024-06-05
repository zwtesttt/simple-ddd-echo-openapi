package po

import "go.mongodb.org/mongo-driver/bson/primitive"

type Node struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	IP        string             `bson:"ip"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty"`
	Location  string             `bson:"location"`
	CreatedAt int64              `bson:"created_at"`
	UpdatedAt int64              `bson:"updated_at"`
	DeletedAt int64              `bson:"deleted_at"`
}
