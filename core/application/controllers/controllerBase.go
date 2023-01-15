package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IControllerBase interface {
	Ok(*gin.Context, interface{})
	Created(*gin.Context, interface{})
	Found(*gin.Context, interface{})
	BadRequest(*gin.Context)
	NotFound(*gin.Context)
	InternalServerError(*gin.Context)
}

type ControllerBase struct {
}

func (c *ControllerBase) Ok(ctx *gin.Context, object interface{}) {
	ctx.JSON(http.StatusOK, object)
}

func (c *ControllerBase) Created(ctx *gin.Context, object interface{}) {
	ctx.JSON(http.StatusCreated, object)
}

func (c *ControllerBase) Found(ctx *gin.Context, object interface{}) {
	ctx.JSON(http.StatusFound, object)
}

func (c *ControllerBase) BadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Bad Request",
	})
}

func (c *ControllerBase) NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}

func (c *ControllerBase) InternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}
