package line

import (
	"context"
	"cosslan/config"
	"cosslan/internal/app/dto"
	"fmt"
	"go.uber.org/zap"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	CreateLine(ctx context.Context, dto dto.CreateLineDTO) error
}

// LineUseCase implements the UseCase interface for Line.
type LineUseCase struct {
	service *Service
	logger  *zap.Logger
}

// NewLineUseCase creates a new LineUseCase.
func NewLineUseCase(cfg config.Config) *LineUseCase {
	return &LineUseCase{service: NewLineService(cfg)}
}

// CreateLine handles the use case for creating a Line.
func (uc *LineUseCase) CreateLine(ctx context.Context, dto dto.CreateLineDTO) error {
	fmt.Println("创建线路", dto)
	return nil
}
