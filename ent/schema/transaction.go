package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),

		field.Int64("amount"),

		field.String("currency"),

		field.UUID("user_id", uuid.UUID{}),

		field.Enum("status").
			Values(
				string(transactiondomain.Pending),
				string(transactiondomain.Completed),
				string(transactiondomain.Failed),
				string(transactiondomain.Rejected),
			).
			Default(string(transactiondomain.Pending)),

		field.Enum("method").
			Values(
				string(transactionmethoddomain.MethodVisa),
				string(transactionmethoddomain.MethodBanking),
				string(transactionmethoddomain.MethodEWallet),
				string(transactionmethoddomain.MethodQR),
			),

		field.JSON("visa_details", &transactionmethoddomain.VisaDetails{}).
			Optional(),

		field.JSON("banking_details", &transactionmethoddomain.BankingDetails{}).
			Optional(),

		field.JSON("ewallet_details", &transactionmethoddomain.EWalletDetails{}).
			Optional(),

		field.JSON("qr_details", &transactionmethoddomain.QRDetails{}).
			Optional(),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("transactions").
			Field("user_id").
			Required().
			Unique(),
	}
}
