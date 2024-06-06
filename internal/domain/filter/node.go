package filter

import "go.mongodb.org/mongo-driver/bson"

// NodeFilter 是用于查询节点的过滤条件
type NodeFilter struct {
	ID       string
	Name     string
	IP       string
	UserID   string
	Location string
}

// ToBSON 将 NodeFilter 转换为 MongoDB 查询文档
func (filter *NodeFilter) ToBSON() bson.M {
	query := bson.M{}

	if filter.ID != "" {
		query["_id"] = filter.ID
	}

	if filter.Name != "" {
		query["name"] = filter.Name
	}

	if filter.IP != "" {
		query["ip"] = filter.IP
	}

	if filter.UserID != "" {
		query["user_id"] = filter.UserID
	}

	if filter.Location != "" {
		query["location"] = filter.Location
	}

	return query
}
