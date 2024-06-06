package service

import (
	"context"
	"cosslan/config"
	"cosslan/internal/app/dto"
	"cosslan/internal/infra/persistence"
	"fmt"
)

var _ NodeService = &NodeServiceImpl{}

type NodeService interface {
	CreateNode(ctx context.Context, dto dto.CreateNodeDTO) error
}

type NodeServiceImpl struct {
	repo *persistence.Repositories
}

func NewNodeServiceImpl(cfg config.Config) NodeService {
	return &NodeServiceImpl{repo: persistence.GetRepositories(cfg)}
}

func (n *NodeServiceImpl) CreateNode(ctx context.Context, dto dto.CreateNodeDTO) error {
	fmt.Println("创建节点")
	return nil
}
