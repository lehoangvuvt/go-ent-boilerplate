package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
			Values("pending", "completed", "failed", "rejected").
			Default("pending"),

		field.Enum("method").
			Values("visa", "banking", "ewallet", "qr"),

		field.JSON("visa_details", []byte{}).Optional(),
		field.JSON("banking_details", []byte{}).Optional(),
		field.JSON("ewallet_details", []byte{}).Optional(),
		field.JSON("qr_details", []byte{}).Optional(),

		field.Time("created_at").Default(time.Now),
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
			Unique().
			Required(),
	}
}
