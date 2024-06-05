package line

import (
	"cosslan/internal/domain/line"
)

type LineService struct {
	repo line.LineService
}

func NewLineService(repo line.LineService) *LineService {
	return &LineService{repo: repo}
}

//// CreateLan handles the logic for creating a LAN.
//func (s *LineService) CreateLine(ctx context.Context, dto CreateLanDTO) (line.Lan, error) {
//	// Convert DTO to domain entity
//	lan := line.NewLan(dto.Name, dto.CIDR)
//	if err := s.repo.Save(ctx, lan); err != nil {
//		return lan, err
//	}
//	return lan, nil
//}
