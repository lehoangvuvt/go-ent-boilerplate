package httptransaction

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *TransactionHandler) {
	r.Route("/transactions", func(tr chi.Router) {
		tr.Post("/", h.Create)
		tr.Get("/", h.List)
		tr.Route("/{id}", func(ir chi.Router) {
			ir.Get("/", h.FindByID)
			ir.Post("/confirm", h.Confirm)
			ir.Post("/cancel", h.Cancel)
		})
	})
}
