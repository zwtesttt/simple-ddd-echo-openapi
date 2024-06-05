package line

import (
	"context"
	"cosslan/internal/domain/line"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	CreateLine(ctx context.Context, dto CreateLineDTO) (line.Line, error)
}

// LineUseCase implements the UseCase interface for Line.
type LineUseCase struct {
	service *LineService
}

// NewLineUseCase creates a new LineUseCase.
func NewLineUseCase(service *LineService) *LineUseCase {
	return &LineUseCase{service: service}
}

//// CreateLine handles the use case for creating a Line.
//func (uc *LineUseCase) CreateLine(ctx context.Context, dto CreateLineDTO) (line.Line, error) {
//	return uc.service.CreateLine(ctx, dto)
//}
