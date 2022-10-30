package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewOrderCredits(ctx context.Context, data *models.OrderCredits) error {
	err := graphql.NewOrderCredits(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetOrderCredits(ctx context.Context, id int64) (models.OrderCredits, error) {
	data, err := graphql.GetOrderCredits(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteOrderCredits(ctx context.Context, id int64) error {
	err := graphql.DeleteOrderCredits(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOrderCredits(ctx context.Context, data models.OrderCredits) error {
	err := graphql.UpdateOrderCredits(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListOrderCreditsForUser(ctx context.Context, user_id int64) ([]models.OrderCredits, error) {
	data, err := graphql.ListOrderCreditsForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
