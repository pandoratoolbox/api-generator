package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_lead_instagram = ReflectToFragment(models.LeadInstagramData{})
)

func NewLeadInstagram(ctx context.Context, data *models.LeadInstagram) error {
	q := `
		mutation CreateLeadInstagram {
			lead_instagram(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.LeadInstagram
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
		LeadInstagram []models.LeadInstagram
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.LeadInstagram) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.LeadInstagram[0].Id

	data.Id = &id

	return nil
}

func DeleteLeadInstagram(ctx context.Context, id int64) error {
	q := `
		mutation DeleteLeadInstagram {
			lead_instagram(where: { id: { eq: $id } }) {
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
		LeadInstagram []models.LeadInstagram
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.LeadInstagram) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateLeadInstagram(ctx context.Context, data models.LeadInstagram) error {
	q := `
		mutation UpdateLeadInstagram {
			lead_instagram(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.LeadInstagram
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
		LeadInstagram []models.LeadInstagram
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.LeadInstagram) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetLeadInstagram(ctx context.Context, id int64) (models.LeadInstagram, error) {
	var data models.LeadInstagram

	q := fragment_lead_instagram + `
			query GetLeadInstagram {
			lead_instagram(where: { id: { eq: $id } }) {
				...LeadInstagram
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
		LeadInstagram []models.LeadInstagram
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.LeadInstagram) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.LeadInstagram[0]

	return data, nil
}
