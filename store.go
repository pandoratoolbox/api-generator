package main

import (
	"strings"
)

var gen_store = `package store

`

var store_get_user_by_username = `
	
func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	out, err := graphql.GetUserByUsername(ctx, username)
	if err != nil {
		return out, err
	}
	
	return out, nil
}`

var store_get_user_by_email = `
	
func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	out, err := graphql.GetUserByEmail(ctx, email)
	if err != nil {
		return out, err
	}
	
	return out, nil
}`

func GenerateStoreFunctions(s Struct, package_name string) (string, error) {

	struct_name := ToUpperCase(s.Name)

	var out string

	getq := `

func Get{{struct_name}}(ctx context.Context, id int64) (*models.{{struct_name}}, error) {
		data, err := {{package_name}}.Get{{struct_name}}(ctx, id)

		if err != nil {
			return data, err
		}

		return data, nil
}
	`

	getq = strings.ReplaceAll(getq, "{{struct_name}}", struct_name)

	getq = strings.ReplaceAll(getq, "{{package_name}}", package_name)

	updateq := `

func Update{{struct_name}}(ctx context.Context, data models.{{struct_name}}) error {
		err := {{package_name}}.Update{{struct_name}}(ctx, data)
		if err != nil {
			return err
		}
	
		return nil
}`

	updateq = strings.ReplaceAll(updateq, "{{struct_name}}", struct_name)

	updateq = strings.ReplaceAll(updateq, "{{package_name}}", package_name)

	deleteq := `

func Delete{{struct_name}}(ctx context.Context, id int64) error {
		err := {{package_name}}.Delete{{struct_name}}(ctx, id)
		if err != nil {
			return err
		}
	
		return nil
}`

	deleteq = strings.ReplaceAll(deleteq, "{{struct_name}}", struct_name)

	deleteq = strings.ReplaceAll(deleteq, "{{package_name}}", package_name)

	createq := `

func New{{struct_name}}(ctx context.Context, data *models.{{struct_name}}) error {
		err := {{package_name}}.New{{struct_name}}(ctx, data)
		if err != nil {
			return err
		}
	
		return nil
}`

	createq = strings.ReplaceAll(createq, "{{struct_name}}", struct_name)

	createq = strings.ReplaceAll(createq, "{{package_name}}", package_name)

	out = createq + `
	
	` +
		getq + `
	
	` +
		deleteq + `
	
	` +
		updateq

	for _, col := range s.Columns {

		if s.Name == "user" {
			if col.Name == "username" {
				out += store_get_user_by_username
			}

			if col.Name == "email" {
				out += store_get_user_by_email
			}

		}

		// 		if col.ForeignKeyTable == "user" {
		// 			list := strings.ReplaceAll(`

		// func List{{struct_name_upper}}ForUser(ctx context.Context, user_id int64) ([]models.{{struct_name_upper}}, error) {
		// 	data, err := graphql.List{{struct_name_upper}}ForUserByUser{{foreign_column_name_upper}}(ctx, user_id)
		// 	if err != nil {
		// 		return data, err
		// 	}

		// 	return data, nil
		// }`, "{{struct_name_upper}}", struct_name)
		// 			list = strings.ReplaceAll(list, "{{foreign_column_name_upper}}", ToUpperCase(col.ForeignKeyColumn))
		// 			out += list
		// 		}

		if col.IsForeignKey {
			if col.IsUnique {
				getfkbyid := `

					func Get{{struct_name_upper}}By{{column_name_upper}}(ctx context.Context, id int64) (*models.{{struct_name_upper}}, error) {
									data, err := graphql.Get{{struct_name_upper}}By{{column_name_upper}}(ctx, id)
									if err != nil {
										return data, err
									}
			
									return data, nil
					}`

				getfkbyid = strings.ReplaceAll(getfkbyid, "{{struct_name_upper}}", ToUpperCase(s.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{column_name_upper}}", ToUpperCase(col.Name))
				out += getfkbyid
			} else {
				getfkbyid := `

				func List{{struct_name_upper}}By{{column_name_upper}}(ctx context.Context, id int64) ([]models.{{struct_name_upper}}, error) {
								data, err := graphql.List{{struct_name_upper}}By{{column_name_upper}}(ctx, id)
								if err != nil {
									return data, err
								}
		
								return data, nil
				}`

				getfkbyid = strings.ReplaceAll(getfkbyid, "{{struct_name_upper}}", ToUpperCase(s.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{column_name_upper}}", ToUpperCase(col.Name))
				out += getfkbyid
			}
		}
	}

	return out, nil
}
