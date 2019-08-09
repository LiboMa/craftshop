package markets

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/LiboMa/otcmarket/common"
	"github.com/jasonlvhit/gocron"
)

func GetMarketData(tradeType string) (*OTCTradeMarket, error) {
	// get data from models
	var otcTradeMarket OTCTradeMarket
	// log, err := myLogger()
	requestURL := fmt.Sprintf("https://otc-api.eiijo.cn/v1/data/trade-market?country=37&currency=1&payMod=0&currPage=1&coinId=2&tradeType=%s&blockType=general&online=1", tradeType)
	err := HttpGetDataBinding(requestURL, &otcTradeMarket)
	if err != nil {
		log.Println(err)
	}

	return &otcTradeMarket, err
}

func GetHuobiMarket() (*HuobiMarket, error) {
	var huobiMarket HuobiMarket
	requestURL := "https://api.huobipro.com/market/tickers"
	err := HttpGetDataBinding(requestURL, &huobiMarket)
	if err != nil {
		log.Println(err)
	}
	return &huobiMarket, err
}

type Task struct {
	counter int
}

func (t *Task) handler(tradeType string) {

	//client := common.InitCache("")
	client := common.GetCache()
	MarketData, err := GetMarketData(tradeType)

	value, err := json.Marshal(&MarketData)

	if err != nil {
		log.Println(err)
	}
	key := fmt.Sprintf("market-price-%s", tradeType)
	err = client.Set(key, value, 0).Err()

	if err != nil {
		log.Println(err)
	}
}

func (t *Task) huobiHandler() {

	//client := common.InitCache("")
	var market MarketData
	client := common.GetCache()
	_huobiMarketData, err := GetHuobiMarket()

	value, err := json.Marshal(_huobiMarketData)

	if err != nil {
		log.Println(err)
	}
	key := fmt.Sprintf("market-huobi")
	err = client.Set(key, value, 0).Err()
	if err != nil {
		log.Println(err)
	}

	for _, m := range _huobiMarketData.Data {

		if m.Amount != 0 && m.Close != 0 {
			key := fmt.Sprintf("market-huobi-%s", m.Symbol)
			value, _ := json.Marshal(&market)
			err = client.Set(key, value, 0).Err()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func TaskRunner(duration uint64) {

	s := gocron.NewScheduler()
	var task Task
	s.Every(duration).Seconds().Do(task.handler, "buy")
	s.Every(duration).Seconds().Do(task.handler, "sell")
	s.Every(duration).Seconds().Do(task.huobiHandler)
	<-s.Start()
}
