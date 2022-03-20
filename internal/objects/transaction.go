package objects

type Transaction struct {
	Id int
	Ticker string
	TransactionDate int
	Price float64
	Shares int
	TransactionCategory int
	PurchaseValue float64
	MarketValue float64
}
