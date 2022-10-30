package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewLeadInstagram(ctx context.Context, data *models.LeadInstagram) error {
	err := graphql.NewLeadInstagram(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetLeadInstagram(ctx context.Context, id int64) (models.LeadInstagram, error) {
	data, err := graphql.GetLeadInstagram(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteLeadInstagram(ctx context.Context, id int64) error {
	err := graphql.DeleteLeadInstagram(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateLeadInstagram(ctx context.Context, data models.LeadInstagram) error {
	err := graphql.UpdateLeadInstagram(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
