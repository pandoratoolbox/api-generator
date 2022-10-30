package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_payment_provider_source = ReflectToFragment(models.PaymentProviderSourceData{})
)

func NewPaymentProviderSource(ctx context.Context, data *models.PaymentProviderSource) error {
	q := `
		mutation CreatePaymentProviderSource {
			payment_provider_source(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.PaymentProviderSource
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
		PaymentProviderSource []models.PaymentProviderSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProviderSource) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.PaymentProviderSource[0].Id

	data.Id = &id

	return nil
}

func DeletePaymentProviderSource(ctx context.Context, id int64) error {
	q := `
		mutation DeletePaymentProviderSource {
			payment_provider_source(where: { id: { eq: $id } }) {
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
		PaymentProviderSource []models.PaymentProviderSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProviderSource) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdatePaymentProviderSource(ctx context.Context, data models.PaymentProviderSource) error {
	q := `
		mutation UpdatePaymentProviderSource {
			payment_provider_source(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.PaymentProviderSource
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
		PaymentProviderSource []models.PaymentProviderSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.PaymentProviderSource) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetPaymentProviderSource(ctx context.Context, id int64) (models.PaymentProviderSource, error) {
	var data models.PaymentProviderSource

	q := fragment_payment_provider_source + `
			query GetPaymentProviderSource {
			payment_provider_source(where: { id: { eq: $id } }) {
				...PaymentProviderSource
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
		PaymentProviderSource []models.PaymentProviderSource
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.PaymentProviderSource) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.PaymentProviderSource[0]

	return data, nil
}

func ListPaymentProviderSourceForUserByUserId(ctx context.Context, user_id int64) ([]models.PaymentProviderSource, error) {
	var out []models.PaymentProviderSource

	q := fragment_user + `query ListPaymentProviderSourceForUser {
		payment_provider_source(where: { user_id: { eq: $user_id } }) {
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
		PaymentProviderSource []models.PaymentProviderSource `json:"payment_provider_source"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.PaymentProviderSource

	return out, nil
}
