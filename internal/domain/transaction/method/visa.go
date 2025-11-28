package transactionmethoddomain

type CardNetwork string

const (
	CardVisa       CardNetwork = "visa"
	CardMastercard CardNetwork = "mastercard"
	CardAmex       CardNetwork = "amex"
	CardJCB        CardNetwork = "jcb"
)

type VisaDetails struct {
	CardLast4         string      `json:"card_last_4"`
	CardNetwork       CardNetwork `json:"card_network"`
	AuthorizationCode string      `json:"authorization_code"`
	ReferenceID       string      `json:"reference_id"`
	Is3DSecure        bool        `json:"is_3d_secure"`
}
