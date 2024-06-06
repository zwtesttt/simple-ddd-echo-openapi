package repository

import (
	"context"
	"cosslan/internal/domain/entity"
	"cosslan/internal/domain/filter"
)

type NodeRepository interface {
	Find(ctx context.Context, filter *filter.NodeFilter) ([]*entity.Node, error)
	Create(ctx context.Context, node *entity.Node) error
	Get(ctx context.Context, id string) (*entity.Node, error)
	Update(ctx context.Context, node *entity.Node) error
	Delete(ctx context.Context, id string) error
}
