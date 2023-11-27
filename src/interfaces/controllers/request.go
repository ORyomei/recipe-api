package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// BindUintPathParam bind uint path param
func BindUintPathParam(ctx *gin.Context, paramName string) (uint, *Error, error) {
	param, err := strconv.ParseUint(ctx.Param(paramName), 10, 64)
	if err != nil {
		e := NewBadRequestError("リクエストが間違っています")
		return 0, e, errors.Wrapf(err, "strconv.ParseUint %s failed", paramName)
	}
	return uint(param), nil, nil
}

// BindStringPathParam bind string path param
func BindStringPathParam(ctx *gin.Context, paramName string) string {
	return ctx.Param(paramName)
}

// BindJSON bind json body
func BindJSON(ctx *gin.Context, body interface{}) (*Error, error) {
	err := ctx.BindJSON(body)
	e := NewBadRequestError("パラメータが間違っています")
	return e, err
}
