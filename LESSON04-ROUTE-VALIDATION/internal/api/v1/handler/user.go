package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mamba.com/route-group/utils"
)

type UserHandler struct {
}

type GetUsersByIdV1Param struct {
	ID int `uri:"id" binding:"gt=0"`
}

type GetUsersByUuidV1Param struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// User API

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all user (v1)"})
}

func (u *UserHandler) GetUsersByIdV1(ctx *gin.Context) {

	var params GetUsersByIdV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (v1)",
		"user_id": params.ID,
	})

	// idStr := ctx.Param("id")

	// id, err := utils.ValidationPositiveInt("ID", idStr)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Get user by ID (v1)",
	// 	"user_id": id,
	// })

}

func (u *UserHandler) GetUsersByUuidV1(ctx *gin.Context) {

	var params GetUsersByUuidV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by UUID (v1)",
		"user_id": params.Uuid,
	})

	// uuidStr := ctx.Param("uuid")

	// uid, err := utils.ValidationUuid("UUID", uuidStr)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message":   "Get user by UUID (v1)",
	// 	"user_uuid": uid,
	// })

}

func (u *UserHandler) PostUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create User (v1)"})
}

func (u *UserHandler) PutUsersByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update User By ID (v1)"})
}

func (u *UserHandler) DeleteUsersByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete User By ID (v1)"})
}
