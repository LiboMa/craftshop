package markets

import (
	"github.com/gin-gonic/gin"
)

type MarketPriceSerializer struct {
	C           *gin.Context
	MarketPrice OTCTradeMarket
	TradeType   string
}
type MarketPriceResponse struct {
	Key       string  `json:"-"`
	Price     float64 `json:"price"`
	Currency  int     `json:"currency"`
	TradeType string  `json:"type",omitedempty`
	Status    bool    `json:"status"`
}

func (m *MarketPriceSerializer) Response() *MarketPriceResponse {
	var offset float64
	if m.TradeType == "buy" {
		offset = 1 * (1 + 0.01)
	} else if m.TradeType == "sell" {
		offset = 1 * (1 - 0.01)
	}
	response := MarketPriceResponse{
		Price:     m.MarketPrice.Data[0].Price * offset,
		Currency:  m.MarketPrice.Data[0].Currency,
		Status:    m.MarketPrice.Success,
		TradeType: m.TradeType,
	}

	return &response
}

// Huobi Market Data Serializer
type HuobiMarketSerializer struct {
	C           *gin.Context
	huobiMarket *HuobiMarket
}
type HuobiMarketResponse struct {
	Status string          `json:"status"`
	Ts     int64           `json:"ts"`
	Data   []marketDataRes `json:"data"`
}

type marketDataRes struct {
	Price  float64 `json:"price"`
	Symbol string  `json:"symbol"`
}

func (m *HuobiMarketSerializer) Response() *HuobiMarketResponse {
	var data marketDataRes
	newData := make([]marketDataRes, len(m.huobiMarket.Data))
	for _, _m := range m.huobiMarket.Data {

		data.Price = _m.Close
		data.Symbol = _m.Symbol

		// log.Println("Before-->", data.Price)
		if data.Price == 0 || data.Symbol == "" {
			continue
		}
		// log.Println("after-->", data.Price)
		newData = append(newData, data)
	}
	response := HuobiMarketResponse{
		Status: m.huobiMarket.Status,
		Ts:     m.huobiMarket.Ts,
		Data:   newData,
	}

	return &response
}
