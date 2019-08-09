package markets

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CurrencyMessage struct {
	Status string   `json: status`
	Data   []string `json: data`
}

/* trade data format demo
{
    "code": 200,
    "message": "成功",
    "totalCount": 183,
    "pageSize": 10,
    "totalPage": 19,
    "currPage": 1,
    "data": [
        {
            "id": 283459,
            "uid": 21198620,
            "userName": "三万里",
            "merchantLevel": 2,
            "coinId": 2,
            "currency": 1,
            "tradeType": 1,
            "blockType": 1,
            "payMethod": "1",
            "payTerm": 15,
            "payName": "[{\"bankName\":\"工商银行\",\"bankType\":1,\"id\":2376119}]",
            "minTradeLimit": 50000.0000000000,
            "maxTradeLimit": 651704,
            "price": 7.06,
            "tradeCount": 92309.4475980000,
            "isOnline": true,
            "tradeMonthTimes": 764,
            "orderCompleteRate": 100,
            "takerLimit": 0,
            "gmtSort": 1558615971000
        },
        {
            "id": 351182,
            "uid": 1759834,
            "userName": "三万里高空",
            "merchantLevel": 3,
            "coinId": 2,
            "currency": 1,
            "tradeType": 1,
            "blockType": 1,
            "payMethod": "1",
            "payTerm": 15,
            "payName": "[{\"bankName\":\"工商银行\",\"bankType\":1,\"id\":2376336}]",
            "minTradeLimit": 50000.0000000000,
            "maxTradeLimit": 230800,
            "price": 7.06,
            "tradeCount": 32691.2181330000,
            "isOnline": true,
            "tradeMonthTimes": 2162,
            "orderCompleteRate": 99,
            "takerLimit": 0,
            "gmtSort": 1558616062000
        }

    ],
    "success": true
}
*/
type OTCTradeMarket struct {
	Code       int         `json: code`       // "code": 200,
	Message    string      `json: message`    //"message": "成功",
	totalCount int         `json: totalCount` // "totalCount": 183,
	PageSize   int         `json: pageSize`   // "pageSize": 10,
	TotalPage  int         `json: totalPage`  // "totalPage": 19,
	CurrPage   int         `json: currPage`   // "currPage": 1,
	Data       []*DataList `json: data`       //"data": xx
	Success    bool        `json: success`    // "success": true

}

type HuobiMarket struct {
	Status string        `json: status`
	Ts     int64         `json: ts`
	Data   []*MarketData `json: data`
}

type MarketData struct {
	Open   float64 `json: open`
	Close  float64 `json: close`
	Low    float64 `json: low`
	High   float64 `json: high`
	Amount float64 `json: amount`
	Count  float64 `json: count`
	Volume float64 `json: vol`
	Symbol float64 `json: Symbol`
}

type DataList struct {
	ID       int    `json: id`       //	"id": 351182,
	UID      int    `json: uid`      //	"uid": 1759834,
	UserName string `json: userName` //	"userName": "三万里高空",
	// "merchantLevel": 3,
	CoinID   int `json: coinId,`   //	"coinId": 2(USDT)
	Currency int `json: currency,` // "currency": 1,
	// "tradeType": 1,
	// "blockType": 1,
	// "payMethod": "1",
	// "payTerm": 15,
	// PayName       []struct{} `json: payName, omitempty`       // "payName": "[{\"bankName\":\"工商银行\",\"bankType\":1,\"id\":2376336}]",
	MinTradeLimit float64 `json: minTradeLimit` //"minTradeLimit": 50000.0000000000,
	MaxTradeLimit float64 `json: maxTradeLimit` //"maxTradeLimit": 230800,
	Price         float64 `json: price`         //"price": 7.06,
	TradeCount    float64 `json: tradeCount `   // "tradeCount": 32691.2181330000,
	// "isOnline": true,
	// "tradeMonthTimes": 2162,
	// "orderCompleteRate": 99,
	// "takerLimit": 0,
	// "gmtSort": 155861606200
}

// func GetHttpRequestBody(url string, datachan chan []byte) []byte {
func GetHttpRequestBody(url string, datachan chan []byte) []byte {

	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	log.Println(resp)
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
	datachan <- data
	return data

}

//func HttpGetBinding(url string, dataStruct interface{}) (interface{}, error) {
func HttpGetDataBinding(url string, dataStruct interface{}) error {

	// make channel for fetching url
	datachan := make(chan []byte, 1)

	go GetHttpRequestBody(url, datachan)
	err := json.Unmarshal(<-datachan, &dataStruct)

	if err != nil {
		log.Println(err)
	}

	return err
}

func HttpGetBindData(url string, dataStruct interface{}) error {

	// make channel for fetching url
	datachan := make(chan []byte, 1)

	go GetHttpRequestBody(url, datachan)
	var data []byte
	data = (<-datachan)

	log.Println("received data -->", string(data))
	// var currencymessage CurrencyMessage
	err := json.Unmarshal(data, &dataStruct)

	if err != nil {
		log.Println(err)
	}

	return err
}

// type Products struct {
// 	ID          int     `db:"id"`
// 	Name        string  `db:"name"`
// 	Model       string  `db:"model"`
// 	Price       float64 `db:"price"`
// 	Description string  `db:"description"`
// 	Image_url   string  `db:"image_url"`
// 	Video_url   string  `db:"video_url"`
// 	Capacity    int     `db:"capacity"`
// 	Created_on  int64   `db:"created_on"`
// 	Created_by  string  `db:"created_by"`
// 	Modified_on int64   `db:"modified_on"`
// 	Modified_by string  `db:"modified_by"`
// 	Labels      string  `db:"labels"`
// 	State       int     `db:"state"`
// }

// func GetProductList() ([]Products, error) {

// 	_sql := "SELECT id, name, model, price, description, image_url, video_url, capacity, created_on, created_by, modified_on, modified_by, labels, state from shop_products"
// 	rows, err := common.FetchAll(_sql)

// 	productList := make([]Products, 0)
// 	for rows.Next() {
// 		var p Products
// 		//rows.StructScan(&p.ID, &p.Name, &p.Model, &p.Price, &p.Description, &p.Image_url, &p.Video_url, &p.Capacity)
// 		rows.StructScan(&p)
// 		productList = append(productList, p)
// 	}

// 	return productList, err

// }

// //func GetProduct(p Products) ProductModel {
// //}

// func GetProductByID(id int) (Products, error) {
// 	//_sql := "SELECT id, name, price, model, description, image_url, video_url, capacity FROM shop_products WHERE id = ?"
// 	_sql := "SELECT * FROM shop_products WHERE id = ? Limit 1"

// 	db := common.Getdb()
// 	var p Products
// 	err := db.Get(&p, _sql, id)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// rows, err := common.FetchOne(_sql, p.ID)
// 	// defer rows.Close()

// 	// for rows.Next() {
// 	// 	rows.StructScan(p)
// 	// }
// 	fmt.Printf("p: %T, %v\n", p, p)
// 	return p, err
// }

// func CreateProduct(p *Products) {

// 	_sql, err := common.CreateQuery(*p, "shop_products")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//fmt.Println(_sql)
// 	db := common.Getdb()
// 	db.Exec(_sql)
// }

// func UpdateProductByID(p *Products) (sql.Result, error) {

// 	_sql := fmt.Sprintf("UPDATE shop_products SET name='%s', model='%s', price=%f, description='%s', image_url='%s', video_url='%s', capacity=%d, created_on=%d, created_by='%s', modified_on=%d, modified_by='%s',labels='%s', state=%d WHERE id=%d",
// 		p.Name, p.Model, p.Price, p.Description, p.Image_url, p.Video_url, p.Capacity,
// 		p.Created_on, p.Created_by, p.Modified_on, p.Modified_by,
// 		p.Labels, p.State, p.ID,
// 	)
// 	db := common.Getdb()
// 	log.Println(_sql)
// 	result, err := db.Exec(_sql)
// 	// result, err := db.NamedExec(`UPDATE shop_products SET name:name, model=:model, price=:price, description=:description,
// 	//  image_url=:image_url, video_url=:video_url, capacity=:capacity,
// 	//  created_on=:created_on created_by=:created_by, modified_on=:modified_on,modified_by=:modified_by,
// 	//  labels=:labels,state=:state WHERE id=:id`, p)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return result, err
// }

// func DeleteProductByID(p *Products) (sql.Result, error) {

// 	_sql := fmt.Sprintf("DELETE FROM shop_products WHERE id=%d", p.ID)
// 	db := common.Getdb()
// 	log.Println(_sql)
// 	result, err := db.Exec(_sql)
// 	// result, err := db.NamedExec(`UPDATE shop_products SET name:name, model=:model, price=:price, description=:description,
// 	//  image_url=:image_url, video_url=:video_url, capacity=:capacity,
// 	//  created_on=:created_on created_by=:created_by, modified_on=:modified_on,modified_by=:modified_by,
// 	//  labels=:labels,state=:state WHERE id=:id`, p)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return result, err
// }
