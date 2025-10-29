package methods

// GetStarTransactionsRequest represents the request for getStarTransactions method.
type GetStarTransactionsRequest struct {
	// Number of transactions to skip in the response
	Offset int `json:"offset,omitempty"`

	// The maximum number of transactions to be retrieved
	Limit int `json:"limit,omitempty"`
}

// GetStarTransactionsResponse represents the response for getStarTransactions method.
type GetStarTransactionsResponse struct {
	Transactions []interface{} `json:"transactions"`
}
