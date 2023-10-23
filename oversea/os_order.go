package oversea

// Some of the tickers does not support mock-order
// Order Requests outside Foreign Exchange will result in errors
// Check the time
// US: 23:30 ~ 06:00 (Summer time: 22:30 ~ 05:00)
// JP: 09:00 ~ 11:30, 12:30 ~ 15:00
// CN: 10:30 ~ 16:00
// HK: 10:30 ~ 13:00, 14:00 ~ 17:00
// For POST API, Body's key value must be printed in Capital letters

type OrderHeader struct {
	Auth          string    `json:"authorization"`
	AppKey        string    `json:"appkey"`
	AppSecret     string    `json:"appsecret"`
	TransactionID OrderType `json:"tr_id"`
	CustomerType  string    `json:"custtype,omitempty"` // For individuals, use P
	HashKey       string    `json:"hashkey"`            // For POST API. Request Body to hash
}

type OrderBody struct {
	AccountNumber      string        `json:"CANO"`         // First 8 letters of the account
	AccountProductCode string        `json:"ACNT_PRDT_CD"` // Last 2 number of the account
	ExchangeCode       ExchangeType  `json:"OVRS_EXCG_CD"`
	ProductNumber      string        `json:"PDNO"`
	OrderQuantity      string        `json:"ORD_QTY"`
	OverseaOrderPer    string        `json:"OVRS_ORD_UNPR"`   // Price per 1 stock
	OrderServerDivCode string        `json:"ORD_SVR_DVSN_CD"` // Default to "0"
	OrderDiv           OrderDivision `json:"ORD_DVSN"`
}

type OrderResponseHeader struct {
	TransactionID       string `json:"tr_id"`
	TransactionContinue string `json:"tr_cont"`
	GlobalUID           string `json:"gt_uid"` // UUID for each transaction
}

type OrderResponseBody struct {
	RequestSuccess string `json:"rt_cd"` // 0 for success, else failed
	MessageCode    string `json:"msg_cd"`
	Message        string `json:"msg1"`
	Output         struct {
		KExchangeOrderNumber string `json:"KRX_RWDG_ORD_ORGNO"`
		OrderNumber          string `json:"ODNO"`
		OrderTime            string `json:"ORD_TMD"` // HHMMSS
	} `json:"output"`
}

// OrderDivision Describes the type of order possible for each market
// Only for US order - tr_id is TTTT1002U or TTTT1006U
type OrderDivision = string

const (
	LimitOrderBuy       = "00" // 지정가 - 모의
	LimitOrderOpenBuy   = "32" // 장개시지정가
	LimitOrderOpenClose = "34" // 장마감지정가

	LimitOrderSell       = "00" // 지정가 - 모의
	MarketOrderOpenSell  = "31" // 장개시시장가
	LimitOrderOpenSell   = "32" // 장개시지정가
	MarketOrderCloseSell = "33" // 장마감시장가
	LimitOrderCloseSell  = "34" // 장마감지정가
)

// ExchangeType Describe which oversea exchange user is ordering from
type ExchangeType = string

const (
	ExchangeNasdaq   = "NASD" // Nasdaq
	ExchangeNYSE     = "NYSE" // New York Stock Exchange
	ExchangeAmex     = "AMEX" // Amex
	ExchangeHongKong = "SEHK" // Hong Kong
	ExchangeSc       = "SZAA" // Shim Chun
	ExchangeShanghai = "SHAA" // ShangHai
	ExchangeJapan    = "TKSE" // Japan Tokyo
)

// OrderType as enum
type OrderType = string

const (
	BuyOrderUS OrderType = "TTTT1002U" // 미국 매수 주문
	BuyOrderJP OrderType = "TTTS0308U" // 일본 매수 주문
	BuyOrderSH OrderType = "TTTS0202U" // 상해 매수 주문
	BuyOrderHK OrderType = "TTTS1002U" // 홍콩 매수 주문
	BuyOrderSC OrderType = "TTTS0305U" // 심천 매수 주문
	BuyOrderVN OrderType = "TTTS0311U" // 베트남 매수 주문

	MockBuyOrderUS OrderType = "VTTT1002U" // 미국 매수 주문
	MockBuyOrderJP OrderType = "VTTS0308U" // 일본 매수 주문
	MockBuyOrderSH OrderType = "VTTS0202U" // 상해 매수 주문
	MockBuyOrderHK OrderType = "VTTS1002U" // 홍콩 매수 주문
	MockBuyOrderSC OrderType = "VTTS0305U" // 심천 매수 주문
	MockBuyOrderVN OrderType = "VTTS0311U" // 베트남 매수 주문
)

const (
	SellOrderUS OrderType = "TTTT1006U" // 미국 매도 주문
	SellOrderJP OrderType = "TTTS0307U" // 일본 매도 주문
	SellOrderSH OrderType = "TTTS1005U" // 상해 매도 주문
	SellOrderHK OrderType = "TTTS1001U" // 홍콩 매도 주문
	SellOrderSC OrderType = "TTTS0304U" // 심천 매도 주문
	SellOrderVN OrderType = "TTTS0310U" // 베트남 매도 주문

	MockSellOrderUS OrderType = "VTTT1001U" // 미국 매도 주문
	MockSellOrderJP OrderType = "VTTS0307U" // 일본 매도 주문
	MockSellOrderSH OrderType = "VTTS1005U" // 상해 매도 주문
	MockSellOrderHK OrderType = "VTTS1001U" // 홍콩 매도 주문
	MockSellOrderSC OrderType = "VTTS0304U" // 심천 매도 주문
	MockSellOrderVN OrderType = "VTTS0310U" // 베트남 매도 주문
)
