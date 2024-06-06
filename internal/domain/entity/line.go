package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Line struct {
	ID         primitive.ObjectID
	Name       string
	Type       string
	Subnets    []string
	UserID     primitive.ObjectID
	NextNodeID primitive.ObjectID
	CreatedAt  int64
	UpdatedAt  int64
	DeletedAt  int64
}

// 判断是否还有下一个节点
func (l *Line) HasNextNode() bool {
	return l.NextNodeID != primitive.NilObjectID
}
