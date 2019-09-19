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
	Pair string `json:"currencyPair" url:"currencyPair"`
}

type GetCurrenciesRequest struct {
	// https://api.valr.com/v1/public/currencies
	// Empty
}

type GetCurrencyPairsRequest struct {
	// https://api.valr.com/v1/public/pairs
	// Empty
}

type GetOrderTypesRequest struct {
	// https://api.valr.com/v1/public/ordertypes
	// Empty
}

type GetOrderTypesForPairRequest struct {
	// https://api.valr.com/v1/public/:currencyPair/ordertypes
	// Currency Pair
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
}

type GetMarketSummaryRequest struct {
	// https://api.valr.com/v1/public/marketsummary
	// Empty
}

type GetMarketSummaryForPairRequest struct {
	// https://api.valr.com/v1/public/:currencyPair/marketsummary
	// Currency Pair
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
}

type GetServerTimeRequest struct {
	// https://api.valr.com/v1/public/time
	// Empty
}


/*
PRIVATE API GET REQUESTS
 */


type GetAccountBalancesRequest struct {
	// https://api.valr.com/v1/account/balances
	// Empty
}

type GetTransactionHistoryRequest struct {
	// https://api.valr.com/v1/account/transactionhistory?skip=0&limit=100
	// Skip
	// Limit
	// required: false
	Skip int `json:"skip" param:"currencyPair"`
	Limit int `json:"limit" param:"skip"`
}

type GetTradeHistoryForPairRequest struct {
	// https://api.valr.com/v1/account/:currencyPair/tradehistory?limit=10
	// Currency Pair
	// required: true
	// Limit
	// required: false
	Pair string `json:"currencyPair" url:"currencyPair"`
	Limit int   `json:"limit" param:"limit"`
}

type GetDepositAddressRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/deposit/address
	// Currency Code
	// required: true
	Asset string `json:"currencyCode" url:"currencyCode"`
}

type GetWithdrawInfoRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw
	// Currency Code
	// required: true
	Asset string `json:"currencyCode" url:"currencyCode"`
}

type GetWithdrawStatusRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw/:withdrawId
	// Currency Code
	// Withdraw ID
	// required: true
	Asset string `json:"currencyCode" url:"currencyCode"`
	ID    string `json:"withdrawId" url:"withdrawId"`
}

type GetDepositHistoryForAssetRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/deposit/history?skip=0&limit=10
	// Currency Code
	// required: true
	// Skip
	// Limit
	// required: false
	Asset string `json:"currencyCode" url:"currencyCode"`
	Skip  int    `json:"skip" param:"skip"`
	Limit int    `json:"limit" param:"limit"`
}

type GetWithdrawHistoryForAssetRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw/history?skip=0&limit=10
	// Currency Code
	// required: true
	// Skip
	// Limit
	// required: false
	Asset string `json:"currencyCode" url:"currencyCode"`
	Skip  int    `json:"skip" param:"skip"`
	Limit int    `json:"limit" param:"limit"`
}

type GetBankAccountForAssetRequest struct {
	// https://api.valr.com/v1/wallet/fiat/:currencyCode/accounts
	// Currency Code
	// required: true
	Asset string `json:"currencyCode" url:"currencyCode"`
}

type GetAuthOrderBookRequest struct {
	// https://api.valr.com/v1/marketdata/:currencyPair/orderbook
	// Currency pair
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
}

type GetAuthFullOrderBookRequest struct {
	// https://api.valr.com/v1/marketdata/:currencyPair/orderbook/full
	// Currency pair
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
}

type GetAuthTradeHistoryForPairRequest struct {
	// https://api.valr.com/v1/marketdata/:currencyPair/tradehistory?limit=10
	// Currency Pair
	// required: true
	// Limit
	// required: false
	Pair string `json:"currencyPair" url:"currencyPair"`
	Limit int   `json:"limit" param:"limit"`
}

type GetSimpleBuyOrSellOrderStatusRequest struct {
	// https://api.valr.com/v1/simple/:currencyPair/order/:orderId
	// Currency Pair
	// Order ID
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
	ID   string `json:"orderId" url:"orderId"`
}

type GetOrderStatusByOrderIDRequest struct {
	// https://api.valr.com/v1/orders/:currencyPair/orderid/:orderId
	// Currency Pair
	// Order ID
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
	ID   string `json:"orderId" url:"orderId"`
}

type GetOrderStatusByCustomerOrderIDRequest struct {
	// https://api.valr.com/v1/orders/:currencyPair/customerorderid/:customerOrderId
	// Currency Pair
	// Customer Order ID
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
	ID   string `json:"customerOrderId" url:"customerOrderId"`
}

type GetAllOpenOrdersRequest struct {
	// https://api.valr.com/v1/orders/open
	// Empty
}

type GetOrderHistoryRequest struct {
	// https://api.valr.com/v1/orders/history?skip=0&limit=2
	// Skip
	// Limit
	// required: false
	Skip  int    `json:"skip" param:"skip"`
	Limit int    `json:"limit" param:"limit"`
}

type GetOrderHistorySummaryByOrderIDRequest struct {
	 // https://api.valr.com/v1/orders/history/summary/orderid/:orderId
	 // Order ID
	 // required: true
	 ID string `json:"orderId" url:"orderId"`
}

type GetOrderHistorySummaryByCustomerOrderIDRequest struct {
	// https://api.valr.com/v1/orders/history/summary/customerorderid/:customerOrderId
	// Customer Order ID
	// required: true
	ID string `json:"customerOrderId" url:"customerOrderId"`
}

type GetOrderHistoryDetailsByOrderIDRequest struct {
	// https://api.valr.com/v1/orders/history/detail/orderid/:orderId
	// Order ID
	// required: true
	ID string `json:"orderId" url:"orderId"`
}

type GetOrderHistoryDetailsByCustomerOrderIDRequest struct {
	// https://api.valr.com/v1/orders/history/detail/customerorderid/:customerOrderId
	// Order ID
	// required: true
	ID string `json:"customerOrderId" url:"customerOrderId"`
}

/*
PRIVATE API POST REQUESTS
 */

type PostNewCryptoWithdrawRequest struct {
	// https://api.valr.com/v1/wallet/crypto/:currencyCode/withdraw
	// Currency Code
	// Amount
	// Address
	// required: true
	Asset string `json:"currencyCode" url:"currencyCode"`
	Amount decimal.Decimal `json:"amount" body:"amount"`
	Address string `json:"address" body:"address"`
}

type PostNewFiatWithdrawRequest struct {
	// https://api.valr.com/v1/wallet/fiat/:currencyCode/withdraw
	// Currency Code
	// Amount
	// BankAccountID
	// required: true
	Asset string `json:"currencyCode" url:"currencyCode"`
	Amount decimal.Decimal `json:"amount" body:"amount"`
	BankAccountID string `json:"linkedBankAccountId" body:"linkedBankAccountId"`
}

type PostSimpleBuyOrSellQuoteRequest struct {
	// https://api.valr.com/v1/simple/:currencyPair/quote
	// Currency Pair
	// Pay in currency
	// Pay amount
	// Side
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
	PayInCurrency string `json:"payInCurrency" body:"payInCurrency"`
	PayAmount decimal.Decimal `json:"payAmount" body:"payAmount"`
	Side Side `json:"side" body:"side"`
}

type PostSimpleBuyOrSellOrderRequest struct {
	// https://api.valr.com/v1/simple/:currencyPair/order
	// Currency Pair
	// Pay in currency
	// Pay amount
	// Side
	// required: true
	Pair string `json:"currencyPair" url:"currencyPair"`
	PayInCurrency string `json:"payInCurrency" body:"payInCurrency"`
	PayAmount decimal.Decimal `json:"payAmount" body:"payAmount"`
	Side Side `json:"side" body:"side"`
}

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
	Pair string `json:"pair" body:"pair"`
	Quantity decimal.Decimal `json:"quantity" body:"quantity"`
	Price decimal.Decimal `json:"price" body:"price"`
	Side Side `json:"side" body:"side"`
	PostOnly bool `json:"postOnly" body:"postOnly"`
	CustomerOrderID string `json:"customerOrderId" body:"customerOderId"`
}

type PostMarketOrderRequest struct {
	// https://api.valr.com/v1/orders/market
	// Currency Pair
	// Quantity
	// Side
	// required: true
	// Customer Order ID
	// required: false
	Pair string `json:"pair" body:"pair"`
	Quantity decimal.Decimal `json:"quoteAmount" body:"quoteAmount"`
	Side Side `json:"side" body:"side"`
	CustomerOrderID string `json:"customerOrderId" body:"customerOderId"`
}

/*
PRIVATE API DEL REQUESTS
 */

type DelOrderRequest struct {
	// https://api.valr.com/v1/orders/order
	// Currency Pair
	// Order ID
	// required: true
	Pair string `json:"pair" body:"pair"`
	ID string `json:"orderId" body:"orderId"`
}