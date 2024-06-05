package line

import (
	"context"
	"cosslan/internal/app/line"
	"fmt"
)

var _ LineService = &LineServiceImpl{}

type LineService interface {
	CreateLine(ctx context.Context, dto line.CreateLineDTO) error
}

type LineServiceImpl struct {
	repo LineRepository
}

func NewLineService(repo LineRepository) *LineServiceImpl {
	return &LineServiceImpl{repo: repo}
}

func (l *LineServiceImpl) CreateLine(ctx context.Context, dto line.CreateLineDTO) error {
	fmt.Println("创建线路")
	return nil
}
