package handler

import (
	userv1 "cosslan/internal/api/http/v1/user"
	"cosslan/internal/app/dto"
	"cosslan/internal/app/user"
	pkghttp "cosslan/pkg/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

var _ userv1.ServerInterface = &UserHandler{}

type UserHandler struct {
	useCase user.UseCase
}

func NewUserHandler(useCase user.UseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (u *UserHandler) Login(ctx echo.Context) error {
	err := u.useCase.CreateUser(ctx.Request().Context(), dto.CreateUserDTO{})
	if err != nil {
		return pkghttp.Fail(ctx, http.StatusInternalServerError, "create user error")
	}

	return pkghttp.Success(ctx, "create user success", nil)
}
