package transactionmethoddomain

type QRDetails struct {
	Provider    string `json:"provider"`
	QRString    string `json:"qr_string"`
	SessionID   string `json:"session_id"`
	ReferenceID string `json:"reference_id"`
}
