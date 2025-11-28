package transactionmethoddomain

type EWalletDetails struct {
	Provider   string `json:"provider"`
	WalletID   string `json:"wallet_id"`
	RefID      string `json:"ref_id"`
	CustomerID string `json:"customer_id"`
}
