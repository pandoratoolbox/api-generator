package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewOrder(ctx context.Context, data *models.Order) error {
	err := graphql.NewOrder(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetOrder(ctx context.Context, id int64) (models.Order, error) {
	data, err := graphql.GetOrder(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteOrder(ctx context.Context, id int64) error {
	err := graphql.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOrder(ctx context.Context, data models.Order) error {
	err := graphql.UpdateOrder(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListOrderForUser(ctx context.Context, user_id int64) ([]models.Order, error) {
	data, err := graphql.ListOrderForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
