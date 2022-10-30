package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewInvoice(ctx context.Context, data *models.Invoice) error {
	err := graphql.NewInvoice(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetInvoice(ctx context.Context, id int64) (models.Invoice, error) {
	data, err := graphql.GetInvoice(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteInvoice(ctx context.Context, id int64) error {
	err := graphql.DeleteInvoice(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateInvoice(ctx context.Context, data models.Invoice) error {
	err := graphql.UpdateInvoice(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListInvoiceForUser(ctx context.Context, user_id int64) ([]models.Invoice, error) {
	data, err := graphql.ListInvoiceForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
