package transactions

//Interface is interface to describe how to handle transactions
type Interface interface {
	CreateTransaction(transation Transaction) error
	GetTransationsForSellerID(sellerID string) ([]Transaction, error)
}

// Transaction should be representation of model in database
type Transaction struct {
	uid        string
	sellerID   string
	buyerID    string
	serviceFee int32
	payment    int32
}
