package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_order_credits = ReflectToFragment(models.OrderCreditsData{})
)

func NewOrderCredits(ctx context.Context, data *models.OrderCredits) error {
	q := `
		mutation CreateOrderCredits {
			order_credits(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.OrderCredits
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
		OrderCredits []models.OrderCredits
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.OrderCredits) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.OrderCredits[0].Id

	data.Id = &id

	return nil
}

func DeleteOrderCredits(ctx context.Context, id int64) error {
	q := `
		mutation DeleteOrderCredits {
			order_credits(where: { id: { eq: $id } }) {
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
		OrderCredits []models.OrderCredits
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.OrderCredits) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateOrderCredits(ctx context.Context, data models.OrderCredits) error {
	q := `
		mutation UpdateOrderCredits {
			order_credits(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.OrderCredits
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
		OrderCredits []models.OrderCredits
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.OrderCredits) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetOrderCredits(ctx context.Context, id int64) (models.OrderCredits, error) {
	var data models.OrderCredits

	q := fragment_order_credits + `
			query GetOrderCredits {
			order_credits(where: { id: { eq: $id } }) {
				...OrderCredits
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
		OrderCredits []models.OrderCredits
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.OrderCredits) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.OrderCredits[0]

	return data, nil
}

func ListOrderCreditsForUserByUserId(ctx context.Context, user_id int64) ([]models.OrderCredits, error) {
	var out []models.OrderCredits

	q := fragment_user + `query ListOrderCreditsForUser {
		order_credits(where: { user_id: { eq: $user_id } }) {
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
		OrderCredits []models.OrderCredits `json:"order_credits"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.OrderCredits

	return out, nil
}
