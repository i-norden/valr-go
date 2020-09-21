package api

import (
	"context"
	"net/http"
)

/*
PUBLIC API GET REQUESTS
*/

// GetOrderBook
//
// Returns a list of the top 20 bids and asks in the order book.
// Ask orders are sorted by price ascending. Bid orders are sorted by price descending.
// Orders of the same price are aggregated.
func (cl *Client) GetOrderBook(ctx context.Context, req *GetOrderBookRequest) (*GetOrderBookResponse, error) {
	var res GetOrderBookResponse
	err := cl.do(ctx, http.MethodGet, "/public/{currencyPair}/orderbook", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCurrencies
//
// Get a list of currencies supported by VALR.
func (cl *Client) GetCurrencies(ctx context.Context, req *GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	var res GetCurrenciesResponse
	err := cl.do(ctx, http.MethodGet, "/public/currencies", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCurrencyPairs
//
// Get a list of all the currency pairs supported by VALR.
func (cl *Client) GetCurrencyPairs(ctx context.Context, req *GetCurrencyPairsRequest) (*GetCurrencyPairsResponse, error) {
	var res GetCurrencyPairsResponse
	err := cl.do(ctx, http.MethodGet, "/public/pairs", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderTypesRequest
//
// Get all the order types supported for all currency pairs.
// An array of currency pairs is returned along with an array of order types for each currency pair.
// You can only place an order that is supported by that currency pair.
//
// The order types supported are as follows:
// limit post-only : Place a limit order on the Exchange that will either be added to the order book or, should it match, be cancelled completely.
// limit : Place a limit order on the Exchange.
// market : Place a market order on the Exchange (only crypto-to-ZAR pairs).
// simple : Similar to a market order, but allows for crypto-to-crypto pairs.
func (cl *Client) GetOrderTypesRequest(ctx context.Context, req *GetOrderTypesRequest) (*GetOrderTypesResponse, error) {
	var res GetOrderTypesResponse
	err := cl.do(ctx, http.MethodGet, "/public/ordertypes", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderTypesForPairRequest
//
// Get the order types supported for a given currency pair.
// An array of order types is returned.
// You can only place an order that is listed in this array for this currency pair.
//
// The order types supported are as follows:
// limit post-only : Place a limit order on the Exchange that will either be added to the order book or, should it match, be cancelled completely.
// limit : Place a limit order on the Exchange.
// market : Place a market order on the Exchange (only crypto-to-ZAR pairs).
// simple : Similar to a market order, but allows for crypto-to-crypto pairs.
func (cl *Client) GetOrderTypesForPairRequest(ctx context.Context, req *GetOrderTypesForPairRequest) (*GetOrderTypesForPairResponse, error) {
	var res GetOrderTypesForPairResponse
	err := cl.do(ctx, http.MethodGet, "/public/{currencyPair}/ordertypes", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetMarketSummaryRequest
//
// Get the market summary for all supported currency pairs.
func (cl *Client) GetMarketSummaryRequest(ctx context.Context, req *GetMarketSummaryRequest) (*GetMarketSummaryResponse, error) {
	var res GetMarketSummaryResponse
	err := cl.do(ctx, http.MethodGet, "/public/marketsummary", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetMarketSummaryForPairRequest
//
// Get the market summary for a given currency pair.
func (cl *Client) GetMarketSummaryForPairRequest(ctx context.Context, req *GetMarketSummaryForPairRequest) (*GetMarketSummaryForPairResponse, error) {
	var res GetMarketSummaryForPairResponse
	err := cl.do(ctx, http.MethodGet, "/public/{currencyPair}/marketsummary", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetServerTimeRequest
//
// Get the server time.
func (cl *Client) GetServerTimeRequest(ctx context.Context, req *GetServerTimeRequest) (*GetServerTimeResponse, error) {
	var res GetServerTimeResponse
	err := cl.do(ctx, http.MethodGet, "/public/time", req, &res, false)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

/*
PRIVATE API GET REQUESTS
*/

// GetAccountBalancesRequest
//
// Returns the list of all wallets with their respective balances.
func (cl *Client) GetAccountBalancesRequest(ctx context.Context, req *GetAccountBalancesRequest) (*GetAccountBalancesResponse, error) {
	var res GetAccountBalancesResponse
	err := cl.do(ctx, http.MethodGet, "/account/balances", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetTransactionHistoryRequest
//
// Transaction history for your account. Note: This API supports pagination.
func (cl *Client) GetTransactionHistoryRequest(ctx context.Context, req *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error) {
	var res GetTransactionHistoryResponse
	err := cl.do(ctx, http.MethodGet, "/account/transactionhistory", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetTradeHistoryForPairRequest
//
// Get the last 100 recent trades for a given currency pair for your account.
// You can limit the number of trades returned by specifying the limit parameter.
func (cl *Client) GetTradeHistoryForPairRequest(ctx context.Context, req *GetTradeHistoryForPairRequest) (*GetTradeHistoryForPairResponse, error) {
	var res GetTradeHistoryForPairResponse
	err := cl.do(ctx, http.MethodGet, "/account/{currencyPair}/tradehistory", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetDepositAddressRequest
//
// Returns the default deposit address associated with currency specified in the path variable {currencyCode}.
func (cl *Client) GetDepositAddressRequest(ctx context.Context, req *GetDepositAddressRequest) (*GetDepositAddressResponse, error) {
	var res GetDepositAddressResponse
	err := cl.do(ctx, http.MethodGet, "/wallet/crypto/{currencyCode}/deposit/address", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetWithdrawInfoRequest
//
// Get all the information about withdrawing a given currency from your VALR account.
// That will include withdrawal costs, minimum withdrawal amount etc.
func (cl *Client) GetWithdrawInfoRequest(ctx context.Context, req *GetWithdrawInfoRequest) (*GetWithdrawInfoResponse, error) {
	var res GetWithdrawInfoResponse
	err := cl.do(ctx, http.MethodGet, "/wallet/crypto/{currencyCode}/withdraw", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetWithdrawStatusRequest
//
// Check the status of a withdrawal.
func (cl *Client) GetWithdrawStatusRequest(ctx context.Context, req *GetWithdrawStatusRequest) (*GetWithdrawStatusResponse, error) {
	var res GetWithdrawStatusResponse
	err := cl.do(ctx, http.MethodGet, "/wallet/crypto/{currencyCode}/withdraw/{withdrawId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetDepositHistoryForAssetRequest
//
// Get the Deposit History records for a given currency.
func (cl *Client) GetDepositHistoryForAssetRequest(ctx context.Context, req *GetDepositHistoryForAssetRequest) (*GetDepositHistoryForAssetResponse, error) {
	var res GetDepositHistoryForAssetResponse
	err := cl.do(ctx, http.MethodGet, "/wallet/crypto/{currencyCode}/deposit/history", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetWithdrawHistoryForAssetRequest
//
// Get Withdrawal History records for a given currency.
func (cl *Client) GetWithdrawHistoryForAssetRequest(ctx context.Context, req *GetWithdrawHistoryForAssetRequest) (*GetWithdrawHistoryForAssetResponse, error) {
	var res GetWithdrawHistoryForAssetResponse
	err := cl.do(ctx, http.MethodGet, "/wallet/crypto/{currencyCode}/withdraw/history", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetBankAccountForAssetRequest
// Get a list of bank accounts that are linked to your VALR account.
// Bank accounts can be linked by signing in to your account on www.VALR.com.
func (cl *Client) GetBankAccountForAssetRequest(ctx context.Context, req *GetBankAccountForAssetRequest) (*GetBankAccountForAssetResponse, error) {
	var res GetBankAccountForAssetResponse
	err := cl.do(ctx, http.MethodGet, "/wallet/fiat/{currencyCode}/accounts", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAuthOrderBookRequest
//
// Returns a list of the top 20 bids and asks in the order book.
// Ask orders are sorted by price ascending. Bid orders are sorted by price descending.
// Orders of the same price are aggregated.
func (cl *Client) GetAuthOrderBookRequest(ctx context.Context, req *GetAuthOrderBookRequest) (*GetAuthOrderBookResponse, error) {
	var res GetAuthOrderBookResponse
	err := cl.do(ctx, http.MethodGet, "/marketdata/{currencyPair}/orderbook", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAuthFullOrderBookRequest
//
// Returns a list of the top 20 bids and asks in the order book.
// Ask orders are sorted by price ascending.
// Bid orders are sorted by price descending.
// Orders of the same price are aggregated.
func (cl *Client) GetAuthFullOrderBookRequest(ctx context.Context, req *GetAuthFullOrderBookRequest) (*GetAuthFullOrderBookResponse, error) {
	var res GetAuthFullOrderBookResponse
	err := cl.do(ctx, http.MethodGet, "/marketdata/{currencyPair}/orderbook/full", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAuthTradeHistoryForPairRequest
//
// Get the last 100 recent trades for a given currency pair.
// You can limit the number of trades returned by specifying the limit parameter.
func (cl *Client) GetAuthTradeHistoryForPairRequest(ctx context.Context, req *GetAuthTradeHistoryForPairRequest) (*GetAuthTradeHistoryForPairResponse, error) {
	var res GetAuthTradeHistoryForPairResponse
	err := cl.do(ctx, http.MethodGet, "/marketdata/:currencyPair/tradehistory", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetSimpleBuyOrSellOrderStatusRequest
//
// Get the status of a Simple Buy/Sell order.
func (cl *Client) GetSimpleBuyOrSellOrderStatusRequest(ctx context.Context, req *GetSimpleBuyOrSellOrderStatusRequest) (*GetSimpleBuyOrSellOrderStatusResponse, error) {
	var res GetSimpleBuyOrSellOrderStatusResponse
	err := cl.do(ctx, http.MethodGet, "/simple/{currencyPair}/order/{orderId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderStatusByOrderIDRequest
//
// This API returns the status of an order that was placed on the Exchange queried using the id provided by VALR.
// VALR provides an id for every order that is placed on the Exchange.
// Use this id to populate the path variable orderId in this API to query the status of the order.
//
// Note: If a customerOrderId was also specified while placing the order, that customerOrderId will be returned as part of the response.
func (cl *Client) GetOrderStatusByOrderIDRequest(ctx context.Context, req *GetOrderStatusByOrderIDRequest) (*GetOrderStatusByOrderIDResponse, error) {
	var res GetOrderStatusByOrderIDResponse
	err := cl.do(ctx, http.MethodGet, "/orders/{currencyPair}/orderid/{orderId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderStatusByCustomerOrderIDRequest
//
// This API returns the status of an order that was placed on the Exchange queried using customerOrderId.
// The customer can specify a customerOrderId while placing an order on the Exchange.
// Use this API to query the order status using that customerOrderId.
func (cl *Client) GetOrderStatusByCustomerOrderIDRequest(ctx context.Context, req *GetOrderStatusByCustomerOrderIDRequest) (*GetOrderStatusByCustomerOrderIDResponse, error) {
	var res GetOrderStatusByCustomerOrderIDResponse
	err := cl.do(ctx, http.MethodGet, "/orders/{currencyPair}/customerorderid/{customerOrderId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAllOpenOrdersRequest
//
// Get all open orders for your account.
// A customerOrderId field will be returned in the response for all those orders that were created with a customerOrderId field.
func (cl *Client) GetAllOpenOrdersRequest(ctx context.Context, req *GetAllOpenOrdersRequest) (*GetAllOpenOrdersResponse, error) {
	var res GetAllOpenOrdersResponse
	err := cl.do(ctx, http.MethodGet, "/orders/open", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderHistoryRequest
//
// Get historical orders placed by you.
func (cl *Client) GetOrderHistoryRequest(ctx context.Context, req *GetOrderHistoryRequest) (*GetOrderHistoryResponse, error) {
	var res GetOrderHistoryResponse
	err := cl.do(ctx, http.MethodGet, "/orders/history", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderHistorySummaryByOrderIDRequest
//
// An order is considered completed when the "Order Status" call returns one of the following statuses: "Filled", "Cancelled" or "Failed".
// When this happens, you can get a more detailed summary about this order using this call.
// Orders that are not completed are invalid for this request.
func (cl *Client) GetOrderHistorySummaryByOrderIDRequest(ctx context.Context, req *GetOrderHistorySummaryByOrderIDRequest) (*GetOrderHistorySummaryByOrderIDResponse, error) {
	var res GetOrderHistorySummaryByOrderIDResponse
	err := cl.do(ctx, http.MethodGet, "/orders/history/summary/orderid/{orderId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderHistorySummaryByCustomerOrderIDRequest
//
// An order is considered completed when the "Order Status" call returns one of the following statuses: "Filled", "Cancelled" or "Failed".
// When this happens, you can get a more detailed summary about this order using this call.
// Orders that are not completed are invalid for this request.
func (cl *Client) GetOrderHistorySummaryByCustomerOrderIDRequest(ctx context.Context, req *GetOrderHistorySummaryByCustomerOrderIDRequest) (*GetOrderHistorySummaryByCustomerOrderIDResponse, error) {
	var res GetOrderHistorySummaryByCustomerOrderIDResponse
	err := cl.do(ctx, http.MethodGet, "/orders/history/summary/customerorderid/{customerOrderId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderHistoryDetailsByOrderIDRequest
//
// Get a detailed history of an order's statuses. This call returns an array of "Order Status" objects.
// The latest and most up-to-date status of this order is the zeroth element in the array.
func (cl *Client) GetOrderHistoryDetailsByOrderIDRequest(ctx context.Context, req *GetOrderHistoryDetailsByOrderIDRequest) (*GetOrderHistoryDetailsByOrderIDResponse, error) {
	var res GetOrderHistoryDetailsByOrderIDResponse
	err := cl.do(ctx, http.MethodGet, "/orders/history/detail/orderid/{orderId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOrderHistoryDetailsByCustomerOrderIDRequest
//
// Get a detailed history of an order's statuses. This call returns an array of "Order Status" objects.
// The latest and most up-to-date status of this order is the zeroth element in the array.
func (cl *Client) GetOrderHistoryDetailsByCustomerOrderIDRequest(ctx context.Context, req *GetOrderHistoryDetailsByCustomerOrderIDRequest) (*GetOrderHistoryDetailsByCustomerOrderIDResponse, error) {
	var res GetOrderHistoryDetailsByCustomerOrderIDResponse
	err := cl.do(ctx, http.MethodGet, "/orders/history/detail/customerorderid/{customerOrderId}", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

/*
PRIVATE API POST REQUESTS
*/

// PostNewCryptoWithdrawRequest
//
// Withdraw cryptocurrency funds to an address.
// The request body for XRP, XMR, XEM, XLM will accept an optional field called "paymentReference".
// Max length for paymentReference is 256.
func (cl *Client) PostNewCryptoWithdrawRequest(ctx context.Context, req *PostNewCryptoWithdrawRequest) (*PostNewCryptoWithdrawResponse, error) {
	var res PostNewCryptoWithdrawResponse
	err := cl.do(ctx, http.MethodPost, "/wallet/crypto/{currencyCode}/withdraw", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// PostNewFiatWithdrawRequest
//
// Withdraw your ZAR funds into one of your linked bank accounts.
func (cl *Client) PostNewFiatWithdrawRequest(ctx context.Context, req *PostNewFiatWithdrawRequest) (*PostNewFiatWithdrawResponse, error) {
	var res PostNewFiatWithdrawResponse
	err := cl.do(ctx, http.MethodPost, "/wallet/fiat/{currencyCode}/withdraw", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// PostSimpleBuyOrSellQuoteRequest
//
// Get a quote to buy or sell instantly using Simple Buy.
//
// A sample request body is as follows:
//
//{
//    "payInCurrency": "BTC",
//    "payAmount": "0.001",
//    "side": "SELL"
//}
func (cl *Client) PostSimpleBuyOrSellQuoteRequest(ctx context.Context, req *PostSimpleBuyOrSellQuoteRequest) (*PostSimpleBuyOrSellQuoteResponse, error) {
	var res PostSimpleBuyOrSellQuoteResponse
	err := cl.do(ctx, http.MethodPost, "/simple/{currencyPair}/quote", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// PostSimpleBuyOrSellOrderRequest
//
// Submit an order to buy or sell instantly using Simple Buy/Sell.
//
// A sample request body is as follows:
//
//{
//    "payInCurrency": "BTC",
//    "payAmount": "0.001",
//    "side": "SELL"
//}
func (cl *Client) PostSimpleBuyOrSellOrderRequest(ctx context.Context, req *PostSimpleBuyOrSellOrderRequest) (*PostSimpleBuyOrSellOrderResponse, error) {
	var res PostSimpleBuyOrSellOrderResponse
	err := cl.do(ctx, http.MethodPost, "/simple/{currencyPair}/order", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// PostLimitOrderRequest
//
// Create a new limit order.
//
// The JSON body used to create a limit order looks like this:
//
//{
//    "side": "SELL",
//    "quantity": "0.100000",
//    "price": "10000",
//    "pair": "BTCZAR",
//    "postOnly": true,
//    "customerOrderId": "1234"
//}
func (cl *Client) PostLimitOrderRequest(ctx context.Context, req *PostLimitOrderRequest) (*PostLimitOrderResponse, error) {
	var res PostLimitOrderResponse
	err := cl.do(ctx, http.MethodPost, "/orders/limit", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// PostMarketBuyRequest
//
// Create a new market order.
// When the response is 202 Accepted, you can either use the Order Status REST API or use WebSocket API to receive updates about this order.
//
// Example request body:
//
//{
//    "side": "BUY",
//    "quoteAmount": "0.100000",
//    "pair": "BTCZAR",
//    "customerOrderId": "1234"
//}
func (cl *Client) PostMarketBuyRequest(ctx context.Context, req *PostMarketOrderBuyRequest) (*PostMarketOrderResponse, error) {
	var res PostMarketOrderResponse
	err := cl.do(ctx, http.MethodPost, "/orders/market", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// PostMarketBuyRequest
//
// Create a new market order.
// When the response is 202 Accepted, you can either use the Order Status REST API or use WebSocket API to receive updates about this order.
//
// Example request body:
//
//{
//    "side": "SELL",
//    "baseAmount": "0.100000",
//    "pair": "BTCZAR",
//    "customerOrderId": "1234"
//}
func (cl *Client) PostMarketSellRequest(ctx context.Context, req *PostMarketOrderSellRequest) (*PostMarketOrderResponse, error) {
	var res PostMarketOrderResponse
	err := cl.do(ctx, http.MethodPost, "/orders/market", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

/*
PRIVATE API DEL REQUESTS
*/

// DelOrderRequest
//
// Cancel an open order.
// A 202 Accepted response means the request to cancel the order was accepted.
// You can either use the Order Status REST API or use WebSocket API to receive status update about this request.
//
// The DELETE request requires a JSON request body in the following format:
//
//{
//  "orderId": "UUID",
//  "pair": "BTCZAR"
//}
func (cl *Client) DelOrderRequest(ctx context.Context, req *DelOrderRequest) (*DelOrderResponse, error) {
	var res DelOrderResponse
	err := cl.do(ctx, http.MethodDelete, "/orders/order", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// DelOrderByCustomerOrderIDRequest
//
// Cancel an open order.
// A 202 Accepted response means the request to cancel the order was accepted.
// You can either use the Order Status REST API or use WebSocket API to receive status update about this request.
//
// The DELETE request requires a JSON request body in the following format:
//
//{
//  "customerOrderId": "^[0-9a-zA-Z-]{0,50}$",
//  "pair": "BTCZAR"
//}
func (cl *Client) DelOrderByCustomerOrderIDRequest(ctx context.Context, req *DelOrderByCustomerOrderIDRequest) (*DelOrderByCustomerOrderIDResponse, error) {
	var res DelOrderByCustomerOrderIDResponse
	err := cl.do(ctx, http.MethodDelete, "/orders/order", req, &res, true)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
