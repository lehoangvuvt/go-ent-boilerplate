package transactionrepository

import (
	"encoding/json"

	"github.com/lehoangvuvt/go-ent-boilerplate/ent"
	"github.com/lehoangvuvt/go-ent-boilerplate/ent/transaction"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
)

func toDomain(et *ent.Transaction) (*transactiondomain.Transaction, error) {
	dt := &transactiondomain.Transaction{
		ID:        et.ID,
		Amount:    et.Amount,
		Currency:  et.Currency,
		UserID:    et.UserID,
		Method:    transactionmethoddomain.PaymentMethodType(et.Method),
		Status:    transactiondomain.TransactionStatus(et.Status),
		CreatedAt: et.CreatedAt,
		UpdatedAt: et.UpdatedAt,
	}

	if et.VisaDetails != nil {
		var visa transactionmethoddomain.VisaDetails
		_ = json.Unmarshal(et.VisaDetails, &visa)
		dt.Visa = &visa
	}

	if et.BankingDetails != nil {
		var bank transactionmethoddomain.BankingDetails
		_ = json.Unmarshal(et.BankingDetails, &bank)
		dt.Banking = &bank
	}

	if et.EwalletDetails != nil {
		var ew transactionmethoddomain.EWalletDetails
		_ = json.Unmarshal(et.EwalletDetails, &ew)
		dt.EWallet = &ew
	}

	if et.QrDetails != nil {
		var qr transactionmethoddomain.QRDetails
		_ = json.Unmarshal(et.QrDetails, &qr)
		dt.QRPay = &qr
	}

	return dt, nil
}

func toDomainList(list []*ent.Transaction) ([]*transactiondomain.Transaction, error) {
	out := make([]*transactiondomain.Transaction, 0, len(list))
	for _, et := range list {
		dt, err := toDomain(et)
		if err != nil {
			return nil, err
		}
		out = append(out, dt)
	}
	return out, nil
}

func applyDomainToCreate(b *ent.TransactionCreate, dt *transactiondomain.Transaction) {
	b.
		SetID(dt.ID).
		SetAmount(dt.Amount).
		SetCurrency(dt.Currency).
		SetUserID(dt.UserID).
		SetMethod(transaction.Method(dt.Method)).
		SetStatus(transaction.Status(dt.Status)).
		SetCreatedAt(dt.CreatedAt).
		SetUpdatedAt(dt.UpdatedAt)

	if dt.Visa != nil {
		raw, _ := json.Marshal(dt.Visa)
		b.SetVisaDetails(raw)
	}

	if dt.Banking != nil {
		raw, _ := json.Marshal(dt.Banking)
		b.SetBankingDetails(raw)
	}

	if dt.EWallet != nil {
		raw, _ := json.Marshal(dt.EWallet)
		b.SetEwalletDetails(raw)
	}

	if dt.QRPay != nil {
		raw, _ := json.Marshal(dt.QRPay)
		b.SetQrDetails(raw)
	}
}

func applyDomainToUpdate(b *ent.TransactionUpdateOne, dt *transactiondomain.Transaction) {
	b.
		SetAmount(dt.Amount).
		SetCurrency(dt.Currency).
		SetUserID(dt.UserID).
		SetMethod(transaction.Method(dt.Method)).
		SetStatus(transaction.Status(dt.Status)).
		SetUpdatedAt(dt.UpdatedAt)

	if dt.Visa != nil {
		raw, _ := json.Marshal(dt.Visa)
		b.SetVisaDetails(raw)
	} else {
		b.ClearVisaDetails()
	}

	if dt.Banking != nil {
		raw, _ := json.Marshal(dt.Banking)
		b.SetBankingDetails(raw)
	} else {
		b.ClearBankingDetails()
	}

	if dt.EWallet != nil {
		raw, _ := json.Marshal(dt.EWallet)
		b.SetEwalletDetails(raw)
	} else {
		b.ClearEwalletDetails()
	}

	if dt.QRPay != nil {
		raw, _ := json.Marshal(dt.QRPay)
		b.SetQrDetails(raw)
	} else {
		b.ClearQrDetails()
	}
}
