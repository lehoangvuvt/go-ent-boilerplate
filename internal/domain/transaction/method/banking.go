package transactionmethoddomain

type BankingDetails struct {
	BankCode        string `json:"bank_code"`
	BankName        string `json:"bank_name"`
	AccountNumber   string `json:"account_number"`
	ReferenceID     string `json:"reference_id"`
	TransactionTime string `json:"transaction_time"`
}
