package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_user_search_history = ReflectToFragment(models.UserSearchHistoryData{})
)

func NewUserSearchHistory(ctx context.Context, data *models.UserSearchHistory) error {
	q := `
		mutation CreateUserSearchHistory {
			user_search_history(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.UserSearchHistory
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
		UserSearchHistory []models.UserSearchHistory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.UserSearchHistory) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.UserSearchHistory[0].Id

	data.Id = &id

	return nil
}

func DeleteUserSearchHistory(ctx context.Context, id int64) error {
	q := `
		mutation DeleteUserSearchHistory {
			user_search_history(where: { id: { eq: $id } }) {
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
		UserSearchHistory []models.UserSearchHistory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.UserSearchHistory) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateUserSearchHistory(ctx context.Context, data models.UserSearchHistory) error {
	q := `
		mutation UpdateUserSearchHistory {
			user_search_history(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.UserSearchHistory
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
		UserSearchHistory []models.UserSearchHistory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.UserSearchHistory) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetUserSearchHistory(ctx context.Context, id int64) (models.UserSearchHistory, error) {
	var data models.UserSearchHistory

	q := fragment_user_search_history + `
			query GetUserSearchHistory {
			user_search_history(where: { id: { eq: $id } }) {
				...UserSearchHistory
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
		UserSearchHistory []models.UserSearchHistory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.UserSearchHistory) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.UserSearchHistory[0]

	return data, nil
}

func ListUserSearchHistoryForUserByUserId(ctx context.Context, user_id int64) ([]models.UserSearchHistory, error) {
	var out []models.UserSearchHistory

	q := fragment_user + `query ListUserSearchHistoryForUser {
		user_search_history(where: { user_id: { eq: $user_id } }) {
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
		UserSearchHistory []models.UserSearchHistory `json:"user_search_history"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.UserSearchHistory

	return out, nil
}
