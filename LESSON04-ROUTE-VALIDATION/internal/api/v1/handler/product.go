package v1handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"mamba.com/route-group/utils"
)

type ProductHandler struct {
}

var (
	slugRegex   = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// Product API

func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (v1)",
		"limit":   limit,
	})
}

func (p *ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if err := utils.ValidationRegex("Product", slug, slugRegex, "must contain only lowercase letters, numbers, hyphens and dots"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product By Slug (v1)",
		"slug":    slug,
	})
}

func (p *ProductHandler) PostProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create Product (v1)"})
}

func (p *ProductHandler) PutProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update Product By ID (v1)"})
}

func (p *ProductHandler) DeleteProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete Product By ID (v1)"})
}
