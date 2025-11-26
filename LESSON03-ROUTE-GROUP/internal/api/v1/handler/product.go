package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

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
	search := ctx.Query("search")

	if err := utils.ValidationRequired("Search", search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidationStringLength("Search", search, 3, 50); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidationRegex("Search", search, searchRegex, "must contain only letters, numbers and spaces"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be a positive number"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all Product (v1)",
		"search":  search,
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
