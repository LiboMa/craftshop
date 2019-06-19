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
	router.GET("/cny", MarketCNY)
	// router.GET("/:id", ProductRetrieve)
	//router.GET("/:slug/comments", ProductCommentList)
}

func MarketList(c *gin.Context) {
	// name := c.Query("name")

	// get data from models
	// marketList, err := GetMarketList()
	//articleModels, modelCount, err := FindManyArticle(tag, author, limit, offset, favorited)

	// serialized to json
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
	c.Header("Host", "")

	// type Result struct {
	// 	Price  float64
	// 	Status bool
	// }

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

	serializer := MarketPriceSerializer{c, otcTradeMarket, tradetype}
	// result := Result{Price: otcTradeMarket.Data[0].Price, Status: otcTradeMarket.Success}

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("get data failed")))
		return
	}
	//c.JSON(http.StatusOK, gin.H{"market-price": (*otcTradeMarket.Data)[0].Price, "status": otcTradeMarket.Success})
	c.JSON(http.StatusOK, gin.H{key: serializer.Response()})
}

func MarketCNY(c *gin.Context) {
	//condition := ArticleModel{}
	tradetype := c.Query("tradeType")
	c.Header("Host", "")

	// type Result struct {
	// 	Price  float64
	// 	Status bool
	// }

	if tradetype == "" || (tradetype != "sell" && tradetype != "buy") {
		c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("using params, typeType=sell|buy")))
		return
	}

	// Code       int         `json: code`       // "code": 200,
	// Message    string      `json: message`    //"message": "成功",
	// totalCount int         `json: totalCount` // "totalCount": 183,
	// PageSize   int         `json: pageSize`   // "pageSize": 10,
	// TotalPage  int         `json: totalPage`  // "totalPage": 19,
	// CurrPage   int         `json: currPage`   // "currPage": 1,
	// Data       []*DataList `json: data`       //"data": xx
	// Success    bool        `json: success`    // "success": true

	var otcTradeMarket OTCTradeMarket

	// get data from cache
	key := fmt.Sprintf("market-price-%s", tradetype)
	val := json.RawMessage(`{"code":200,"message":"成功","totalCount":300,"pageSize":10,"totalPage":30,"currPage":1,"data":[{"id":354157,"uid":86613404,"userName":"潮人码头","merchantLevel":2,"coinId":2,"currency":1,"tradeType":1,"blockType":1,"payMethod":"1","payTerm":15,"payName":"[{\"bankName\":\"商家小号和搬砖的不交易，请取消否则收款卡退回\",\"bankType\":1,\"id\":2594223}]","minTradeLimit":50000.0000000000,"maxTradeLimit":1174800,"price":6.96,"tradeCount":168793.1839090000,"isOnline":true,"tradeMonthTimes":766,"orderCompleteRate":99,"takerLimit":0,"gmtSort":1560927014000}], "success":"true"}`)

	//client := common.GetCache() //2. 3 * times increated
	//val, err := common.GetCacheItem(key) //3. almost the same to method 2
	//val, err := client.Get(key).Result()
	//if err != nil {
	//	log.Println(err)
	//	}
	// get data from db if failure
	json.Unmarshal([]byte(val), &otcTradeMarket)

	otcTradeMarket.Data[0].Price = 1
	otcTradeMarket.Data[0].Currency = 2
	otcTradeMarket.Success = true

	serializer := MarketPriceSerializer{c, otcTradeMarket, tradetype}
	// result := Result{Price: otcTradeMarket.Data[0].Price, Status: otcTradeMarket.Success}

	c.JSON(http.StatusOK, gin.H{key: serializer.Response()})
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("get data failed")))
	// 	return
	// }
	//c.JSON(http.StatusOK, gin.H{"market-price": (*otcTradeMarket.Data)[0].Price, "status": otcTradeMarket.Success})
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
