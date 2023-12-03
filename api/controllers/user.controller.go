package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	entytes "tot/api/elements"
	"tot/core/helpers"
)

type ControllerUser struct {
	user entytes.InfUserReg
}

func NewControllerUser(user entytes.InfUserReg) *ControllerUser {
	return &ControllerUser{user: user}
}

//func (s *ControllerUser) HandlerRegGet(ctx *gin.Context) {
//	var body entytess2.ShmValidUserReg
//	err := ctx.ShouldBind(&body)
//	if err != nil {
//		helpers.ValidateErrorResponse(ctx, err)
//		return
//	}
//	result, e := s.user.RegUser(body)
//	if e != nil {
//		helpers.ErrorResponse(ctx, err)
//		return
//	}
//	println(result)
//
//	helpers.APIResponse(ctx, "sssssss", http.StatusOK, nil)
//
//}

func (s *ControllerUser) HandlerRegPOST(ctx *gin.Context) {
	var body entytes.ShmValidUserReg
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		helpers.ValidateErrorResponse(ctx, err)
		return
	}
	result, err := s.user.RegUser(body)
	if err != nil {
		helpers.HandlerError(ctx, err)
		return
	}
	helpers.APIResponse(ctx, "ok", http.StatusOK, result)
}

func (s *ControllerUser) HandlerValidSmsPOST(ctx *gin.Context) {
	var body entytes.ShmValidSms
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		helpers.ValidateErrorResponse(ctx, err)
		return
	}
	result, err := s.user.ValidSmsUser(body)
	if err != nil {
		helpers.HandlerError(ctx, err)
		return
	}
	helpers.APIResponse(ctx, "ok", http.StatusOK, result)
}

func (s *ControllerUser) HandlerLoginPOST(ctx *gin.Context) {
	var body entytes.ShmValidUserReg
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		helpers.ValidateErrorResponse(ctx, err)
		return
	}
	result, err := s.user.LoginUser(body)
	if err != nil {
		helpers.HandlerError(ctx, err)
		return
	}
	helpers.APIResponse(ctx, "ok", http.StatusOK, result)
}

func (s *ControllerUser) HandlerRefreshTokenPOST(ctx *gin.Context) {
	var body entytes.ShmValidRefresh
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		helpers.ValidateErrorResponse(ctx, err)
		return
	}

	result, err := s.user.RefreshTokenUser(body)
	if err != nil {
		helpers.HandlerError(ctx, err)
		return
	}
	helpers.APIResponse(ctx, "ok", http.StatusOK, result)
}
