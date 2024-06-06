package handler

import (
	linev1 "cosslan/internal/api/http/v1/line"
	"cosslan/internal/app/dto"
	"cosslan/internal/app/line"
	pkghttp "cosslan/pkg/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

var _ linev1.ServerInterface = &LineHandler{}

type LineHandler struct {
	useCase line.UseCase
}

func NewLineHandler(useCase line.UseCase) *LineHandler {
	return &LineHandler{useCase: useCase}
}

func (l *LineHandler) Line(ctx echo.Context) error {
	err := l.useCase.CreateLine(ctx.Request().Context(), dto.CreateLineDTO{
		Name: ctx.QueryParam("name"),
		CIDR: ctx.QueryParam("cidr"),
	})
	if err != nil {
		return pkghttp.Fail(ctx, http.StatusInternalServerError, "create user error")

	}
	return pkghttp.Success(ctx, "create line success", nil)
}
