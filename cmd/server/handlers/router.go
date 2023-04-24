package handlers

import (
	"github.com/Andrea-Reyna/go-web/cmd/server/middlewares"
	"github.com/Andrea-Reyna/go-web/internal/products"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func (router *Router) Setup() {
	router.Engine.Use(gin.Recovery())
	router.Engine.Use(gin.Logger())
	//router.Engine.Use(middlewares.Logger)

	router.SetProductsRoutes()
	router.Engine.Run()
}

func (router *Router) SetProductsRoutes() {

	repository, err := products.NewSliceBasedRepository()
	if err != nil {
		panic("error loading repository")
	}

	service := products.DefaultService{
		Storage: repository,
	}

	handler := ProductHandlers{
		Service: service,
	}

	group := router.Engine.Group("products")

	group.POST("", middlewares.ValidateToken, handler.Create())
	group.GET("", handler.GetAll())
	group.GET("/:id", handler.FindById())
	group.GET("/search", handler.Search())
	group.PUT("/:id", middlewares.ValidateToken, handler.Update())
	group.PATCH("/:id", middlewares.ValidateToken, handler.UpdatePartial())
	group.DELETE("/:id", middlewares.ValidateToken, handler.Delete())
	group.GET("/consumer_price", handler.ConsumerPrice())
}
