package routers

import (
	// "go-test-api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	// product := controllers.New()
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	})

	// router.POST("/products", product.CreateProduct)
	// router.GET("/products", product.GetProducts)
	// router.GET("/products/:id", product.GetProductById)

	return router
}
