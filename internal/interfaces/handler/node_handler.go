package handler

import (
	nodev1 "cosslan/internal/api/http/v1/node"
	"cosslan/internal/app/dto"
	"cosslan/internal/app/node"
	"github.com/labstack/echo/v4"
)

var _ nodev1.ServerInterface = &NodeHandler{}

type NodeHandler struct {
	useCase node.UseCase
}

func NewNodeHandler(useCase node.UseCase) *NodeHandler {
	return &NodeHandler{useCase: useCase}
}

func (n *NodeHandler) Node(ctx echo.Context) error {
	err := n.useCase.CreateNode(ctx.Request().Context(), dto.CreateNodeDTO{})
	if err != nil {
		return err
	}
	return nil
}
