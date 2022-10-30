package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_lead = ReflectToFragment(models.LeadData{})
)

func NewLead(ctx context.Context, data *models.Lead) error {
	q := `
		mutation CreateLead {
			lead(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Lead
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
		Lead []models.Lead
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Lead) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Lead[0].Id

	data.Id = &id

	return nil
}

func DeleteLead(ctx context.Context, id int64) error {
	q := `
		mutation DeleteLead {
			lead(where: { id: { eq: $id } }) {
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
		Lead []models.Lead
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Lead) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateLead(ctx context.Context, data models.Lead) error {
	q := `
		mutation UpdateLead {
			lead(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Lead
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
		Lead []models.Lead
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Lead) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetLead(ctx context.Context, id int64) (models.Lead, error) {
	var data models.Lead

	q := fragment_lead + `
			query GetLead {
			lead(where: { id: { eq: $id } }) {
				...Lead
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
		Lead []models.Lead
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Lead) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Lead[0]

	return data, nil
}
