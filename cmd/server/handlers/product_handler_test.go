package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Andrea-Reyna/go-web/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// var productsList = []domain.Product{
// 	{
// 		ID:          1,
// 		Name:        "Oil - Margarine",
// 		Quantity:    439,
// 		CodeValue:   "S82254D",
// 		IsPublished: true,
// 		Expiration:  "15/12/2021",
// 		Price:       71.42,
// 	},
// 	{
// 		ID:          2,
// 		Name:        "Pineapple - Canned, Rings MOD",
// 		Quantity:    345,
// 		CodeValue:   "M4637",
// 		IsPublished: true,
// 		Expiration:  "09/08/2021",
// 		Price:       352.79,
// 	},
// 	{
// 		ID:          3,
// 		Name:        "Wine - Red Oakridge Merlot",
// 		Quantity:    367,
// 		CodeValue:   "T65812",
// 		IsPublished: false,
// 		Expiration:  "24/05/2021",
// 		Price:       179.23,
// 	},
// }

var newProduct = CreateProductRequest{
	Name:        "New product by test",
	Quantity:    130,
	CodeValue:   "M71599",
	IsPublished: false,
	Expiration:  "28/01/2022",
	Price:       275.47,
}
var sameCode = CreateProductRequest{
	Name:        "New product by test",
	Quantity:    130,
	CodeValue:   "M4315",
	IsPublished: false,
	Expiration:  "28/01/2022",
	Price:       275.47,
}

var updateProduct = CreateProductRequest{
	Name:        "Update by test",
	Quantity:    138,
	CodeValue:   "S82254D",
	IsPublished: false,
	Expiration:  "01/01/2022",
	Price:       555.47,
}

func createServerForTestPrductsHandler() *gin.Engine {
	os.Setenv("FILE", "/Users/areyna/Documents/Bootcamp/Modulo 4 - Go Web/Activities/products.json")

	repository, err := products.NewSliceBasedRepository()
	if err != nil {
		panic(err)
	}

	service := &products.DefaultService{
		Storage: repository,
	}

	handler := &ProductHandlers{
		Service: service,
	}

	gin.SetMode(gin.TestMode)

	server := gin.New()

	group := server.Group("products")
	group.POST("", handler.Create())
	group.GET("", handler.GetAll())
	group.GET("/:id", handler.FindById())
	group.PUT("/:id", handler.Update())
	group.DELETE("/:id", handler.Delete())
	return server
}

func TestProductsHandler_GetAll(t *testing.T) {
	t.Run("should return a product list", func(t *testing.T) {
		var (
			expectedStatusCode = http.StatusOK
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
		)
		request := httptest.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.True(t, len(response.Body.String()) > 0)
		//assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}

func TestProductsHandler_GetProductByID(t *testing.T) {
	t.Run("should return a product", func(t *testing.T) {
		var (
			productIdToSearch  = "4"
			expectedStatusCode = http.StatusOK
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}

			expectedResponse = `{
					"id": 4,
					"name": "Cookie - Oatmeal",
					"quantity": 130,
					"code_value": "M7157",
					"is_published": false,
					"expiration": "28/01/2022",
					"price": 275.47
			}`
		)

		request := httptest.NewRequest(http.MethodGet, "/products/"+productIdToSearch, nil)
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return an error if the product is not found", func(t *testing.T) {
		var (
			productIdToSearch  = "777"
			expectedStatusCode = http.StatusNotFound
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
			expectedResponse = `{"error":"product not found"}`
		)

		request := httptest.NewRequest(http.MethodGet, "/products/"+productIdToSearch, nil)
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return an format error", func(t *testing.T) {
		var (
			productIdToSearch  = "qwer"
			expectedStatusCode = http.StatusBadRequest
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
			expectedResponse = `{"error":"invalid data"}`
		)

		request := httptest.NewRequest(http.MethodGet, "/products/"+productIdToSearch, nil)
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}

func TestProductsHandler_Create(t *testing.T) {
	t.Run("should create product", func(t *testing.T) {
		newProductJson, err := json.Marshal(newProduct)
		if err != nil {
			t.Log(err)
			return
		}
		var (
			expectedStatusCode = http.StatusCreated
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
			expectedResponse = `{
				"id":501,
				"name": "New product by test",
				"quantity": 130,
				"code_value": "M71599",
				"is_published": false,
				"expiration": "28/01/2022",
				"price": 275.47
			}`
		)
		request := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte(string(newProductJson))))
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return an error if product code exist", func(t *testing.T) {
		newProductJson, err := json.Marshal(sameCode)
		if err != nil {
			t.Log(err)
			return
		}
		var (
			expectedStatusCode = http.StatusConflict
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
			expectedResponse = `{"code":"Conflict", "message":"canot create the given product it already exist", "status":409}`
		)
		request := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte(string(newProductJson))))
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}

func TestProductsHandler_Update(t *testing.T) {
	t.Run("should update product", func(t *testing.T) {
		newProductJson, err := json.Marshal(updateProduct)
		if err != nil {
			t.Log(err)
			return
		}
		var (
			productIdToUpdate  = "1"
			expectedStatusCode = http.StatusOK
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
			expectedResponse = `{
				"id":1,
				"name": "Update by test",
				"quantity": 138,
				"code_value": "S82254D",
				"is_published": false,
				"expiration": "01/01/2022",
				"price": 555.47
			}`
		)
		request := httptest.NewRequest(http.MethodPut, "/products/"+productIdToUpdate, bytes.NewBuffer([]byte(string(newProductJson))))
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return an error if product code exist", func(t *testing.T) {
		newProductJson, err := json.Marshal(sameCode)
		if err != nil {
			t.Log(err)
			return
		}
		var (
			productIdToUpdate  = "1"
			expectedStatusCode = http.StatusConflict
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
			expectedResponse = `{"code":"Conflict", "message":"code already exist", "status":409}`
		)
		request := httptest.NewRequest(http.MethodPut, "/products/"+productIdToUpdate, bytes.NewBuffer([]byte(string(newProductJson))))
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}

func TestProductsHandler_Delete(t *testing.T) {
	t.Run("should delete product", func(t *testing.T) {
		var (
			productIdToDelete  = "1"
			expectedStatusCode = http.StatusNoContent
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
		)
		request := httptest.NewRequest(http.MethodDelete, "/products/"+productIdToDelete, nil)
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("should return an error if product not exist", func(t *testing.T) {
		var (
			productIdToDelete  = "1"
			expectedStatusCode = http.StatusNotFound
			expectedHeaders    = http.Header{
				"Content-Type": []string{
					"application/json; charset=utf-8",
				},
			}
			expectedResponse = `{"code":"NotFound", "message":"product not found", "status":404}`
		)
		request := httptest.NewRequest(http.MethodDelete, "/products/"+productIdToDelete, nil)
		response := httptest.NewRecorder()

		server := createServerForTestPrductsHandler()

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}
