package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_payment_provider_customer = ReflectToFragment(models.PaymentProviderCustomerData{})
)

func NewPaymentProviderCustomer(ctx context.Context, data *models.PaymentProviderCustomer) error {
	q := `
		mutation CreatePaymentProviderCustomer {
			payment_provider_customer(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.PaymentProviderCustomer
	}{
		Data: *data,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		PaymentProviderCustomer []models.PaymentProviderCustomer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProviderCustomer) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.PaymentProviderCustomer[0].Id

	data.Id = &id

	return nil
}

func DeletePaymentProviderCustomer(ctx context.Context, id int64) error {
	q := `
		mutation DeletePaymentProviderCustomer {
			payment_provider_customer(where: { id: { eq: $id } }) {
				id
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		PaymentProviderCustomer []models.PaymentProviderCustomer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProviderCustomer) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdatePaymentProviderCustomer(ctx context.Context, data models.PaymentProviderCustomer) error {
	q := `
		mutation UpdatePaymentProviderCustomer {
			payment_provider_customer(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.PaymentProviderCustomer
	}{
		Id: *data.Id,
	}

	data.Id = nil
	input.Data = data

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		PaymentProviderCustomer []models.PaymentProviderCustomer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProviderCustomer) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetPaymentProviderCustomer(ctx context.Context, id int64) (models.PaymentProviderCustomer, error) {
	var data models.PaymentProviderCustomer

	q := fragment_payment_provider_customer + `
			query GetPaymentProviderCustomer {
			payment_provider_customer(where: { id: { eq: $id } }) {
				...PaymentProviderCustomer
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return data, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		PaymentProviderCustomer []models.PaymentProviderCustomer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.PaymentProviderCustomer) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.PaymentProviderCustomer[0]

	return data, nil
}

func ListPaymentProviderCustomerForUserByUserId(ctx context.Context, user_id int64) ([]models.PaymentProviderCustomer, error) {
	var out []models.PaymentProviderCustomer

	q := fragment_user + `query ListPaymentProviderCustomerForUser {
		payment_provider_customer(where: { user_id: { eq: $user_id } }) {
			...User
		}
	}`

	input := struct {
		UserId int64 `json:"user_id"`
	}{
		UserId: user_id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	rt := struct {
		PaymentProviderCustomer []models.PaymentProviderCustomer `json:"payment_provider_customer"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.PaymentProviderCustomer

	return out, nil
}
