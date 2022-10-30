package handlers

import (
	"audienceviral-api/models"
	"audienceviral-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "invoice_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetInvoice(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.Invoice{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewInvoice(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.Invoice{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateInvoice(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "invoice_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteInvoice(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListInvoiceForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListInvoiceForUserById(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
