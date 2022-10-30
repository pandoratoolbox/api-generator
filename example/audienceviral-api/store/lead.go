package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewLead(ctx context.Context, data *models.Lead) error {
	err := graphql.NewLead(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetLead(ctx context.Context, id int64) (models.Lead, error) {
	data, err := graphql.GetLead(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteLead(ctx context.Context, id int64) error {
	err := graphql.DeleteLead(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateLead(ctx context.Context, data models.Lead) error {
	err := graphql.UpdateLead(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
