package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_user = ReflectToFragment(models.UserData{})
)

func NewUser(ctx context.Context, data *models.User) error {
	q := `
		mutation CreateUser {
			user(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.User
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
		User []models.User
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.User) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.User[0].Id

	data.Id = &id

	return nil
}

func DeleteUser(ctx context.Context, id int64) error {
	q := `
		mutation DeleteUser {
			user(where: { id: { eq: $id } }) {
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
		User []models.User
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.User) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateUser(ctx context.Context, data models.User) error {
	q := `
		mutation UpdateUser {
			user(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.User
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
		User []models.User
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.User) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetUser(ctx context.Context, id int64) (models.User, error) {
	var data models.User

	q := fragment_user + `
			query GetUser {
			user(where: { id: { eq: $id } }) {
				...User
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
		User []models.User
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.User) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.User[0]

	return data, nil
}

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var out models.User

	q := fragment_user + `query GetUserByEmail {
		user(where: { email: { eq: $email } }) {
			...User
		}
	}`

	input := struct {
		Email string `json:"email"`
	}{
		Email: email,
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
		User []models.User `json:"user"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	if len(rt.User) < 1 {
		return out, errors.New("Unable to find user")
	}

	return out, nil
}

func ListUserForUserByReferrerId(ctx context.Context, user_id int64) ([]models.User, error) {
	var out []models.User

	q := fragment_user + `query ListUserForUser {
		user(where: { referrer_id: { eq: $user_id } }) {
			...User
		}
	}`

	input := struct {
		ReferrerId int64 `json:"referrer_id"`
	}{
		ReferrerId: user_id,
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
		User []models.User `json:"user"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.User

	return out, nil
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var out models.User

	q := fragment_user + `query GetUserByUsername {
		user(where: { username: { eq: $username } }) {
			...User
		}
	}`

	input := struct {
		Username string `json:"username"`
	}{
		Username: username,
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
		User []models.User `json:"user"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	if len(rt.User) < 1 {
		return out, errors.New("Unable to find user")
	}

	return out, nil
}
