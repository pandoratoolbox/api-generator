package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_order = ReflectToFragment(models.OrderData{})
)

func NewOrder(ctx context.Context, data *models.Order) error {
	q := `
		mutation CreateOrder {
			order(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Order
	}{
		Data: *data,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		Order []models.Order
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Order) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Order[0].Id

	data.Id = &id

	return nil
}

func DeleteOrder(ctx context.Context, id int64) error {
	q := `
		mutation DeleteOrder {
			order(where: { id: { eq: $id } }) {
				id
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		Order []models.Order
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Order) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateOrder(ctx context.Context, data models.Order) error {
	q := `
		mutation UpdateOrder {
			order(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Order
	}{
		Id: *data.Id,
	}

	data.Id = nil
	input.Data = data

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		Order []models.Order
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Order) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetOrder(ctx context.Context, id int64) (models.Order, error) {
	var data models.Order

	q := fragment_order + `
			query GetOrder {
			order(where: { id: { eq: $id } }) {
				...Order
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return data, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		Order []models.Order
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Order) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Order[0]

	return data, nil
}

func ListOrderForUserByUserId(ctx context.Context, user_id int64) ([]models.Order, error) {
	var out []models.Order

	q := fragment_user + `query ListOrderForUser {
		order(where: { user_id: { eq: $user_id } }) {
			...User
		}
	}`

	input := struct {
		UserId int64 `json:"user_id"`
	}{
		UserId: user_id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	rt := struct {
		Order []models.Order `json:"order"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.Order

	return out, nil
}
