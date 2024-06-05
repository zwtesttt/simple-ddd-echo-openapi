package handler

import (
	lanv1 "cosslan/internal/api/http/v1/line"
	nodev1 "cosslan/internal/api/http/v1/node"
	userv1 "cosslan/internal/api/http/v1/user"
	"github.com/labstack/echo/v4"
)

var _ userv1.ServerInterface = &HttpServer{}
var _ lanv1.ServerInterface = &HttpServer{}
var _ nodev1.ServerInterface = &HttpServer{}

type HttpServer struct {
}

func (h HttpServer) Login(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) Lan(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) Node(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
