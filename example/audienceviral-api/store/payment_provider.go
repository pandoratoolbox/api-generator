package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewPaymentProvider(ctx context.Context, data *models.PaymentProvider) error {
	err := graphql.NewPaymentProvider(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetPaymentProvider(ctx context.Context, id int64) (models.PaymentProvider, error) {
	data, err := graphql.GetPaymentProvider(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeletePaymentProvider(ctx context.Context, id int64) error {
	err := graphql.DeletePaymentProvider(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePaymentProvider(ctx context.Context, data models.PaymentProvider) error {
	err := graphql.UpdatePaymentProvider(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
