package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewPaymentProviderSource(ctx context.Context, data *models.PaymentProviderSource) error {
	err := graphql.NewPaymentProviderSource(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetPaymentProviderSource(ctx context.Context, id int64) (models.PaymentProviderSource, error) {
	data, err := graphql.GetPaymentProviderSource(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeletePaymentProviderSource(ctx context.Context, id int64) error {
	err := graphql.DeletePaymentProviderSource(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePaymentProviderSource(ctx context.Context, data models.PaymentProviderSource) error {
	err := graphql.UpdatePaymentProviderSource(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListPaymentProviderSourceForUser(ctx context.Context, user_id int64) ([]models.PaymentProviderSource, error) {
	data, err := graphql.ListPaymentProviderSourceForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
