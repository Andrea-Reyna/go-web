package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Andrea-Reyna/go-web/internal/products"
	"github.com/Andrea-Reyna/go-web/pkg/rest"
	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	Service products.Service
}

// @Summary Create a new product
// @Description This method creates a new product entry in the system by taking a JSON input with the required product information. It returns an error if there is an issue with the input data, if the product already exists, or if there is an internal server error.
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Product Information"
// @Success 201 {object} CreateProductResponse "Successfully created product"
// @Failure 400 {object} rest.ErrorResponse "BadRequest: invalid data"
// @Failure 409 {object} rest.ErrorResponse "Conflict: cannot create the given product, it already exists or error date format"
// @Failure 500 {object} rest.ErrorResponse "InternalServerError: an internal error has occurred"
// @Router /products [post]
func (handler ProductHandlers) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CreateProductRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Status:  400,
				Code:    "BadRequest",
				Message: "invalid data",
			})
			return
		}
		productToCreate := request.ToDomain()
		err := handler.Service.Create(&productToCreate)
		if err != nil {
			switch err {
			case products.ErrProductAlreadyExists:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "canot create the given product it already exist",
				})
			case products.ErrFormateDate:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "error date format",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
					Status:  500,
					Code:    "InternalServerError",
					Message: "an internal error has ocurred",
				})
			}
			return
		}
		ctx.JSON(http.StatusCreated, productToCreate)
	}
}

// @Summary Update a product
// @Description This method update a product entry in the system by taking a JSON input with the required product information and Id. It returns an error if there is an issue with the input data, if the product code already exists, or if there is an internal server error.
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Product Information"
// @Param id path int true "Product ID"
// @Success 201 {object} CreateProductResponse "Successfully updated product"
// @Failure 400 {object} rest.ErrorResponse "BadRequest: invalid data"
// @Failure 409 {object} rest.ErrorResponse "Conflict: cannot update product, code already exists or error date format"
// @Failure 500 {object} rest.ErrorResponse "InternalServerError: an internal error has occurred"
// @Router /products/:id [put]
func (handler ProductHandlers) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}
		var request CreateProductRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Status:  400,
				Code:    "BadRequest",
				Message: "bad request",
			})
			return
		}
		productToCreate := request.ToDomain()
		productToCreate.ID = id

		err = handler.Service.Update(&productToCreate)
		if err != nil {
			switch err {
			case products.ErrProductAlreadyExists:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "code already exist",
				})
			case products.ErrFormateDate:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "error date format",
				})
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, rest.ErrorResponse{
					Status:  404,
					Code:    "NotFound",
					Message: "product not found",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
					Status:  500,
					Code:    "InternalServerError",
					Message: "an internal error has ocurred",
				})
			}
			return
		}
		ctx.JSON(http.StatusOK, productToCreate)
	}
}

// @Summary Update partial a product
// @Description This method update a product entry in the system by taking a JSON input with the required product information and Id. It returns an error if there is an issue with the input data, if the product code already exists, or if there is an internal server error.
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Product Information"
// @Param id path int true "Product ID"
// @Success 201 {object} CreateProductResponse "Successfully updated product"
// @Failure 400 {object} rest.ErrorResponse "BadRequest: invalid data"
// @Failure 409 {object} rest.ErrorResponse "Conflict: cannot update product, code already exists or error date format"
// @Failure 500 {object} rest.ErrorResponse "InternalServerError: an internal error has occurred"
// @Router /products/:id [patch]
func (handler ProductHandlers) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Status:  400,
				Code:    "BadRequest",
				Message: "invalid data",
			})
			return
		}

		prod, err := handler.Service.FindById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, rest.ErrorResponse{
				Status:  404,
				Code:    "NotFound",
				Message: "product not found",
			})
		}

		if err := json.NewDecoder(ctx.Request.Body).Decode(&prod); err != nil {
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Status:  400,
				Code:    "BadRequest",
				Message: "invalid data",
			})
			return
		}
		prod.ID = id


		err = handler.Service.Update(&prod)
		if err != nil {
			switch err {
			case products.ErrProductAlreadyExists:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "code already exist",
				})
			case products.ErrFormateDate:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "error date format",
				})
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, rest.ErrorResponse{
					Status:  404,
					Code:    "NotFound",
					Message: "product not found",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
					Status:  500,
					Code:    "InternalServerError",
					Message: "an internal error has ocurred",
				})
			}
			return
		}
		ctx.JSON(http.StatusOK, prod)
	}

}

func (handler ProductHandlers) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Status:  400,
				Code:    "BadRequest",
				Message: "invalid data",
			})
			return
		}
		name := ctx.Query("name")

		product, err := handler.Service.UpdateName(id, name)
		if err != nil {
			switch err {
			case products.ErrProductAlreadyExists:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "code already exist",
				})
			case products.ErrFormateDate:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Status:  409,
					Code:    "Conflict",
					Message: "error date format",
				})
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, rest.ErrorResponse{
					Status:  404,
					Code:    "NotFound",
					Message: "product not found",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
					Status:  500,
					Code:    "InternalServerError",
					Message: "an internal error has ocurred",
				})
			}
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

// @Summary Get All products
// @Description This method get a list with all products.
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} domain.Product "Successfully list of products"
// @Router /products [get]
func (handler ProductHandlers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := handler.Service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
				Status:  500,
				Code:    "InternalServerError",
				Message: "an internal error has ocurred",
			})
		}
		ctx.JSON(http.StatusOK, products)
	}
}

// @Summary Get product by ID
// @Description Retrieves a specific product by its ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} domain.Product "Successfully retrieved product"
// @Failure 400 {object} map[string]string "Invalid data"
// @Failure 404 {object} map[string]string "Product not found"
// @Router /products/{id} [get]
func (handler ProductHandlers) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}
		product, err := handler.Service.FindById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

// @Summary Search products by price
// @Description Retrieves a list of products with a price greater than the specified value
// @Tags products
// @Accept  json
// @Produce  json
// @Param   priceGt     query    float64     true    "Minimum product price"
// @Success 200 {array} domain.Product "Successfully retrieved list of products"
// @Failure 400 {object} rest.ErrorResponse "Invalid data"
// @Failure 404 {object} map[string]string "Price must be greater than 0"
// @Router /products/search [get]
func (handler ProductHandlers) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		priceGt, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Status:  400,
				Code:    "BadRequest",
				Message: "invalid data",
			})
			return
		}
		filterProducts, err := handler.Service.Search(priceGt)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "el precio ingresado debe ser mayor a 0"})
			return
		}
		ctx.JSON(http.StatusOK, filterProducts)
	}
}

// @Summary Delete product by ID
// @Description Deletes a specific product by its ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 204 "Successfully deleted product"
// @Failure 400 {object} rest.ErrorResponse "Invalid data"
// @Failure 404 {object} rest.ErrorResponse "Product not found"
// @Failure 500 {object} rest.ErrorResponse "An internal error has occurred"
// @Router /products/{id} [delete]
func (handler ProductHandlers) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Status:  400,
				Code:    "BadRequest",
				Message: "invalid data",
			})
			return
		}

		err = handler.Service.Delete(id)
		if err != nil {
			switch err {
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, rest.ErrorResponse{
					Status:  404,
					Code:    "NotFound",
					Message: "product not found",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
					Status:  500,
					Code:    "InternalServerError",
					Message: "an internal error has ocurred",
				})
			}
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

// @Summary Get consumer prices for a list of product IDs
// @Description Retrieves a list of products with prices for the specified product IDs
// @Tags products
// @Accept  json
// @Produce  json
// @Param   list     query    string     true    "Comma-separated list of product IDs"
// @Success 200 {array} domain.Product "Successfully retrieved list of consumer products"
// @Failure 400 {object} rest.ErrorResponse "Invalid data"
// @Failure 500 {object} rest.ErrorResponse "An internal error has occurred"
// @Router /products/consumer-price [get]
func (handler ProductHandlers) ConsumerPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := ctx.Query("list")
		var nums []int
		for _, numStr := range strings.Split(params, ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
					Status:  400,
					Code:    "BadRequest",
					Message: "invalid data",
				})
				return
			}
			nums = append(nums, num)
		}

		consumerProducts, err := handler.Service.ConsumerPrice(nums)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
				Status:  500,
				Code:    "InternalServerError",
				Message: "an internal error has ocurred",
			})
			return
		}
		ctx.JSON(http.StatusOK, consumerProducts)
	}
}
