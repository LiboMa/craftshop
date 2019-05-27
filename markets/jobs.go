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

func TaskRunner(duration uint64) {

	s := gocron.NewScheduler()
	var task Task
	s.Every(duration).Seconds().Do(task.handler, "buy")
	s.Every(duration).Seconds().Do(task.handler, "sell")
	<-s.Start()
}
