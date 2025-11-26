package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// User API

func (u *UserHandler) GetUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all user (v1)"})
}

func (u *UserHandler) GetUsersByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get User By ID (v1)"})
}

func (u *UserHandler) PostUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create User (v1)"})
}

func (u *UserHandler) PutUsersByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update User By ID (v1)"})
}

func (u *UserHandler) DeleteUsersByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete User By ID (v1)"})
}
