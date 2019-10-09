package api

import (
	"time"

	"github.com/shopspring/decimal"
)

/*
PUBLIC API GET RESPONSES
*/

// GetOrderBookResponse is the struct that GetOrderBook responses are unpacked into
type GetOrderBookResponse struct {
	OrderBook
}

// GetCurrenciesResponse is the struct that GetCurrencies responses are unpacked into
type GetCurrenciesResponse struct {
	Currencies []CurrencyInfo
}

// GetCurrencyPairsResponse is the struct that GetCurrencyPairs responses are unpacked into
type GetCurrencyPairsResponse struct {
	CurrencyPairs []PairInfo
}

// GetOrderTypesResponse is the struct that GetOrderTypes responses are unpacked into
type GetOrderTypesResponse struct {
	OrderTypes []OrderTypes
}

// GetOrderTypesForPairResponse is the struct that GetOrderTypesForPair responses are unpacked into
type GetOrderTypesForPairResponse struct {
	OrderTypes []string
}

// GetMarketSummaryResponse is the struct that GetMarketSummary responses are unpacked into
type GetMarketSummaryResponse struct {
	Markets []MarketSummary
}

// GetMarketSummaryForPairResponse is the struct that GetMarketSummaryForPair responses are unpacked into
type GetMarketSummaryForPairResponse struct {
	MarketSummary
}

// GetServerTimeResponse is the struct that GetServerTime responses are unpacked into
type GetServerTimeResponse struct {
	EpochTime int       `json:"epochTime"`
	Time      time.Time `json:"time"`
}

/*
PRIVATE API GET RESPONSES
*/

// GetAccountBalancesResponse is the struct that GetAccountBalances responses are unpacked into
type GetAccountBalancesResponse struct {
	Balances []AccountBalance
}

// GetTransactionHistoryResponse is the struct that GetTransactionHistory responses are unpacked into
type GetTransactionHistoryResponse struct {
	Transactions []TransactionInfo
}

// GetTradeHistoryForPairResponse is the struct that GetTradeHistoryForPair responses are unpacked into
type GetTradeHistoryForPairResponse struct {
	Trades []TradeInfo
}

// GetDepositAddressResponse is the struct that GetDepositAddress responses are unpacked into
type GetDepositAddressResponse struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
}

// GetWithdrawInfoResponse is the struct that GetWithdrawInfo responses are unpacked into
type GetWithdrawInfoResponse struct {
	Currency                 string          `json:"currency"`
	MinimumWithdrawAmount    decimal.Decimal `json:"minimumWithdrawAmount"`
	Active                   bool            `json:"isActive"`
	WithdrawCost             decimal.Decimal `json:"withdrawCost"`
	SupportsPaymentReference bool            `json:"supportsPaymentReference"`
}

// GetWithdrawStatusResponse is the struct that GetWithdrawStatus responses are unpacked into
type GetWithdrawStatusResponse struct {
	WithdrawInfo
}

// GetDepositHistoryForAssetResponse is the struct that GetDepositHistoryForAsset responses are unpacked into
type GetDepositHistoryForAssetResponse struct {
	Deposits []DepositInfo
}

// GetWithdrawHistoryForAssetResponse is the struct that GetWithdrawHistoryForAsset responses are unpacked into
type GetWithdrawHistoryForAssetResponse struct {
	Withdraws []WithdrawInfo
}

// GetBankAccountForAssetResponse is the struct that GetBankAccountForAsset responses are unpacked into
type GetBankAccountForAssetResponse struct {
	BankAccounts []BankInfo
}

// GetAuthOrderBookResponse is the struct that GetAuthOrderBook responses are unpacked into
type GetAuthOrderBookResponse struct {
	GetOrderBookResponse
}

// GetAuthFullOrderBookResponse is the struct that GetAuthFullOrderBook responses are unpacked into
type GetAuthFullOrderBookResponse struct {
	OrderBook
}

// GetAuthTradeHistoryForPairResponse is the struct that GetAuthTradeHistoryForPair responses are unpacked into
type GetAuthTradeHistoryForPairResponse struct {
	Trades []TradeHistoryInfo
}

// GetSimpleBuyOrSellOrderStatusResponse is the struct that GetSimpleBuyOrSellOrderStatus responses are unpacked into
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

// GetOrderStatusByOrderIDResponse is the struct that GetOrderStatusByOrderID responses are unpacked into
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

// GetOrderStatusByCustomerOrderIDResponse is the struct that GetOrderStatusByCustomerOrderID responses are unpacked into
type GetOrderStatusByCustomerOrderIDResponse struct {
	GetOrderStatusByOrderIDResponse
}

// GetAllOpenOrdersResponse is the struct that GetAllOpenOrders responses are unpacked into
type GetAllOpenOrdersResponse struct {
	Orders []OpenOrder
}

// GetOrderHistoryResponse is the struct that GetOrderHistory responses are unpacked into
type GetOrderHistoryResponse struct {
	Orders []OrderReceipt
}

// GetOrderHistorySummaryByOrderIDResponse is the struct that GetOrderHistorySummaryByOrderID responses are unpacked into
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

// GetOrderHistorySummaryByCustomerOrderIDResponse is the struct that GetOrderHistorySummaryByCustomerOrderID responses are unpacked into
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

// GetOrderHistoryDetailsByOrderIDResponse is the struct that GetOrderHistoryDetailsByOrderID responses are unpacked into
type GetOrderHistoryDetailsByOrderIDResponse struct {
	OrderStatuses []OrderStatus
}

// GetOrderHistoryDetailsByCustomerOrderIDResponse is the struct that GetOrderHistoryDetailsByCustomerOrderID responses are unpacked into
type GetOrderHistoryDetailsByCustomerOrderIDResponse struct {
	OrderStatuses []OrderStatus
}

/*
PRIVATE API POST RESPONSES
*/

// PostNewCryptoWithdrawResponse is the struct that PostNewCryptoWithdraw responses are unpacked into
type PostNewCryptoWithdrawResponse struct {
	IDResponse
}

// PostNewFiatWithdrawResponse is the struct that PostNewFiatWithdraw responses are unpacked into
type PostNewFiatWithdrawResponse struct {
	IDResponse
}

// PostSimpleBuyOrSellQuoteResponse is the struct that PostSimpleBuyOrSellQuote responses are unpacked into
type PostSimpleBuyOrSellQuoteResponse struct {
	Pair          string          `json:"currencyPair"`
	PayAmount     decimal.Decimal `json:"payAmount"`
	ReceiveAmount decimal.Decimal `json:"receiveAmount"`
	Fee           decimal.Decimal `json:"fee"`
	FeeCurrency   string          `json:"feeCurrency"`
	CreatedAt     time.Time       `json:"createdAt"`
	OrderID       string          `json:"id"`
}

// PostSimpleBuyOrSellOrderResponse is the struct that PostSimpleBuyOrSellOrder responses are unpacked into
type PostSimpleBuyOrSellOrderResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// PostLimitOrderResponse is the struct that PostLimitOrder responses are unpacked into
type PostLimitOrderResponse struct {
	IDResponse
}

// PostMarketOrderResponse is the struct that PostMarketOrder responses are unpacked into
type PostMarketOrderResponse struct {
	IDResponse
}

/*
PRIVATE API DEL RESPONSES
*/

// DelOrderResponse is the struct that DelOrder responses are unpacked into
type DelOrderResponse struct {
	// Empty 202 Response
}
