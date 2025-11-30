package v1handler

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"mamba.com/route-group/utils"
)

type ProductHandler struct {
}

type GetProductsBySlugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=3,max=5"`
}

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=3,max=50,search"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}

type ProductImage struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required,file_ext=jpg png gif"`
}

type PostProductsV1Param struct {
	Name         string       `json:"name" binding:"required,min=3,max=100"`
	Price        int          `json:"price" binding:"required,min_int=100000"`
	Display      *bool        `json:"display" binding:"omitempty"`
	ProductImage ProductImage `json:"product_image" binding:"required"`
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
	var params GetProductsV1Param
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	if params.Limit == 0 {
		params.Limit = 1
	}

	if params.Email == "" {
		params.Email = "No Email"
	}

	if params.Date == "" {
		params.Date = time.Now().Format("2006-01-02")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product (v1)",
		"search":  params.Search,
		"limit":   params.Limit,
		"email":   params.Email,
		"date":    params.Date,
	})
	// limit := ctx.DefaultQuery("limit", "10")
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "List all products (v1)",
	// 	"limit":   limit,
	// })
}

func (p *ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {

	var params GetProductsBySlugV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product By Slug (v1)",
		"slug":    params.Slug,
	})

	// slug := ctx.Param("slug")

	// if err := utils.ValidationRegex("Product", slug, slugRegex, "must contain only lowercase letters, numbers, hyphens and dots"); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Get Product By Slug (v1)",
	// 	"slug":    slug,
	// })
}

func (p *ProductHandler) PostProductsV1(ctx *gin.Context) {

	var params PostProductsV1Param
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	if params.Display == nil {
		defaultDisplay := true
		params.Display = &defaultDisplay
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":       "Create Product (v1)",
		"name":          params.Name,
		"price":         params.Price,
		"display":       params.Display,
		"product_image": params.ProductImage,
	})

	// body, err := ctx.GetRawData()
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, "Error read body request")
	// }

	// ctx.JSON(http.StatusCreated, gin.H{
	// 	"message": "Create Product (v1)",
	// 	"data":    string(body),
	// })
}

func (p *ProductHandler) PutProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update Product By ID (v1)"})
}

func (p *ProductHandler) DeleteProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete Product By ID (v1)"})
}
