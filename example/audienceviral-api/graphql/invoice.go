package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_invoice = ReflectToFragment(models.InvoiceData{})
)

func NewInvoice(ctx context.Context, data *models.Invoice) error {
	q := `
		mutation CreateInvoice {
			invoice(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Invoice
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
		Invoice []models.Invoice
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Invoice) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Invoice[0].Id

	data.Id = &id

	return nil
}

func DeleteInvoice(ctx context.Context, id int64) error {
	q := `
		mutation DeleteInvoice {
			invoice(where: { id: { eq: $id } }) {
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
		Invoice []models.Invoice
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Invoice) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateInvoice(ctx context.Context, data models.Invoice) error {
	q := `
		mutation UpdateInvoice {
			invoice(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Invoice
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
		Invoice []models.Invoice
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Invoice) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetInvoice(ctx context.Context, id int64) (models.Invoice, error) {
	var data models.Invoice

	q := fragment_invoice + `
			query GetInvoice {
			invoice(where: { id: { eq: $id } }) {
				...Invoice
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
		Invoice []models.Invoice
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Invoice) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Invoice[0]

	return data, nil
}

func ListInvoiceForUserByUserId(ctx context.Context, user_id int64) ([]models.Invoice, error) {
	var out []models.Invoice

	q := fragment_user + `query ListInvoiceForUser {
		invoice(where: { user_id: { eq: $user_id } }) {
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
		Invoice []models.Invoice `json:"invoice"`
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	out = rt.Invoice

	return out, nil
}
