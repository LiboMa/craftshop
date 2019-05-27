package markets

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/LiboMa/otcmarket/common"
	"github.com/gin-gonic/gin"
)

func MarketsRegister(router *gin.RouterGroup) {
	// router.POST("/", ProductCreate)
	// router.PUT("/:id", ProductUpdate)
	// router.DELETE("/:id", ProductDelete)

}

func MarketsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", MarketList)
	router.GET("/usdtcny", MarketUsdt)
	// router.GET("/:id", ProductRetrieve)
	//router.GET("/:slug/comments", ProductCommentList)
}

func MarketList(c *gin.Context) {
	//condition := ArticleModel{}
	// name := c.Query("name")

	// get data from models
	// marketList, err := GetMarketList()
	//articleModels, modelCount, err := FindManyArticle(tag, author, limit, offset, favorited)

	// serialized to json
	//log, err := common.myLogger()

	var currencyMessage CurrencyMessage
	err := HttpGetDataBinding("https://api.huobi.com/v1/common/currencys", &currencyMessage)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("Invalid param")))
		return
	}
	// return http with json body
	//var users = json.RawMessage(`[{"username" : "akbar", "email": "akb@r.app"}, {"username" : "arkan", "email": "ark@n.app"}]`)

	// serializer := MarketsSerializer{c, productList}
	// c.JSON(http.StatusOK, gin.H{"markets": serializer.Response()})
	c.JSON(http.StatusOK, gin.H{"markets": currencyMessage})
}

func MarketUsdt(c *gin.Context) {
	//condition := ArticleModel{}
	tradetype := c.Query("tradeType")

	if tradetype == "" || (tradetype != "sell" && tradetype != "buy") {
		c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("using params, typeType=sell|buy")))
		return
	}

	var otcTradeMarket OTCTradeMarket

	// get data from cache
	key := fmt.Sprintf("market-price-%s", tradetype)
	//client := common.InitCache() //1. slowest method
	client := common.GetCache() //2. 3 * times increated
	//val, err := common.GetCacheItem(key) //3. almost the same to method 2
	val, err := client.Get(key).Result()
	if err != nil {
		log.Println(err)
	}
	// get data from db if failure
	json.Unmarshal([]byte(val), &otcTradeMarket)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("get data failed")))
		return
	}
	//c.JSON(http.StatusOK, gin.H{"market-price": (*otcTradeMarket.Data)[0].Price, "status": otcTradeMarket.Success})
	c.JSON(http.StatusOK, gin.H{key: otcTradeMarket.Data[0], "status": otcTradeMarket.Success})
}

// func MarketRetrieve(c *gin.Context) {
// 	//id := c.Param("id")
// 	id, err := strconv.Atoi(c.Param("id"))
// 	productmodel, err := GetProductByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, common.NewError("product", errors.New("Invalid id")))
// 		return
// 	}

// 	fmt.Println(productmodel)
// 	serializer := ProductSerializer{c, productmodel}
// 	c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
// }
