package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_user_finance = ReflectToFragment(models.UserFinanceData{})
)

func NewUserFinance(ctx context.Context, data *models.UserFinance) error {
	q := `
		mutation CreateUserFinance {
			user_finance(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.UserFinance
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
		UserFinance []models.UserFinance
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.UserFinance) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.UserFinance[0].Id

	data.Id = &id

	return nil
}

func DeleteUserFinance(ctx context.Context, id int64) error {
	q := `
		mutation DeleteUserFinance {
			user_finance(where: { id: { eq: $id } }) {
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
		UserFinance []models.UserFinance
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.UserFinance) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateUserFinance(ctx context.Context, data models.UserFinance) error {
	q := `
		mutation UpdateUserFinance {
			user_finance(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.UserFinance
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
		UserFinance []models.UserFinance
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.UserFinance) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetUserFinance(ctx context.Context, id int64) (models.UserFinance, error) {
	var data models.UserFinance

	q := fragment_user_finance + `
			query GetUserFinance {
			user_finance(where: { id: { eq: $id } }) {
				...UserFinance
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
		UserFinance []models.UserFinance
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.UserFinance) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.UserFinance[0]

	return data, nil
}

func ListUserFinanceForUserByUserId(ctx context.Context, user_id int64) ([]models.UserFinance, error) {
	var out []models.UserFinance

	q := fragment_user + `query ListUserFinanceForUser {
		user_finance(where: { user_id: { eq: $user_id } }) {
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
		UserFinance []models.UserFinance `json:"user_finance"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.UserFinance

	return out, nil
}
