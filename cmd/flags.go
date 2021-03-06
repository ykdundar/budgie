package cmd

// flags.
var (
	token     string  // token represents an API token obtained from marketstack.
	name      string  // name represents portfolio and stock names.
	rename    string  // rename represents the new name is to be given for a portfolio.
	portfolio string  // portfolio represents portfolio name belongs to a selected stock.
	ticker    string  // ticker represents stock symbols.
	price     float64 // price represents market value of selected stock.
	shares    int     // shares represents number of shares traded.
	date      string  // date represents the day of transaction.
	id        int     // id represents the order of stocks, portfolios and transactions.
)
