package main

import (
	"github.com/LiboMa/craftshop/common"
	"github.com/LiboMa/craftshop/products"
	"github.com/gin-gonic/gin"
)

func main() {

	// init db connection
	db := common.InitDB()
	defer db.Close()
	// init redis connection

	// start gin
	r := gin.Default()
	// init routers
	//
	v1 := r.Group("/api")

	products.ProductsRegister(v1.Group("/products"))
	products.ProductsAnonymousRegister(v1.Group("/products"))

	testMock := r.Group("/api/ping")

	testMock.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// db operationss
	r.Run(":8080")
}
