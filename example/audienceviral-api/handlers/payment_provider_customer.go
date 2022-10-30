package handlers

import (
	"audienceviral-api/models"
	"audienceviral-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetPaymentProviderCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "payment_provider_customer_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetPaymentProviderCustomer(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewPaymentProviderCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.PaymentProviderCustomer{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewPaymentProviderCustomer(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdatePaymentProviderCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.PaymentProviderCustomer{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdatePaymentProviderCustomer(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeletePaymentProviderCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "payment_provider_customer_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeletePaymentProviderCustomer(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListPaymentProviderCustomerForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListPaymentProviderCustomerForUserById(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
