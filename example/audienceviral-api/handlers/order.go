package handlers

import (
	"audienceviral-api/models"
	"audienceviral-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "order_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetOrder(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.Order{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewOrder(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.Order{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateOrder(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "order_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteOrder(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListOrderForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListOrderForUserById(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
