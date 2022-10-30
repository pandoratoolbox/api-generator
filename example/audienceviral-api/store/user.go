package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewUser(ctx context.Context, data *models.User) error {
	err := graphql.NewUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(ctx context.Context, id int64) (models.User, error) {
	data, err := graphql.GetUser(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteUser(ctx context.Context, id int64) error {
	err := graphql.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(ctx context.Context, data models.User) error {
	err := graphql.UpdateUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	out, err := graphql.GetUserByEmail(ctx, email)
	if err != nil {
		return out, err
	}

	return out, nil
}

func ListUserForUser(ctx context.Context, user_id int64) ([]models.User, error) {
	data, err := graphql.ListUserForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	out, err := graphql.GetUserByUsername(ctx, username)
	if err != nil {
		return out, err
	}

	return out, nil
}
