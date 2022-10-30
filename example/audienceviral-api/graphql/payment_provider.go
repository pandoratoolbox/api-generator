package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_payment_provider = ReflectToFragment(models.PaymentProviderData{})
)

func NewPaymentProvider(ctx context.Context, data *models.PaymentProvider) error {
	q := `
		mutation CreatePaymentProvider {
			payment_provider(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.PaymentProvider
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
		PaymentProvider []models.PaymentProvider
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProvider) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.PaymentProvider[0].Id

	data.Id = &id

	return nil
}

func DeletePaymentProvider(ctx context.Context, id int64) error {
	q := `
		mutation DeletePaymentProvider {
			payment_provider(where: { id: { eq: $id } }) {
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
		PaymentProvider []models.PaymentProvider
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProvider) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdatePaymentProvider(ctx context.Context, data models.PaymentProvider) error {
	q := `
		mutation UpdatePaymentProvider {
			payment_provider(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.PaymentProvider
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
		PaymentProvider []models.PaymentProvider
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProvider) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetPaymentProvider(ctx context.Context, id int64) (models.PaymentProvider, error) {
	var data models.PaymentProvider

	q := fragment_payment_provider + `
			query GetPaymentProvider {
			payment_provider(where: { id: { eq: $id } }) {
				...PaymentProvider
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
		PaymentProvider []models.PaymentProvider
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.PaymentProvider) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.PaymentProvider[0]

	return data, nil
}
