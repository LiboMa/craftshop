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
	router.GET("/usdt_cny", MarketUsdt)
	router.GET("/cny_cny", MarketCNY)
	router.GET("/usdtcny", MarketUsdtv2)
	router.GET("/cnycny", MarketCNYv2)
	router.GET("/symbols", CryptoMarket)
	router.GET("/symbols/:symbol", SingleCryptoMarket)
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

func MarketUsdtv2(c *gin.Context) {
	//condition := ArticleModel{}
	tradetype := c.Query("tradeType")
	c.Header("Host", "")
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
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}
func MarketCNY(c *gin.Context) {
	tradetype := c.Query("tradeType")
	c.Header("Host", "")

	var otcTradeMarket OTCTradeMarket

	// get data from cache
	key := fmt.Sprintf("market-price-%s", tradetype)
	val := json.RawMessage(`{"code":200,"message":"成功","totalCount":300,"pageSize":10,"totalPage":30,"currPage":1,"data":[{"id":354157,"uid":86613404,"userName":"潮人码头","merchantLevel":2,"coinId":2,"currency":1,"tradeType":1,"blockType":1,"payMethod":"1","payTerm":15,"payName":"[{\"bankName\":\"商家小号和搬砖的不交易，请取消否则收款卡退回\",\"bankType\":1,\"id\":2594223}]","minTradeLimit":50000.0000000000,"maxTradeLimit":1174800,"price":6.96,"tradeCount":168793.1839090000,"isOnline":true,"tradeMonthTimes":766,"orderCompleteRate":99,"takerLimit":0,"gmtSort":1560927014000}], "success":"true"}`)

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

func MarketCNYv2(c *gin.Context) {
	tradetype := c.Query("tradeType")
	c.Header("Host", "")

	var otcTradeMarket OTCTradeMarket

	// get data from cache
	// key := fmt.Sprintf("market-price-%s", tradetype)
	val := json.RawMessage(`{"code":200,"message":"成功","totalCount":300,"pageSize":10,"totalPage":30,"currPage":1,"data":[{"id":354157,"uid":86613404,"userName":"潮人码头","merchantLevel":2,"coinId":2,"currency":1,"tradeType":1,"blockType":1,"payMethod":"1","payTerm":15,"payName":"[{\"bankName\":\"商家小号和搬砖的不交易，请取消否则收款卡退回\",\"bankType\":1,\"id\":2594223}]","minTradeLimit":50000.0000000000,"maxTradeLimit":1174800,"price":6.96,"tradeCount":168793.1839090000,"isOnline":true,"tradeMonthTimes":766,"orderCompleteRate":99,"takerLimit":0,"gmtSort":1560927014000}], "success":"true"}`)

	json.Unmarshal([]byte(val), &otcTradeMarket)

	otcTradeMarket.Data[0].Price = 1
	otcTradeMarket.Data[0].Currency = 2
	otcTradeMarket.Success = true

	serializer := MarketPriceSerializer{c, otcTradeMarket, tradetype}
	// result := Result{Price: otcTradeMarket.Data[0].Price, Status: otcTradeMarket.Success}

	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("get data failed")))
	// 	return
	// }
	//c.JSON(http.StatusOK, gin.H{"market-price": (*otcTradeMarket.Data)[0].Price, "status": otcTradeMarket.Success})
}

func CryptoMarket(c *gin.Context) {
	symbol := c.Query("symbol")
	c.Header("Host", "")

	var huobiMarket HuobiMarket
	var key string
	var val string

	// if symbol == "" || (symbol != "sell" && symbol != "buy") {
	// 	c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("using params, symbol=eth_usdt")))
	// 	return
	// }
	// get data from cache
	if symbol != "" {
		key = fmt.Sprintf("market-huobi-%s", symbol)
	} else {
		key = fmt.Sprintf("market-huobi")
	}

	client := common.GetCache() //2. 3 * times increated
	val, err := client.Get(key).Result()
	if err != nil {
		log.Println(err)
	}
	// get data from db if failure
	json.Unmarshal([]byte(val), &huobiMarket)
	serializer := HuobiMarketSerializer{c, &huobiMarket}

	// empty bug here:?
	c.JSON(http.StatusOK, serializer.Response())

}

func SingleCryptoMarket(c *gin.Context) {
	symbol := c.Param("symbol")
	c.Header("Host", "")

	var huobiMarketData MarketData
	var key string
	var val string

	// if symbol == "" || (symbol != "sell" && symbol != "buy") {
	// 	c.JSON(http.StatusNotFound, common.NewError("markets", errors.New("using params, symbol=eth_usdt")))
	// 	return
	// }
	// get data from cache
	key = fmt.Sprintf("market-huobi-%s", symbol)

	client := common.GetCache() //2. 3 * times increated
	val, err := client.Get(key).Result()
	if err != nil {
		log.Println(err)
	}
	// get data from db if failure
	json.Unmarshal([]byte(val), &huobiMarketData)
	//serializer := HuobiMarketSerializer{c, &huobiMarket}
	response := marketDataRes{}
	response.Price = huobiMarketData.Close
	response.Symbol = huobiMarketData.Symbol

	// empty bug here:?
	c.JSON(http.StatusOK, response)

}
