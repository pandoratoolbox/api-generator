package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewUserSearchHistory(ctx context.Context, data *models.UserSearchHistory) error {
	err := graphql.NewUserSearchHistory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetUserSearchHistory(ctx context.Context, id int64) (models.UserSearchHistory, error) {
	data, err := graphql.GetUserSearchHistory(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteUserSearchHistory(ctx context.Context, id int64) error {
	err := graphql.DeleteUserSearchHistory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserSearchHistory(ctx context.Context, data models.UserSearchHistory) error {
	err := graphql.UpdateUserSearchHistory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListUserSearchHistoryForUser(ctx context.Context, user_id int64) ([]models.UserSearchHistory, error) {
	data, err := graphql.ListUserSearchHistoryForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
