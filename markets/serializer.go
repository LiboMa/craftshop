package markets

import "github.com/gin-gonic/gin"

type MarketPriceSerializer struct {
	C           *gin.Context
	MarketPrice OTCTradeMarket
	TradeType   string
}
type MarketPriceResponse struct {
	Key       string  `json:"-"`
	Price     float64 `json:"price"`
	Currency  int     `json:"currency"`
	TradeType string  `json:"type", omitedempty`
	Status    bool    `json:"status"`
}

func (m *MarketPriceSerializer) Response() *MarketPriceResponse {

	response := MarketPriceResponse{
		Price:     m.MarketPrice.Data[0].Price,
		Currency:  m.MarketPrice.Data[0].Currency,
		Status:    m.MarketPrice.Success,
		TradeType: m.TradeType,
	}

	return &response
}

// Huobi Market Data Serializer
type HuobiMarketSerializer struct {
	C           *gin.Context
	huobiMarket HuobiMarket
}
type HuobiMarketResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []*MarketData
}

func (m *HuobiMarketSerializer) Response() *HuobiMarketResponse {

	response := HuobiMarketResponse{
		Status: m.huobiMarket.Status,
		Ts:     m.huobiMarket.Ts,
		Data:   m.huobiMarket.Data,
	}

	return &response
}
