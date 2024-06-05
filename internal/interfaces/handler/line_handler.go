package handler

import (
	linev1 "cosslan/internal/api/http/v1/line"
	"cosslan/internal/app/line"
	"github.com/labstack/echo/v4"
)

var _ linev1.ServerInterface = &LineHandler{}

type LineHandler struct {
	useCase line.UseCase
}

func NewLineHandler(useCase line.UseCase) *LineHandler {
	return &LineHandler{useCase: useCase}
}

func (l *LineHandler) Line(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
