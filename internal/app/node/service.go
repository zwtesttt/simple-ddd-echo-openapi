package node

import (
	"cosslan/config"
	"cosslan/internal/domain/service"
)

type Service struct {
	svc service.NodeService
}

func NewService(cfg config.Config) *Service {
	return &Service{svc: service.NewNodeServiceImpl(cfg)}
}
