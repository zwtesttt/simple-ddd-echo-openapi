package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Node struct {
	ID        primitive.ObjectID
	Name      string
	IP        string
	UserID    primitive.ObjectID
	Location  string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}
