package service

import (
	"context"
	"cosslan/config"
	"cosslan/internal/app/dto"
	"cosslan/internal/infra/persistence"
	"fmt"
)

var _ LineService = &LineServiceImpl{}

type LineService interface {
	CreateLine(ctx context.Context, dto dto.CreateLineDTO) error
}

type LineServiceImpl struct {
	repo *persistence.Repositories
}

func NewLineService(cfg config.Config) *LineServiceImpl {
	return &LineServiceImpl{repo: persistence.GetRepositories(cfg)}
}

func (l *LineServiceImpl) CreateLine(ctx context.Context, dto dto.CreateLineDTO) error {
	fmt.Println("创建线路")
	return nil
}
