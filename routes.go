package main

import (
	"strings"
)

func GenerateRoutes(sts []Struct) string {
	var out string

	var user_routes string

	out = `
	
	r := chi.NewRouter()

	corsParams := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsParams.Handler)
	r.Use(middleware.Logger)
	r.Use(handlers.Authenticator)
	
	`

	for _, s := range sts {
		q := `
		
	r.Route("/{{struct_name_snake}}", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.New{{struct_name_upper}})
			r.Route("/{{{struct_name_snake}}_id}", func(r chi.Router) {
				r.Get("/", handlers.Get{{struct_name_upper}})
				r.Put("/", handlers.Update{{struct_name_upper}})
				r.Delete("/", handlers.Delete{{struct_name_upper}})
			})
		})
	})`

		q = strings.ReplaceAll(q, "{{struct_name_upper}}", ToUpperCase(s.Name))
		q = strings.ReplaceAll(q, "{{struct_name_snake}}", ToSnakeCase(s.Name))
		out += q

		for _, c := range s.Columns {

			if IsUserOwned(c) {
				if c.IsUnique {
					continue
				}

				ur := `
r.Get("/{{struct_name_snake}}", handlers.List{{struct_name_upper}}ForUserBy{{foreign_column_name_upper}})`

				ur = strings.ReplaceAll(ur, "{{foreign_column_name_upper}}", ToUpperCase(c.ForeignKeyColumn))
				ur = strings.ReplaceAll(ur, "{{struct_name_upper}}", ToUpperCase(s.Name))
				ur = strings.ReplaceAll(ur, "{{struct_name_snake}}", ToSnakeCase(s.Name))

				user_routes += ur
			}
		}
	}

	out += `

	r.Route("/me", func(r chi.Router) {
		r.Use(handlers.RestrictAuth)
		` + user_routes + `
	})`

	return out
}

func IsUserOwned(c Column) bool {

	if c.ForeignKeyTable == "user" && c.IsForeignKey {
		return true
	}

	return false
}
