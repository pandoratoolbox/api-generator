package handlers

import (
	"audienceviral-api/models"
	"audienceviral-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "subscription_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetSubscription(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.Subscription{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewSubscription(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.Subscription{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateSubscription(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "subscription_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteSubscription(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListSubscriptionForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListSubscriptionForUserById(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
