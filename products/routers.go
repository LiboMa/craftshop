package products

import (
	"errors"
	"net/http"

	"github.com/LiboMa/craftshop/common"
	"github.com/gin-gonic/gin"
)

// "errors"
// "net/http"
// "strconv"

func ProductsRegister(router *gin.RouterGroup) {
	//router.POST("/", ProductCreate)
	//router.PUT("/:slug", ProductUpdate)
	//router.DELETE("/:slug", ProductDelete)
	//router.POST("/:slug/favorite", ProductFavorite)
	//router.DELETE("/:slug/favorite", ProductUnfavorite)
	//router.POST("/:slug/comments", ProductCommentCreate)
	//router.DELETE("/:slug/comments/:id", ProductCommentDelete)
}

func ProductsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", ProductList)
	//router.GET("/:slug", ProductRetrieve)
	//router.GET("/:slug/comments", ProductCommentList)
}

func ProductList(c *gin.Context) {
	//condition := ArticleModel{}
	// name := c.Query("name")
	// model := c.Query("model")
	// description := c.Query("description")
	// price := c.Query("price")
	// image := c.Query("image_url")
	// video := c.Query("video_url")
	// capacity := c.Query("capacity")

	// get data from models
	productList, err := GetProductList()
	//articleModels, modelCount, err := FindManyArticle(tag, author, limit, offset, favorited)

	// serialized to json

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
		return
	}
	// return http with json body
	//var users = json.RawMessage(`[{"username" : "akbar", "email": "akb@r.app"}, {"username" : "arkan", "email": "ark@n.app"}]`)

	serializer := ProductsSerializer{c, productList}
	c.JSON(http.StatusOK, gin.H{"products": serializer.Response()})
}
