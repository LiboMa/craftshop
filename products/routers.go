package products

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LiboMa/craftshop/common"
	"github.com/gin-gonic/gin"
)

// "errors"
// "net/http"
// "strconv"

func ProductsRegister(router *gin.RouterGroup) {
	router.POST("/", ProductCreate)
	router.PUT("/:id", ProductUpdate)
	//router.DELETE("/:slug", ProductDelete)
	//router.POST("/:slug/favorite", ProductFavorite)
	//router.DELETE("/:slug/favorite", ProductUnfavorite)
	//router.POST("/:slug/comments", ProductCommentCreate)
	//router.DELETE("/:slug/comments/:id", ProductCommentDelete)
}

func ProductsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", ProductList)
	router.GET("/:id", ProductRetrieve)
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

func ProductRetrieve(c *gin.Context) {
	//id := c.Param("id")
	id, err := strconv.Atoi(c.Param("id"))
	productmodel, err := GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("product", errors.New("Invalid id")))
		return
	}

	fmt.Println(productmodel)
	serializer := ProductSerializer{c, productmodel}
	c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
}

func ProductCreate(c *gin.Context) {

	// data is the test data supposed received from request body
	// data := `{"name":"english_A3",
	// "model":"A3",
	// "price":999,
	// "description":"english lessons for children age of 6",
	// "image_url":"http://s3.edushop.com/static/images/en_a3.jepg",
	// "video_url":"http://s3.edushop.com/static/images/en_a3.jepg",
	// "Capacity":99
	// }`

	var product Products

	// validator should be added here..
	c.ShouldBind(&product)

	//log.Println(reflect.TypeOf(product))
	product.Created_on = common.MakeTimeStamp()
	product.Created_by = "admin"
	product.Modified_on = common.MakeTimeStamp()
	product.Modified_by = "admin"
	product.Labels = "test_lebels"
	product.State = 1

	fmt.Println("fe convert product: ", product)
	CreateProduct(&product)
	c.JSON(http.StatusCreated, gin.H{"product": product})

	//=======
	// err := json.Unmarshal([]byte(data), &product)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// product.Created_on = common.MakeTimeStamp()
	// product.Created_by = "admin"
	// product.Modified_on = common.MakeTimeStamp()
	// product.Modified_by = "admin"
	// product.Labels = "test_lebels"
	// product.State = 1

	// serializer := ProductSerializer{c, product}
	// c.JSON(http.StatusCreated, gin.H{"product": serializer.Response(), "result": "OK"})

}

func ProductUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	var product Products
	//var product_body Products
	c.ShouldBind(&product)

	productmodel, err := GetProductByID(id)

	log.Println(productmodel)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"product": product, "status": "updated"})

}
