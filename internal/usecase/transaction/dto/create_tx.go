package transactionusecasedto

import (
	"github.com/go-playground/validator/v10"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type CreateTransactionRequest struct {
	Amount   int64                                     `json:"amount" validate:"required,gt=0"`
	Currency string                                    `json:"currency" validate:"required,len=3"`
	Method   transactionmethoddomain.PaymentMethodType `json:"method" validate:"required,payment_method"`

	Visa    *transactionmethoddomain.VisaDetails    `json:"visa"`
	Banking *transactionmethoddomain.BankingDetails `json:"banking"`
	EWallet *transactionmethoddomain.EWalletDetails `json:"ewallet"`
	QRPay   *transactionmethoddomain.QRDetails      `json:"qr_pay"`
}

func init() {
	validate.RegisterValidation("payment_method", validatePaymentMethod)
	validate.RegisterStructValidation(transactionStructLevelValidation, CreateTransactionRequest{})
}

func (r *CreateTransactionRequest) Validate() error {
	return validate.Struct(r)
}

func validatePaymentMethod(fl validator.FieldLevel) bool {
	m := fl.Field().String()

	switch transactionmethoddomain.PaymentMethodType(m) {
	case transactionmethoddomain.MethodVisa,
		transactionmethoddomain.MethodBanking,
		transactionmethoddomain.MethodEWallet,
		transactionmethoddomain.MethodQR:
		return true
	}

	return false
}

func transactionStructLevelValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(CreateTransactionRequest)

	switch req.Method {
	case transactionmethoddomain.MethodVisa:
		if req.Visa == nil {
			sl.ReportError(req.Visa, "visa", "Visa", "visa_required", "")
		}
	case transactionmethoddomain.MethodBanking:
		if req.Banking == nil {
			sl.ReportError(req.Banking, "banking", "Banking", "banking_required", "")
		}
	case transactionmethoddomain.MethodEWallet:
		if req.EWallet == nil {
			sl.ReportError(req.EWallet, "ewallet", "EWallet", "ewallet_required", "")
		}
	case transactionmethoddomain.MethodQR:
		if req.QRPay == nil {
			sl.ReportError(req.QRPay, "qr_pay", "QRPay", "qr_required", "")
		}
	}
}
