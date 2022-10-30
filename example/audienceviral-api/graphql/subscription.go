package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_subscription = ReflectToFragment(models.SubscriptionData{})
)

func NewSubscription(ctx context.Context, data *models.Subscription) error {
	q := `
		mutation CreateSubscription {
			subscription(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Subscription
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
		Subscription []models.Subscription
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Subscription) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Subscription[0].Id

	data.Id = &id

	return nil
}

func DeleteSubscription(ctx context.Context, id int64) error {
	q := `
		mutation DeleteSubscription {
			subscription(where: { id: { eq: $id } }) {
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
		Subscription []models.Subscription
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Subscription) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateSubscription(ctx context.Context, data models.Subscription) error {
	q := `
		mutation UpdateSubscription {
			subscription(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Subscription
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
		Subscription []models.Subscription
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Subscription) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetSubscription(ctx context.Context, id int64) (models.Subscription, error) {
	var data models.Subscription

	q := fragment_subscription + `
			query GetSubscription {
			subscription(where: { id: { eq: $id } }) {
				...Subscription
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
		Subscription []models.Subscription
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Subscription) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Subscription[0]

	return data, nil
}

func ListSubscriptionForUserByUserId(ctx context.Context, user_id int64) ([]models.Subscription, error) {
	var out []models.Subscription

	q := fragment_user + `query ListSubscriptionForUser {
		subscription(where: { user_id: { eq: $user_id } }) {
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
		Subscription []models.Subscription `json:"subscription"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.Subscription

	return out, nil
}
