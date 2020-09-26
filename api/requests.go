package api

import "github.com/shopspring/decimal"

/*
PUBLIC API GET REQUESTS
*/

// GetOrderBookRequest is the request struct for GetOrderBook.
type GetOrderBookRequest struct {
	// https://api.valr.com/v1/public/:currencyPair/orderbook
	// Currency pair
	// required: true
	Pair string `json:"-" url:"currencyPair"`
}

// GetCurrenciesRequest is the request struct for GetCurrencies
type GetCurrenciesRequest struct {
	// https://api.valr.com/v1/public/currencies
	// Empty
}

// GetCurrencyPairsRequest is the request struct for GetCurrencyPairs
type GetCurrencyPairsRequest struct {
	// https://api.valr.com/v1/public/pairs
	// Empty
}

// GetOrderTypesRequest is the request struct for GetOrderTypes
type GetOrderTypesRequest struct {
	// https://api.valr.com/v1/public/ordertypes
	// Empty
}

// GetOrderTypesForPairRequest is the request struct for GetOrderTypesForPair
type GetOrderTypesForPairRequest struct {
	// https://api.valr.com/v1/public/:currencyPair/ordertypes
	// Currency Pair
	// required: true
	Pair string `json:"-" url:"currencyPair"`
}

// GetMarketSummaryRequest is the request struct for GetMarketSummary
type GetMarketSummaryRequest struct {
	// https://api.valr.com/v1/public/marketsummary
	// Empty
}

// GetMarketSummaryForPairRequest is the request struct for GetMarketSummaryForPair
type GetMarketSummaryForPairRequest struct {
	// https://api.valr.com/v1/public/:currencyPair/marketsummary
	// Currency Pair
	// required: true
	Pair string `json:"-" url:"currencyPair"`
}

// GetServerTimeRequest is the request struct for GetServerTime
type GetServerTimeRequest struct {
	// https://api.valr.com/v1/public/time
	// Empty
}

/*
PRIVATE API GET REQUESTS
*/

// GetAccountBalancesRequest is the request struct for GetAccountBalances
type GetAccountBalancesRequest struct {
	// https://api.valr.com/v1/account/balances
	// Empty
}

// GetTransactionHistoryRequest is the request struct for GetTransactionHistory
type GetTransactionHistoryRequest struct {
	// https://api.valr.com/v1/account/transactionhistory?skip=0&limit=100
	// Skip
	// Limit
	// required: false
	Skip  int `json:"-" url:"skip"`
	Limit int `json:"-" url:"limit"`
}

// GetTradeHistoryForPairRequest is the request struct for GetTradeHistoryForPair
type GetTradeHistoryForPairRequest struct {
	// https://api.valr.com/v1/account/:currencyPair/tradehistory?limit=10
	// Currency Pair
	// required: true
	// Limit
	// required: false
	Pair  string `json:"-" url:"currencyPair"`
	Limit int    `json:"-" url:"limit"`
}

// GetDepositAddressRequest is the request struct for GetDepositAddress
type GetDepositAddressRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/deposit/address
	// Currency Code
	// required: true
	Asset string `json:"-" url:"currencyCode"`
}

// GetWithdrawInfoRequest is the request struct for GetWithdrawInfo
type GetWithdrawInfoRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw
	// Currency Code
	// required: true
	Asset string `json:"-" url:"currencyCode"`
}

// GetWithdrawStatusRequest is the request struct for GetWithdrawStatus
type GetWithdrawStatusRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw/:withdrawId
	// Currency Code
	// Withdraw ID
	// required: true
	Asset string `json:"-" url:"currencyCode"`
	ID    string `json:"-" url:"withdrawId"`
}

// GetDepositHistoryForAssetRequest is the request struct for GetDepositHistoryForAsset
type GetDepositHistoryForAssetRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/deposit/history?skip=0&limit=10
	// Currency Code
	// required: true
	// Skip
	// Limit
	// required: false
	Asset string `json:"-" url:"currencyCode"`
	Skip  int    `json:"-" url:"skip"`
	Limit int    `json:"-" url:"limit"`
}

// GetWithdrawHistoryForAssetRequest is the request struct for GetWithdrawHistoryForAsset
type GetWithdrawHistoryForAssetRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw/history?skip=0&limit=10
	// Currency Code
	// required: true
	// Skip
	// Limit
	// required: false
	Asset string `json:"-" url:"currencyCode"`
	Skip  int    `json:"-" url:"skip"`
	Limit int    `json:"-" url:"limit"`
}

// GetBankAccountForAssetRequest is the request struct for GetBankAccountForAsset
type GetBankAccountForAssetRequest struct {
	// https://api.valr.com/v1/wallet/fiat/:currencyCode/accounts
	// Currency Code
	// required: true
	Asset string `json:"-" url:"currencyCode"` // note: ZAR is the only supported asset at this time
}

// GetAuthOrderBookRequest is the request struct for GetAuthOrderBook
type GetAuthOrderBookRequest struct {
	// https://api.valr.com/v1/marketdata/:currencyPair/orderbook
	// Currency pair
	// required: true
	Pair string `json:"-" url:"currencyPair"`
}

// GetAuthFullOrderBookRequest is the request struct for GetAuthFullOrderBook
type GetAuthFullOrderBookRequest struct {
	// https://api.valr.com/v1/marketdata/:currencyPair/orderbook/full
	// Currency pair
	// required: true
	Pair string `json:"-" url:"currencyPair"`
}

// GetAuthTradeHistoryForPairRequest is the request struct for GetAuthTradeHistoryForPair
type GetAuthTradeHistoryForPairRequest struct {
	// https://api.valr.com/v1/marketdata/:currencyPair/tradehistory?limit=10
	// Currency Pair
	// required: true
	// Limit
	// required: false
	Pair  string `json:"-" url:"currencyPair"`
	Limit int    `json:"-" url:"limit"`
}

// GetSimpleBuyOrSellOrderStatusRequest is the request struct for GetSimpleBuyOrSellOrderStatus
type GetSimpleBuyOrSellOrderStatusRequest struct {
	// https://api.valr.com/v1/simple/:currencyPair/order/:orderId
	// Currency Pair
	// Order ID
	// required: true
	Pair string `json:"-" url:"currencyPair"`
	ID   string `json:"-" url:"orderId"`
}

// GetOrderStatusByOrderIDRequest is the request struct for GetOrderStatusByOrderID
type GetOrderStatusByOrderIDRequest struct {
	// https://api.valr.com/v1/orders/:currencyPair/orderid/:orderId
	// Currency Pair
	// Order ID
	// required: true
	Pair string `json:"-" url:"currencyPair"`
	ID   string `json:"-" url:"orderId"`
}

// GetOrderStatusByCustomerOrderIDRequest is the request struct for GetOrderStatusByCustomerOrderID
type GetOrderStatusByCustomerOrderIDRequest struct {
	// https://api.valr.com/v1/orders/:currencyPair/customerorderid/:customerOrderId
	// Currency Pair
	// Customer Order ID
	// required: true
	Pair string `json:"-" url:"currencyPair"`
	ID   string `json:"-" url:"customerOrderId"`
}

// GetAllOpenOrdersRequest is the request struct for GetAllOpenOrders
type GetAllOpenOrdersRequest struct {
	// https://api.valr.com/v1/orders/open
	// Empty
}

// GetOrderHistoryRequest is the request struct for GetOrderHistory
type GetOrderHistoryRequest struct {
	// https://api.valr.com/v1/orders/history?skip=0&limit=2
	// Skip
	// Limit
	// required: false
	Skip  int `json:"-" url:"skip"`
	Limit int `json:"-" url:"limit"`
}

// GetOrderHistorySummaryByOrderIDRequest is the request struct for GetOrderHistorySummaryByOrderID
type GetOrderHistorySummaryByOrderIDRequest struct {
	// https://api.valr.com/v1/orders/history/summary/orderid/:orderId
	// Order ID
	// required: true
	ID string `json:"-" url:"orderId"`
}

// GetOrderHistorySummaryByCustomerOrderIDRequest is the request struct for GetOrderHistorySummaryByCustomerOrderID
type GetOrderHistorySummaryByCustomerOrderIDRequest struct {
	// https://api.valr.com/v1/orders/history/summary/customerorderid/:customerOrderId
	// Customer Order ID
	// required: true
	ID string `json:"-" url:"customerOrderId"`
}

// GetOrderHistoryDetailsByOrderIDRequest is the request struct for GetOrderHistoryDetailsByOrderID
type GetOrderHistoryDetailsByOrderIDRequest struct {
	// https://api.valr.com/v1/orders/history/detail/orderid/:orderId
	// Order ID
	// required: true
	ID string `json:"-" url:"orderId"`
}

// GetOrderHistoryDetailsByCustomerOrderIDRequest is the request struct for GetOrderHistoryDetailsByCustomerOrderID
type GetOrderHistoryDetailsByCustomerOrderIDRequest struct {
	// https://api.valr.com/v1/orders/history/detail/customerorderid/:customerOrderId
	// Order ID
	// required: true
	ID string `json:"-" url:"customerOrderId"`
}

/*
PRIVATE API POST REQUESTS
*/

// PostNewCryptoWithdrawRequest is the request struct for PostNewCryptoWithdraw
type PostNewCryptoWithdrawRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw
	// Currency Code
	// Amount
	// Address
	// required: true
	Asset   string          `json:"currencyCode" url:"-"`
	Amount  decimal.Decimal `json:"amount" url:"-"`
	Address string          `json:"address" url:"-"`
}

// PostNewFiatWithdrawRequest is the request struct for PostNewFiatWithdraw
type PostNewFiatWithdrawRequest struct {
	// https://api.valr.com/v1/wallet/fiat/:currencyCode/withdraw
	// Currency Code
	// Amount
	// BankAccountID
	// required: true
	Asset         string          `json:"currencyCode" url:"-"`
	Amount        decimal.Decimal `json:"amount" url:"-"`
	BankAccountID string          `json:"linkedBankAccountId" url:"-"`
}

// PostSimpleBuyOrSellQuoteRequest is the request stuct for PostSimpleBuyOrSellQuote
type PostSimpleBuyOrSellQuoteRequest struct {
	// https://api.valr.com/v1/simple/:currencyPair/quote
	// Currency Pair
	// Pay in currency
	// Pay amount
	// Side
	// required: true
	Pair          string          `json:"currencyPair" url:"-"`
	PayInCurrency string          `json:"payInCurrency" url:"-"`
	PayAmount     decimal.Decimal `json:"payAmount" url:"-"`
	Side          RequestSide     `json:"side" url:"-"`
}

// PostSimpleBuyOrSellOrderRequest is the request struct for PostSimpleBuyOrSellOrder
type PostSimpleBuyOrSellOrderRequest struct {
	// https://api.valr.com/v1/simple/:currencyPair/order
	// Currency Pair
	// Pay in currency
	// Pay amount
	// Side
	// required: true
	Pair          string          `json:"currencyPair" url:"-"`
	PayInCurrency string          `json:"payInCurrency" url:"-"`
	PayAmount     decimal.Decimal `json:"payAmount" url:"-"`
	Side          RequestSide     `json:"side" url:"-"`
}

// PostLimitOrderRequest is the request struct for PostLimitOrder
type PostLimitOrderRequest struct {
	// https://api.valr.com/v1/orders/limit
	// Currency Pair
	// Quantity
	// Price
	// Side
	// required: true
	// Post Only
	// Customer Order ID
	// required: false
	Pair            string          `json:"pair" url:"-"`
	Quantity        decimal.Decimal `json:"quantity" url:"-"`
	Price           decimal.Decimal `json:"price" url:"-"`
	Side            RequestSide     `json:"side" url:"-"`
	PostOnly        bool            `json:"postOnly" url:"-"`
	CustomerOrderID string          `json:"customerOrderId" url:"-"`
}

// PostMarketOrderBuyRequest is the request struct for PostMarketOrder
type PostMarketOrderBuyRequest struct {
	// https://api.valr.com/v1/orders/market
	// Currency Pair
	// Quantity
	// Side
	// required: true
	// Customer Order ID
	// required: false
	Side            RequestSide     `json:"side" url:"-"`
	Quantity        decimal.Decimal `json:"quoteAmount" url:"-"`
	Pair            string          `json:"pair" url:"-"`
	CustomerOrderID string          `json:"customerOrderId" url:"-"`
}

// PostMarketOrderBuyRequest is the request struct for PostMarketOrder
type PostMarketOrderSellRequest struct {
	// https://api.valr.com/v1/orders/market
	// Currency Pair
	// Quantity
	// Side
	// required: true
	// Customer Order ID
	// required: false
	Side            RequestSide     `json:"side" url:"-"`
	Quantity        decimal.Decimal `json:"baseAmount" url:"-"`
	Pair            string          `json:"pair" url:"-"`
	CustomerOrderID string          `json:"customerOrderId" url:"-"`
}

/*
PRIVATE API DEL REQUESTS
*/

// DelOrderRequest is the request struct for DelOrder
type DelOrderRequest struct {
	// https://api.valr.com/v1/orders/order
	// Currency Pair
	// Order ID
	// required: true
	Pair string `json:"pair" url:"-"`
	ID   string `json:"orderId" url:"-"`
}

// DelOrderByCustomerOrderIDRequest is the request struct for DelOrderByCustomerOrderID
type DelOrderByCustomerOrderIDRequest struct {
	// https://api.valr.com/v1/orders/order
	// Currency Pair
	// required: true
	Pair string `json:"pair" url:"-"`
	// Customer Order ID
	// required: true
	ID string `json:"customerOrderId" url:"-"`
}
