package line

type Line struct {
	ID         string
	Name       string
	Type       string
	Subnets    []string
	UserID     string
	NextNodeID string
	CreatedAt  int64
	UpdatedAt  int64
	DeletedAt  int64
}

// 判断是否还有下一个节点
func (l *Line) HasNextNode() bool {
	return l.NextNodeID != ""
}
