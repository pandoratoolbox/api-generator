package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewUserFinance(ctx context.Context, data *models.UserFinance) error {
	err := graphql.NewUserFinance(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetUserFinance(ctx context.Context, id int64) (models.UserFinance, error) {
	data, err := graphql.GetUserFinance(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteUserFinance(ctx context.Context, id int64) error {
	err := graphql.DeleteUserFinance(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserFinance(ctx context.Context, data models.UserFinance) error {
	err := graphql.UpdateUserFinance(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListUserFinanceForUser(ctx context.Context, user_id int64) ([]models.UserFinance, error) {
	data, err := graphql.ListUserFinanceForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
