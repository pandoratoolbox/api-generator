package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewSubscription(ctx context.Context, data *models.Subscription) error {
	err := graphql.NewSubscription(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetSubscription(ctx context.Context, id int64) (models.Subscription, error) {
	data, err := graphql.GetSubscription(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteSubscription(ctx context.Context, id int64) error {
	err := graphql.DeleteSubscription(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSubscription(ctx context.Context, data models.Subscription) error {
	err := graphql.UpdateSubscription(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListSubscriptionForUser(ctx context.Context, user_id int64) ([]models.Subscription, error) {
	data, err := graphql.ListSubscriptionForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
