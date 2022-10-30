package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewDataPlatform(ctx context.Context, data *models.DataPlatform) error {
	err := graphql.NewDataPlatform(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetDataPlatform(ctx context.Context, id int64) (models.DataPlatform, error) {
	data, err := graphql.GetDataPlatform(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteDataPlatform(ctx context.Context, id int64) error {
	err := graphql.DeleteDataPlatform(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDataPlatform(ctx context.Context, data models.DataPlatform) error {
	err := graphql.UpdateDataPlatform(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
