package api

import (
	"time"

	"github.com/shopspring/decimal"
)

// OrderBookEntry is a single entry in an OrderBook
type OrderBookEntry struct {
	Price      decimal.Decimal `json:"price"`
	Quantity   decimal.Decimal `json:"quantity"`
	Side       string          `json:"side"`
	Pair       string          `json:"currencyPair"`
	OrderCount int             `json:"orderCount"`
}

// OrderBook holds OrderBookEntries
type OrderBook struct {
	Asks []OrderBookEntry `json:"Asks"`
	Bids []OrderBookEntry `json:"Bids"`
}

// CurrencyInfo holds info for a specific asset
type CurrencyInfo struct {
	Symbol    string `json:"symbol"`
	Active    bool   `json:"isActive"`
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}

// PairInfo holds info for a specific pair
type PairInfo struct {
	Symbol         string          `json:"symbol"`
	BaseCurrency   string          `json:"baseCurrency"`
	QuoteCurrency  string          `json:"quoteCurrency"`
	ShortName      string          `json:"shortName"`
	Active         bool            `json:"active"`
	MinBaseAmount  decimal.Decimal `json:"minBaseAmount"`
	MaxBaseAmount  decimal.Decimal `json:"maxBaseAmount"`
	MinQuoteAmount decimal.Decimal `json:"minQuoteAmount"`
	MaxQuoteAmount decimal.Decimal `json:"maxQuoteAmount"`
}

// OrderTypes associates order types with a specific pair
type OrderTypes struct {
	Pair       string   `json:"currencyPair"`
	OrderTypes []string `json:"orderTypes"`
}

// MarketSummary holds market summary information
type MarketSummary struct {
	Pair               string          `json:"currencyPair"`
	AskPrice           decimal.Decimal `json:"askPrice"`
	BidPrice           decimal.Decimal `json:"bidPrice"`
	LastPrice          decimal.Decimal `json:"lastTradedPrice"`
	ClosePrice         decimal.Decimal `json:"previousClosePrice"`
	BaseVolume         decimal.Decimal `json:"baseVolume"`
	HighPrice          decimal.Decimal `json:"highPrice"`
	LowPrice           decimal.Decimal `json:"lowPrice"`
	Created            time.Time       `json:"created"`
	ChangeFromPrevious string          `json:"changeFromPrevious"`
}

// AccountBalance represent the balance info for a specific asset
type AccountBalance struct {
	Currency  string          `json:"currency"`
	Available decimal.Decimal `json:"available"`
	Reserved  decimal.Decimal `json:"reserved"`
	Total     decimal.Decimal `json:"total"`
}

// TransactionInfo holds transaction info
type TransactionInfo struct {
	TransactionType TransactionType           `json:"transactionType"`
	DebitCurrency   string                    `json:"debitCurrency"`
	DebitValue      decimal.Decimal           `json:"debitValue"`
	CreditCurrency  string                    `json:"creditCurrency"`
	CreditValue     decimal.Decimal           `json:"creditValue"`
	FeeCurrency     string                    `json:"feeCurrency"`
	FeeValue        decimal.Decimal           `json:"feeValue"`
	EventAt         time.Time                 `json:"eventAt"`
	AdditionalInfo  AdditionalTransactionInfo `json:"additionalInfo"`
}

// TransactionType associates a transction type with its description
type TransactionType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// AdditionalTransactionInfo holds additional info for a transaction
type AdditionalTransactionInfo struct {
	CostPerCoin        decimal.Decimal `json:"costPerCoin"`
	CostPerCoinSymbol  string          `json:"costPerCoinSymbol"`
	CurrencyPairSymbol string          `json:"currencyPairSymbol"`
	OrderID            string          `json:"orderId"`
}

// TradeInfo holds info about a specific trade
type TradeInfo struct {
	Price    decimal.Decimal `json:"price"`
	Quantity decimal.Decimal `json:"quantity"`
	Pair     string          `json:"currencyPair"`
	TradedAt time.Ticker     `json:"tradedAt"`
	Side     ResponseSide    `json:"side"`
	TradeID  int             `json:"tradeId"`
}

// DepositInfo holds info about a specific deposit
type DepositInfo struct {
	CurrencyCode    string          `json:"currencyCode"`
	ReceiveAddress  string          `json:"receiveAddress"`
	TransactionHash string          `json:"transactionHash"`
	Amount          decimal.Decimal `json:"amount"`
	CreatedAt       time.Time       `json:"createdAt"`
	Confirmations   int             `json:"confirmations"`
	Confirmed       bool            `json:"confirmed"`
	ConfirmedAt     time.Time       `json:"confirmedAt"`
}

// WithdrawInfo holds info about a specific withdraw
type WithdrawInfo struct {
	Currency           string          `json:"currency"`
	Address            string          `json:"address"`
	Amount             decimal.Decimal `json:"amount"`
	FeeAmount          decimal.Decimal `json:"feeAmount"`
	TransactionHash    string          `json:"transactionHash"`
	Confirmations      int             `json:"confirmations"`
	LastConfirmationAt time.Time       `json:"lastConfirmationAt"`
	UniqueID           string          `json:"uniqueId"`
	CreateAt           time.Time       `json:"createdAt"`
	Verified           bool            `json:"verified"`
	State              string          `json:"status"`
}

// BankInfo holds data related to a bank account
type BankInfo struct {
	ID            string    `json:"id"`
	Bank          string    `json:"bank"`
	AccountHolder string    `json:"accountHolder"`
	AccountNumber string    `json:"accountNumber"`
	BranchCode    string    `json:"branchCode"`
	AccountType   string    `json:"accountType"`
	CreatedAt     time.Time `json:"createdAt"`
}

// TradeHistoryInfo is the data for a specific trade in trade history
type TradeHistoryInfo struct {
	Price      decimal.Decimal `json:"price"`
	Quantity   decimal.Decimal `json:"quantity"`
	Pair       string          `json:"currencyPair"`
	TradedAt   time.Ticker     `json:"tradedAt"`
	TakerSide  ResponseSide    `json:"takerSide"`
	SequenceID int             `json:"sequenceId"`
	ID         string          `json:"id"`
}

// OpenOrder holds info for an open order on the market
type OpenOrder struct {
	OrderID           string          `json:"orderId"`
	Side              ResponseSide    `json:"side"`
	Price             decimal.Decimal `json:"price"`
	Pair              string          `json:"currencyPair"`
	CreatedAt         time.Time       `json:"createdAt"`
	RemainingQuantity decimal.Decimal `json:"remainingQuantity"`
	OriginalQuantity  decimal.Decimal `json:"originalQuantity"`
	FilledPercentage  decimal.Decimal `json:"filledPercentage"`
	CustomerOrderID   string          `json:"customerOrderId"`
}

// OrderReceipt collects info for a successful order
type OrderReceipt struct {
	OrderID          string          `json:"orderId"`
	CustomerOrderID  string          `json:"customerOrderId"`
	OrderStatusType  bool            `json:"orderStatusType"`
	Pair             string          `json:"currencyPair"`
	AveragePrice     decimal.Decimal `json:"averagePrice"`
	OriginalPrice    decimal.Decimal `json:"originalPrice"`
	Quantity         decimal.Decimal `json:"quantity"`
	OriginalQuantity decimal.Decimal `json:"originalQuantity"`
	Total            decimal.Decimal `json:"total"`
	TotalFee         decimal.Decimal `json:"totalFee"`
	OrderSide        ResponseSide    `json:"orderSide"`
	OrderType        string          `json:"orderType"`
	FailedReason     string          `json:"failedReason"`
	OrderUpdatedAt   time.Time       `json:"orderUpdatedAt"`
	OrderCreatedAt   time.Time       `json:"orderCreatedAt"`
}

// OrderStatus holds info related to the status of a specific order
type OrderStatus struct {
	OrderID           string          `json:"orderId"`
	OrderStatusType   string          `json:"orderStatusType"`
	Pair              string          `json:"currencyPair"`
	OriginalPrice     decimal.Decimal `json:"originalPrice"`
	RemainingQuantity decimal.Decimal `json:"remainingQuantity"`
	OriginalQuantity  decimal.Decimal `json:"originalQuantity"`
	OrderSide         ResponseSide    `json:"orderSide"`
	OrderType         string          `json:"orderType"`
	FailedReason      string          `json:"failedReason"`
	OrderUpdatedAt    time.Time       `json:"orderUpdatedAt"`
	OrderCreatedAt    time.Time       `json:"orderCreatedAt"`
	CustomerOrderID   string          `json:"customerOrderId"`
}

// RequestSide type for explicitly representing the two options
type RequestSide string

const (
	BUY  RequestSide = "BUY"
	SELL RequestSide = "SELL"
)

// ResponseSide type for explicitly representing the two options
type ResponseSide string

const (
	Buy  ResponseSide = "buy"
	Sell ResponseSide = "sell"
)
