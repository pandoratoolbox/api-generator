package main

import (
	"strings"
)

var gen_handlers = `package handlers

`

func GenerateCoreMiddleware() string {
	var out string

	out = `
	
	func RestrictAuth(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !r.Context().Value(models.CTX_is_auth).(bool) {
				http.Error(w, "Restricted access, please log in.", http.StatusUnauthorized)
				return
			}
	
			next.ServeHTTP(w, r)
		})
	}
	`

	out += `
	func Authenticator(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			_, claims, err := jwtauth.FromContext(ctx)
			if err != nil {
				ctx = context.WithValue(ctx, models.CTX_is_auth, false)
			} else {
				claimsRole := claims["role_ids"]
				if claimsRole == nil {
					ctx = context.WithValue(ctx, models.CTX_is_auth, false)
				} else {
					var rids []int64
					rrids := claimsRole.([]interface{})
					for _, rrid := range rrids {
						rids = append(rids, int64(rrid.(float64)))
					}

					ctx = context.WithValue(ctx, models.CTX_user_role_ids, rids)
				}
	
				claimsId := claims["id"]
				if claimsId == nil {
					ctx = context.WithValue(ctx, models.CTX_is_auth, false)
				} else {
					ctx = context.WithValue(ctx, models.CTX_user_id, int64(claimsId.(float64)))
					ctx = context.WithValue(ctx, models.CTX_is_auth, true)
				}
	
			}
	
			r = r.WithContext(ctx)
	
			next.ServeHTTP(w, r)
		})
	}
	`

	return out
}

func GenerateHandlerCore() string {
	out := `
	
	func ServeError(w http.ResponseWriter, message string, code int) {
		fmt.Printf("Http error: %s\n", message)
		http.Error(w, message, code)
	}
	
	func ServeJSON(w http.ResponseWriter, data interface{}) {
		js, err := json.Marshal(data)
	
		if err != nil {
			ServeError(w, err.Error(), 400)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	`

	return out
}

func GenerateAuthHandlers() string {
	pkg := `
	
	var (
		TokenAuth = jwtauth.New("HS256", []byte("h1l32b"), nil)
	)`

	handlers := `

func Login(w http.ResponseWriter, r *http.Request) {
input := models.User{}
ctx := r.Context()

decoder := json.NewDecoder(r.Body)

err := decoder.Decode(&input)
if err != nil {
	ServeError(w, err.Error(), 400)
	return
}

user, err := store.GetUserByUsername(ctx, *input.Username)
if err != nil {
	ServeError(w, err.Error(), 400)
	return
}

if *user.Password != *input.Password {
	ServeError(w, errors.New("Wrong password").Error(), 400)
	return
}

_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
	"id":       	 *user.Id,
	"role_ids":      *user.RoleIds,
})
if err != nil {
	ServeError(w, err.Error(), 400)
	return
}

out := struct{
	Token string
}{
	Token: jwtstring,
}

js, err := json.Marshal(out)
if err != nil {
	ServeError(w, err.Error(), 400)
	return
}

w.Write(js)
}

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	roles := models.Ints{1}

	user := models.User{
	 	UserData: models.UserData{
	 		Username: data.Username,
	 		Password: data.Password,
	 		Email:    data.Email,
	 		RoleIds:     &roles,
	 	},
	}

	err = store.NewUser(ctx, &user)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
		"id":        *user.Id,
		"role_ids":      *user.RoleIds,
	})

	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	response := struct {
		User  models.User
		Token string
	}{
		User:  user,
		Token: jwtstring,
	}

	ServeJSON(w, response)
}` + `

func GetMyUserData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	user, err := store.GetUser(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = nil

	ServeJSON(w, user)
}`

	return pkg + handlers
}

func GenerateHandlers(s Struct) (string, error) {
	var out string

	struct_name_snake := ToSnakeCase(s.Name)
	struct_name := ToUpperCase(s.Name)

	gethandler := `

	func Get{{struct_name}}(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
	
		q := chi.URLParam(r, "{{struct_name_snake}}_id")
		id, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	
		data, err := store.Get{{struct_name}}(ctx, id)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}
	
		ServeJSON(w, data)
	}`

	gethandler = strings.ReplaceAll(gethandler, "{{struct_name_snake}}", struct_name_snake)
	gethandler = strings.ReplaceAll(gethandler, "{{struct_name}}", struct_name)

	updatehandler := `

	func Update{{struct_name}}(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		q := chi.URLParam(r, "{{struct_name_snake}}_id")
		id, err := strconv.ParseInt(q, 10, 64)

		data := models.{{struct_name}}{}
	
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&data)
		if err != nil {
			ServeError(w, err.Error(), 400)
			return
		}

		data.Id = &id
		
	
		err = store.Update{{struct_name}}(ctx, data)
		if err != nil {
			ServeError(w, err.Error(), 400)
			return
		}
	
		w.WriteHeader(200)
	}`

	updatehandler = strings.ReplaceAll(updatehandler, "{{struct_name}}", struct_name)

	deletehandler := `

	func Delete{{struct_name}}(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		q := chi.URLParam(r, "{{struct_name_snake}}_id")
		id, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		err = store.Delete{{struct_name}}(ctx, id)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		}`

	deletehandler = strings.ReplaceAll(deletehandler, "{{struct_name}}", struct_name)
	deletehandler = strings.ReplaceAll(deletehandler, "{{struct_name_snake}}", struct_name_snake)

	createhandler := `

	func New{{struct_name}}(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		input := models.{{struct_name}}{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&input)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.New{{struct_name}}(ctx, &input)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ServeJSON(w, input)
	}
	`

	createhandler = strings.ReplaceAll(createhandler, "{{struct_name}}", struct_name)

	out = gethandler + `
	
	` + createhandler + `
	
	` + updatehandler + `
	
	` + deletehandler + `
	
	`

	for _, c := range s.Columns {
		if c.ForeignKeyTable == "user" {
			// if c.IsUnique {
			//get single foreign object

			// 				getbyuserid := `

			// func Get{{struct_name_upper}}ByUserId(w http.ResponseWriter, r *http.Request) {
			// 	ctx := r.Context()
			// 	mid := ctx.Value(models.CTX_user_id).(int64)

			// 	data, err := store.Get{{struct_name_upper}}ByUserId(ctx, mid)
			// 	if err != nil {
			// 		ServeError(w, err.Error(), 400)
			// 		return
			// 	}

			// 	ServeJSON(w, data)

			// }`

			// 				out += getbyuserid
			// 				continue
			// }

			listbyuserid := `

				func List{{struct_name_upper}}ForUserBy{{foreign_column_name_upper}}(w http.ResponseWriter, r *http.Request) {
					ctx := r.Context()
					mid := ctx.Value(models.CTX_user_id).(int64)

					data, err := store.List{{struct_name_upper}}By{{column_name_upper}}(ctx, mid)
					if err != nil {
						ServeError(w, err.Error(), 400)
						return
					}


					ServeJSON(w, data)
	}`

			listbyuserid = strings.ReplaceAll(listbyuserid, "{{foreign_column_name_upper}}", ToUpperCase(c.ForeignKeyColumn))
			listbyuserid = strings.ReplaceAll(listbyuserid, "{{column_name_upper}}", ToUpperCase(c.Name))
			listbyuserid = strings.ReplaceAll(listbyuserid, "{{struct_name_upper}}", struct_name)

			out += listbyuserid
		}
	}

	return out, nil

}
