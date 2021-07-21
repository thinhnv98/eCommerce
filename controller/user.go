package controller

import (
	"eCommerce/models"
	"eCommerce/service"
	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	IUserService service.IUserService
}

func (_self UserController) Create(ctx *gin.Context) {
	var u models.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := _self.IUserService.Create(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	return
}

func (_self UserController) Login(ctx *gin.Context) {
	var u models.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isSuccess, id, err := _self.IUserService.Login(ctx, u.UserName, u.PasswordHash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if !isSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errors.New("username or password invalid").Error(),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": true,
		"id":      id,
	})
	return
}
