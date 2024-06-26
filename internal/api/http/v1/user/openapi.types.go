// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package v1

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Response defines model for Response.
type Response struct {
	Code int                     `json:"code"`
	Data *map[string]interface{} `json:"data,omitempty"`
	Msg  string                  `json:"msg"`
}

// UserLoginRequest defines model for UserLoginRequest.
type UserLoginRequest struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = UserLoginRequest
