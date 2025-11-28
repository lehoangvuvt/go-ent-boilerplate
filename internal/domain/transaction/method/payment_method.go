package transactionmethoddomain

type PaymentMethodType string

const (
	MethodVisa    PaymentMethodType = "visa"
	MethodBanking PaymentMethodType = "banking"
	MethodEWallet PaymentMethodType = "ewallet"
	MethodQR      PaymentMethodType = "qr"
)
