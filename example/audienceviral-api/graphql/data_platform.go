package graphql

import (
	"audienceviral-api/models"
	"context"
	"errors"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_data_platform = ReflectToFragment(models.DataPlatformData{})
)

func NewDataPlatform(ctx context.Context, data *models.DataPlatform) error {
	q := `
		mutation CreateDataPlatform {
			data_platform(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.DataPlatform
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
		DataPlatform []models.DataPlatform
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.DataPlatform) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.DataPlatform[0].Id

	data.Id = &id

	return nil
}

func DeleteDataPlatform(ctx context.Context, id int64) error {
	q := `
		mutation DeleteDataPlatform {
			data_platform(where: { id: { eq: $id } }) {
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
		DataPlatform []models.DataPlatform
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.DataPlatform) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateDataPlatform(ctx context.Context, data models.DataPlatform) error {
	q := `
		mutation UpdateDataPlatform {
			data_platform(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.DataPlatform
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
		DataPlatform []models.DataPlatform
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.DataPlatform) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetDataPlatform(ctx context.Context, id int64) (models.DataPlatform, error) {
	var data models.DataPlatform

	q := fragment_data_platform + `
			query GetDataPlatform {
			data_platform(where: { id: { eq: $id } }) {
				...DataPlatform
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
		DataPlatform []models.DataPlatform
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.DataPlatform) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.DataPlatform[0]

	return data, nil
}
