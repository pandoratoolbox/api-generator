package handlers

import (
	"audienceviral-api/models"
	"audienceviral-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetUserFinance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "user_finance_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetUserFinance(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewUserFinance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.UserFinance{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewUserFinance(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateUserFinance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.UserFinance{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateUserFinance(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteUserFinance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "user_finance_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteUserFinance(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListUserFinanceForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListUserFinanceForUserById(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
