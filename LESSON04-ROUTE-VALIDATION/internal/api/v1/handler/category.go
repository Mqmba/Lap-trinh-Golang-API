package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mamba.com/route-group/utils"
)

type CategoryHandler struct {
}

type GetCategoryByCategoryV1Param struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}

type PostCategoriesV1Param struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

var validCategory = map[string]bool{
	"php":    true,
	"python": true,
	"golang": true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	var params GetCategoryByCategoryV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": params.Category,
	})
	// category := ctx.Param("category")

	// if err := utils.ValidationInList("Category", category, validCategory); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message":  "Category found",
	// 	"category": category,
	// })
}

func (c *CategoryHandler) PostCategoriesV1(ctx *gin.Context) {
	var param PostCategoriesV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post category (V1)",
		"name":    param.Name,
		"status":  param.Status,
	})
}
