package store

import (
	"audienceviral-api/graphql"
	"audienceviral-api/models"
	"context"
)

func NewPaymentProviderCustomer(ctx context.Context, data *models.PaymentProviderCustomer) error {
	err := graphql.NewPaymentProviderCustomer(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetPaymentProviderCustomer(ctx context.Context, id int64) (models.PaymentProviderCustomer, error) {
	data, err := graphql.GetPaymentProviderCustomer(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeletePaymentProviderCustomer(ctx context.Context, id int64) error {
	err := graphql.DeletePaymentProviderCustomer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePaymentProviderCustomer(ctx context.Context, data models.PaymentProviderCustomer) error {
	err := graphql.UpdatePaymentProviderCustomer(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListPaymentProviderCustomerForUser(ctx context.Context, user_id int64) ([]models.PaymentProviderCustomer, error) {
	data, err := graphql.ListPaymentProviderCustomerForUserById(ctx, user_id)
	if err != nil {
		return data, err
	}

	return data, nil
}
