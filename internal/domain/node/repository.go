package node

import "context"

type NodeRepository interface {
	Find(ctx context.Context, filter *NodeFilter) ([]*Node, error)
	Create(ctx context.Context, node *Node) error
	Get(ctx context.Context, id string) (*Node, error)
	Update(ctx context.Context, node *Node) error
	Delete(ctx context.Context, id string) error
}
