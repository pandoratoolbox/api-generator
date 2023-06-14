package main

import (
	"fmt"
	"strings"
)

var gen_graphql = `package graphql

`

var graphql_init = `package graphql

` + GenerateImport(PACKAGE_GRAPHQL, PACKAGE_ERRORS, PACKAGE_CONNECTIONS, PACKAGE_REFLECT, PACKAGE_STRINGS, PACKAGE_UNICODE) + strings.ReplaceAll(`

var Graph *core.GraphJin
func ToUpperCamelCase(s string) string {
	rs := []rune(s)
	new := []rune{}
	next_upper := true

	for _, r := range rs {
		if next_upper {
			new = append(new, unicode.ToUpper(r))
			next_upper = false
			continue
		}
		if unicode.IsPunct(r) {
			next_upper = true
			continue
		}

		new = append(new, r)
	}

	return string(new)

}

func ToSnakeCase(s string) string {
	var res = make([]rune, 0, len(s))
	var p = '_'
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			res = append(res, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				res = append(res, '_', unicode.ToLower(r))
			} else {
				res = append(res, unicode.ToLower(r))
			}
		} else {
			res = append(res, unicode.ToLower(r))
		}

		p = r
	}
	return string(res)
}

func ReflectToFragment(data interface{}) string {
	var fragment string
	var fieldsList string
	var structName string
	var structSnake string

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	structName = v.Type().Name()
	if strings.Contains(structName, "Data") {
		structName = strings.Replace(structName, "Data", "", -1)
	}
	structSnake = ToSnakeCase(structName)

	for i := 0; i < v.NumField(); i++ {
		var fieldName string
		var fieldSnake string

		fieldName = v.Type().Field(i).Name
		fieldSnake = ToSnakeCase(fieldName)

		if i == 0 {
			fieldsList = fieldSnake
			continue
		}
		fieldsList = fieldsList + @@
		@@ + fieldSnake

	}

	fragment = fmt.Sprintf(@@
	
	fragment %s on %s {
		%s
	}
	
	@@, structName, structSnake, fieldsList)

	fmt.Println(fragment)

	return fragment
}


func Init() error {
	var err error
	Graph, err = core.NewGraphJin(nil, connections.Postgres)

	if err != nil {
		return err
	}

	return nil
}`, "@@", "`")

var graphql_user_by_username = strings.Replace(`

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var out models.User


	q := fragment_user + @@query GetUserByUsername {
		user(where: { username: { eq: $username } }) {
			...User
		}
	}@@


	input := struct{
		Username string @@json:"username"@@
	}{
		Username: username,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	rt := struct{
		User []models.User @@json:"user"@@
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	if len(rt.User) < 1 {
		return out, errors.New("Unable to find user")
	}

	out = rt.User[0]

	return out, nil
	}`, "@@", "`", -1)

var graphql_user_by_email = strings.Replace(`

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var out models.User


	q := fragment_user + @@query GetUserByEmail {
		user(where: { email: { eq: $email } }) {
			...User
		}
	}@@


	input := struct{
		Email string @@json:"email"@@
	}{
		Email: email,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	rt := struct{
		User []models.User @@json:"user"@@
	}{}

	err = json.Unmarshal(res.Data, &rt)
	if err != nil {
		return out, err
	}

	if len(rt.User) < 1 {
		return out, errors.New("Unable to find user")
	}

	out = rt.User[0]

	return out, nil
	}`, "@@", "`", -1)

func GenerateGraphqlQueries(s Struct, list_fields ...string) (string, error) {
	var c string

	structname := ToUpperCase(s.Name)

	structname_snake := ToSnakeCase(s.Name)

	fragment := strings.ReplaceAll(`var (
		fragment_{{struct_name_snake}} = ReflectToFragment(models.{{struct_name}}Data{})
	)
	`, "{{struct_name_snake}}", structname_snake)

	fragment = strings.ReplaceAll(`var (
	fragment_{{struct_name_snake}} = ReflectToFragment(models.{{struct_name}}Data{})
)
`, "{{struct_name}}", structname)

	fragment = strings.ReplaceAll(fragment, "{{struct_name}}", structname)
	fragment = strings.ReplaceAll(fragment, "{{struct_name_snake}}", structname_snake)

	// q_fields := fmt.Sprintf("...%s", structname)

	// q_fields += `
	// ...{{foreign_struct_name}}`

	for _, col := range s.Columns {

		if s.Name == "user" {
			if col.Name == "username" {
				c += graphql_user_by_username
			}
			if col.Name == "email" {
				c += graphql_user_by_email
			}
		}

		// 		if col.ForeignKeyTable == "user" {
		// 			lbu := `

		// func List{{struct_name_upper}}ForUserBy{{column_name_upper}}(ctx context.Context, user_id int64) ([]models.{{struct_name_upper}}, error) {
		// 	var out []models.{{struct_name_upper}}

		// 	q := fragment_{{struct_name_snake}} + @@query List{{struct_name_upper}}ForUser {
		// 		{{struct_name_snake}}(where: { {{column_name_snake}}: { eq: $id } }) {
		// 			...{{struct_table_upper}}
		// 		}
		// 	}@@

		// 	input := struct{
		// 		Id int64 @@json:"id"@@
		// 	}{
		// 		Id: user_id,
		// 	}

		// 	js, err := json.Marshal(input)
		// 	if err != nil {
		// 		return out, err
		// 	}

		// 	res, err := Graph.GraphQL(ctx, q, js, nil)
		// 	if err != nil {
		// 		return out, err
		// 	}

		// 	rt := struct{
		// 		{{struct_name_upper}} []models.{{struct_name_upper}} @@json:"{{struct_name_snake}}"@@
		// 	}{}

		// 	err = json.Unmarshal(res.Data, &rt)
		// 	if err != nil {
		// 		return out, err
		// 	}

		// 	out = rt.{{struct_name_upper}}

		// 	return out, nil
		// }`

		// 			lbu = strings.ReplaceAll(lbu, "@@", "`")
		// 			lbu = strings.ReplaceAll(lbu, "{{struct_name_upper}}", structname)
		// 			lbu = strings.ReplaceAll(lbu, "{{struct_name_snake}}", structname_snake)
		// 			lbu = strings.ReplaceAll(lbu, "{{foreign_table_upper}}", ToUpperCase(col.ForeignKeyTable))
		// 			lbu = strings.ReplaceAll(lbu, "{{foreign_table_snake}}", ToSnakeCase(col.ForeignKeyTable))
		// 			lbu = strings.ReplaceAll(lbu, "{{column_name_snake}}", ToSnakeCase(col.Name))
		// 			lbu = strings.ReplaceAll(lbu, "{{column_name_upper}}", ToUpperCase(col.Name))

		// 			c += lbu
		// 		}

		if col.IsForeignKey {

			if col.IsUnique {
				getfkbyid := `

				func Get{{struct_name_upper}}By{{column_name_upper}}(ctx context.Context, id int64) (models.{{struct_name_upper}}, error) {
					var out models.{{struct_name_upper}}

					q := fragment_{{struct_name_snake}}+""query Get{{struct_name_upper}}By{{column_name_upper}}(where: { {{column_name_snake}}: { eq: $id }}) {
						...{{struct_name_upper}}
					}""

					input := struct{
						Id int64 ""%s""
					}{
						Id: id,
					}

					js, err := json.Marshal(input)
					if err != nil {
						return out, err
					}

					res, err := Graph.GraphQL(ctx, q, js, nil)
					if err != nil {
						return out, err
					}

					ret := struct{
						{{struct_name_upper}} []models.{{struct_name_upper}}
					}{}

					err = json.Unmarshal(res.Data, &ret)
					if err != nil {
						return out, err
					}

					if len(ret.{{struct_name_upper}}) < 1 {
						return out, errors.New("Object not found")
					}

					out = ret.{{struct_name_upper}}[0]

					return out, nil
				}`
				getfkbyid = strings.ReplaceAll(getfkbyid, `""`, "`")
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{struct_name_upper}}", ToUpperCase(s.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{struct_name_snake}}", ToSnakeCase(s.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{column_name_upper}}", ToUpperCase(col.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{column_name_snake}}", ToSnakeCase(col.Name))
				getfkbyid = fmt.Sprintf(getfkbyid, `json:"id"`)
				c += getfkbyid
			} else {
				getfkbyid := `

				func List{{struct_name_upper}}By{{column_name_upper}}(ctx context.Context, id int64) ([]models.{{struct_name_upper}}, error) {
					var out []models.{{struct_name_upper}}
	
					q := fragment_{{struct_name_snake}}+""query List{{struct_name_upper}}By{{column_name_upper}}(where: { {{column_name_snake}}: { eq: $id }}) {
						...{{struct_name_upper}}
					}""
	
					input := struct{
						Id int64 ""%s""
					}{
						Id: id,
					}
	
					js, err := json.Marshal(input)
					if err != nil {
						return out, err
					}
	
					res, err := Graph.GraphQL(ctx, q, js, nil)
					if err != nil {
						return out, err
					}
	
					ret := struct{
						{{struct_name_upper}} []models.{{struct_name_upper}}
					}{}
	
					err = json.Unmarshal(res.Data, &ret)
					if err != nil {
						return out, err
					}
	
					if len(ret.{{struct_name_upper}}) < 1 {
						return out, errors.New("Object not found")
					}
	
					out = ret.{{struct_name_upper}}
	
					return out, nil
				}`
				getfkbyid = strings.ReplaceAll(getfkbyid, `""`, "`")
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{struct_name_upper}}", ToUpperCase(s.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{struct_name_snake}}", ToSnakeCase(s.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{column_name_upper}}", ToUpperCase(col.Name))
				getfkbyid = strings.ReplaceAll(getfkbyid, "{{column_name_snake}}", ToSnakeCase(col.Name))
				getfkbyid = fmt.Sprintf(getfkbyid, `json:"id"`)
				c += getfkbyid
			}

		}

	}

	getbyid := `
	
	func Get{{struct_name}}(ctx context.Context, id int64) (models.{{struct_name}}, error) {
		var data models.{{struct_name}}
	
		q := fragment_{{struct_name_snake}} + ""
			query Get{{struct_name}} {
			{{struct_name_snake}}(where: { id: { eq: $id } }) {
				...{{struct_name}}
			}
		}
		""
	
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
			{{struct_name}} []models.{{struct_name}}
		}
	
		err = json.Unmarshal(res.Data, &out)
		if err != nil {
			return data, err
		}
	
		if len(out.{{struct_name}}) < 1 {
			return data, errors.New("Unable to retrieve object")
		}
	
		data = out.{{struct_name}}[0]
	
		return data, nil
	}`

	getbyid = strings.ReplaceAll(getbyid, `""`, "`")
	getbyid = strings.ReplaceAll(getbyid, "{{struct_name}}", structname)
	getbyid = strings.ReplaceAll(getbyid, "{{struct_name_snake}}", structname_snake)

	deletebyid := `
	
	func Delete{{struct_name}}(ctx context.Context, id int64) error {
		q := ""
		mutation Delete{{struct_name}} {
			{{struct_name_snake}}(where: { id: { eq: $id } }) {
				id
			}
		}
		""
	
		input := struct {
			Id  int64
		}{
			Id:  id,
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
			{{struct_name}} []models.{{struct_name}}
		}
	
		err = json.Unmarshal(res.Data, &out)
		if err != nil {
			return err
		}
	
		if len(out.{{struct_name}}) < 1 {
			return errors.New("Unable to delete object")
		}
	
		return nil
	}`

	deletebyid = strings.ReplaceAll(deletebyid, `""`, "`")
	deletebyid = strings.ReplaceAll(deletebyid, "{{struct_name}}", structname)
	deletebyid = strings.ReplaceAll(deletebyid, "{{struct_name_snake}}", structname_snake)

	updatebyid := `
	
	func Update{{struct_name}}(ctx context.Context, data models.{{struct_name}}) error {
		q := ""
		mutation Update{{struct_name}} {
			{{struct_name_snake}}(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		""

		input := struct {
			Id   int64
			Data models.{{struct_name}}
		}{
			Id:  *data.Id,
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
			{{struct_name}} []models.{{struct_name}}
		}
	
		err = json.Unmarshal(res.Data, &out)
		if err != nil {
			return err
		}
	
		if len(out.{{struct_name}}) < 1 {
			return errors.New("Unable to update object")
		}
	
		return nil
	}
	`

	updatebyid = strings.ReplaceAll(updatebyid, `""`, "`")
	updatebyid = strings.ReplaceAll(updatebyid, "{{struct_name}}", structname)
	updatebyid = strings.ReplaceAll(updatebyid, "{{struct_name_snake}}", structname_snake)

	createbyid := strings.ReplaceAll(`
	
	func New{{struct_name}}(ctx context.Context, data *models.{{struct_name}}) error {
		q := ""
		mutation Create{{struct_name}} {
			{{struct_name_snake}}(insert: $data) {
				id
			}
		}
		""

		input := struct {
			Data models.{{struct_name}}
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
			{{struct_name}} []models.{{struct_name}}
		}
	
		err = json.Unmarshal(res.Data, &out)
		if err != nil {
			return err
		}
	
		if len(out.{{struct_name}}) < 1 {
			return errors.New("Unable to insert object")
		}
	
		id := *out.{{struct_name}}[0].Id
	
		data.Id = &id
	
		return nil
	}`, `""`, "`")

	createbyid = strings.ReplaceAll(createbyid, "{{struct_name}}", structname)
	createbyid = strings.ReplaceAll(createbyid, "{{struct_name_snake}}", structname_snake)

	c = fragment + `
	
	` + createbyid + `
	
	` + deletebyid + `
	
	` + updatebyid + `
	
	` + getbyid + c

	return c, nil
}
