package api

import (
	"time"

	"github.com/shopspring/decimal"
)

/*
PUBLIC API GET RESPONSES
*/

type GetOrderBookResponse struct {
	OrderBook
}

type GetCurrenciesResponse struct {
	Currencies []CurrencyInfo
}

type GetCurrencyPairsResponse struct {
	CurrencyPairs []PairInfo
}

type GetOrderTypesResponse struct {
	OrderTypes []OrderTypes
}

type GetOrderTypesForPairResponse struct {
	OrderTypes []string
}

type GetMarketSummaryResponse struct {
	Markets []MarketSummary
}

type GetMarketSummaryForPairResponse struct {
	MarketSummary
}

type GetServerTimeResponse struct {
	EpochTime int       `json:"epochTime"`
	Time      time.Time `json:"time"`
}

/*
PRIVATE API GET RESPONSES
*/

type GetAccountBalancesResponse struct {
	Balances []AccountBalance
}

type GetTransactionHistoryResponse struct {
	Transactions []TransactionInfo
}

type GetTradeHistoryForPairResponse struct {
	Trades []TradeInfo
}

type GetDepositAddressResponse struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
}

type GetWithdrawInfoResponse struct {
	Currency                 string          `json:"currency"`
	MinimumWithdrawAmount    decimal.Decimal `json:"minimumWithdrawAmount"`
	Active                   bool            `json:"isActive"`
	WithdrawCost             decimal.Decimal `json:"withdrawCost"`
	SupportsPaymentReference bool            `json:"supportsPaymentReference"`
}

type GetWithdrawStatusResponse struct {
	WithdrawInfo
}

type GetDepositHistoryForAssetResponse struct {
	Deposits []DepositInfo
}

type GetWithdrawHistoryForAssetResponse struct {
	Withdraws []WithdrawInfo
}

type GetBankAccountForAssetResponse struct {
	BankAccounts []BankInfo
}

type GetAuthOrderBookResponse struct {
	GetOrderBookResponse
}

type GetAuthFullOrderBookResponse struct {
	OrderBook
}

type GetAuthTradeHistoryForPairResponse struct {
	Trades []TradeHistoryInfo
}

type GetSimpleBuyOrSellOrderStatusResponse struct {
	OrderID         string          `json:"orderId"`
	Success         bool            `json:"success"`
	Processing      bool            `json:"processing"`
	PaidAmount      decimal.Decimal `json:"paidAmount"`
	PaidCurrency    string          `json:"paidCurrency"`
	ReceiveAmount   decimal.Decimal `json:"receivedAmount"`
	FeeAmount       decimal.Decimal `json:"feeAmount"`
	FeeCurrency     string          `json:"feeCurrency"`
	OrderExecutedAt time.Time       `json:"orderExecutedAt"`
}

type GetOrderStatusByOrderIDResponse struct {
	OrderID           string          `json:"orderId"`
	OrderStatusType   bool            `json:"orderStatusType"`
	CurrencyPair      string          `json:"currencyPair"`
	OriginalPrice     decimal.Decimal `json:"originalPrice"`
	RemainingQuantity decimal.Decimal `json:"remainingQuantity"`
	OriginalQuantity  decimal.Decimal `json:"originalQuantity"`
	OrderSide         ResponseSide    `json:"orderSide"`
	OrderType         string          `json:"orderType"`
	FailedReason      string          `json:"failedReason"`
	CustomerOrderID   string          `json:"customerOrderId"`
	OrderUpdatedAt    time.Time       `json:"orderUpdatedAt"`
	OrderCreatedAt    time.Time       `json:"orderCreatedAt"`
}

type GetOrderStatusByCustomerOrderIDResponse struct {
	GetOrderStatusByOrderIDResponse
}

type GetAllOpenOrdersResponse struct {
	Orders []OpenOrder
}

type GetOrderHistoryResponse struct {
	Orders []OrderReceipt
}

type GetOrderHistorySummaryByOrderIDResponse struct {
	OrderID           string          `json:"orderId"`
	OrderStatusType   bool            `json:"orderStatusType"`
	Pair              string          `json:"currencyPair"`
	AveragePrice      decimal.Decimal `json:"averagePrice"`
	OriginalPrice     decimal.Decimal `json:"originalPrice"`
	RemainingQuantity decimal.Decimal `json:"remainingQuantity"`
	OriginalQuantity  decimal.Decimal `json:"originalQuantity"`
	Total             decimal.Decimal `json:"total"`
	TotalFee          decimal.Decimal `json:"totalFee"`
	FeeCurrency       string          `json:"feeCurrency"`
	OrderSide         ResponseSide    `json:"orderSide"`
	OrderType         string          `json:"orderType"`
	FailedReason      string          `json:"failedReason"`
	OrderUpdatedAt    time.Time       `json:"orderUpdatedAt"`
	OrderCreatedAt    time.Time       `json:"orderCreatedAt"`
}

type GetOrderHistorySummaryByCustomerOrderIDResponse struct {
	OrderID           string          `json:"orderId"`
	CustomerOrderID   string          `json:"customerOrderId"`
	OrderStatusType   bool            `json:"orderStatusType"`
	Pair              string          `json:"currencyPair"`
	AveragePrice      decimal.Decimal `json:"averagePrice"`
	OriginalPrice     decimal.Decimal `json:"originalPrice"`
	RemainingQuantity decimal.Decimal `json:"remainingQuantity"`
	OriginalQuantity  decimal.Decimal `json:"originalQuantity"`
	Total             decimal.Decimal `json:"total"`
	TotalFee          decimal.Decimal `json:"totalFee"`
	FeeCurrency       string          `json:"feeCurrency"`
	OrderSide         ResponseSide    `json:"orderSide"`
	OrderType         string          `json:"orderType"`
	FailedReason      string          `json:"failedReason"`
	OrderUpdatedAt    time.Time       `json:"orderUpdatedAt"`
	OrderCreatedAt    time.Time       `json:"orderCreatedAt"`
}

type GetOrderHistoryDetailsByOrderIDResponse struct {
	OrderStatuses []OrderStatus
}

type GetOrderHistoryDetailsByCustomerOrderIDResponse struct {
	OrderStatuses []OrderStatus
}

/*
PRIVATE API POST RESPONSES
*/

type PostNewCryptoWithdrawResponse struct {
	IDResponse
}

type PostNewFiatWithdrawResponse struct {
	IDResponse
}

type PostSimpleBuyOrSellQuoteResponse struct {
	Pair          string          `json:"currencyPair"`
	PayAmount     decimal.Decimal `json:"payAmount"`
	ReceiveAmount decimal.Decimal `json:"receiveAmount"`
	Fee           decimal.Decimal `json:"fee"`
	FeeCurrency   string          `json:"feeCurrency"`
	CreatedAt     time.Time       `json:"createdAt"`
	OrderID       string          `json:"id"`
}

type PostSimpleBuyOrSellOrderResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PostLimitOrderResponse struct {
	IDResponse
}

type PostMarketOrderResponse struct {
	IDResponse
}

/*
PRIVATE API DEL RESPONSES
*/

type DelOrderResponse struct {
	// Empty 202 Response
}
