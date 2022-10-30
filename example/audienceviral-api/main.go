package main

import (
	"audienceviral-api/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {

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

	r.Route("/subscription", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewSubscription)
			r.Route("/{subscription_id}", func(r chi.Router) {
				r.Get("/", handlers.GetSubscription)
				r.Put("/", handlers.UpdateSubscription)
				r.Delete("/", handlers.DeleteSubscription)
			})
		})
	})

	r.Route("/user_finance", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewUserFinance)
			r.Route("/{user_finance_id}", func(r chi.Router) {
				r.Get("/", handlers.GetUserFinance)
				r.Put("/", handlers.UpdateUserFinance)
				r.Delete("/", handlers.DeleteUserFinance)
			})
		})
	})

	r.Route("/invoice", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewInvoice)
			r.Route("/{invoice_id}", func(r chi.Router) {
				r.Get("/", handlers.GetInvoice)
				r.Put("/", handlers.UpdateInvoice)
				r.Delete("/", handlers.DeleteInvoice)
			})
		})
	})

	r.Route("/user_search_history", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewUserSearchHistory)
			r.Route("/{user_search_history_id}", func(r chi.Router) {
				r.Get("/", handlers.GetUserSearchHistory)
				r.Put("/", handlers.UpdateUserSearchHistory)
				r.Delete("/", handlers.DeleteUserSearchHistory)
			})
		})
	})

	r.Route("/payment_provider_source", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewPaymentProviderSource)
			r.Route("/{payment_provider_source_id}", func(r chi.Router) {
				r.Get("/", handlers.GetPaymentProviderSource)
				r.Put("/", handlers.UpdatePaymentProviderSource)
				r.Delete("/", handlers.DeletePaymentProviderSource)
			})
		})
	})

	r.Route("/lead", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewLead)
			r.Route("/{lead_id}", func(r chi.Router) {
				r.Get("/", handlers.GetLead)
				r.Put("/", handlers.UpdateLead)
				r.Delete("/", handlers.DeleteLead)
			})
		})
	})

	r.Route("/payment_provider_customer", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewPaymentProviderCustomer)
			r.Route("/{payment_provider_customer_id}", func(r chi.Router) {
				r.Get("/", handlers.GetPaymentProviderCustomer)
				r.Put("/", handlers.UpdatePaymentProviderCustomer)
				r.Delete("/", handlers.DeletePaymentProviderCustomer)
			})
		})
	})

	r.Route("/user", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewUser)
			r.Route("/{user_id}", func(r chi.Router) {
				r.Get("/", handlers.GetUser)
				r.Put("/", handlers.UpdateUser)
				r.Delete("/", handlers.DeleteUser)
			})
		})
	})

	r.Route("/order", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewOrder)
			r.Route("/{order_id}", func(r chi.Router) {
				r.Get("/", handlers.GetOrder)
				r.Put("/", handlers.UpdateOrder)
				r.Delete("/", handlers.DeleteOrder)
			})
		})
	})

	r.Route("/payment_provider", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewPaymentProvider)
			r.Route("/{payment_provider_id}", func(r chi.Router) {
				r.Get("/", handlers.GetPaymentProvider)
				r.Put("/", handlers.UpdatePaymentProvider)
				r.Delete("/", handlers.DeletePaymentProvider)
			})
		})
	})

	r.Route("/order_credits", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewOrderCredits)
			r.Route("/{order_credits_id}", func(r chi.Router) {
				r.Get("/", handlers.GetOrderCredits)
				r.Put("/", handlers.UpdateOrderCredits)
				r.Delete("/", handlers.DeleteOrderCredits)
			})
		})
	})

	r.Route("/lead_instagram", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewLeadInstagram)
			r.Route("/{lead_instagram_id}", func(r chi.Router) {
				r.Get("/", handlers.GetLeadInstagram)
				r.Put("/", handlers.UpdateLeadInstagram)
				r.Delete("/", handlers.DeleteLeadInstagram)
			})
		})
	})

	r.Route("/data_platform", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewDataPlatform)
			r.Route("/{data_platform_id}", func(r chi.Router) {
				r.Get("/", handlers.GetDataPlatform)
				r.Put("/", handlers.UpdateDataPlatform)
				r.Delete("/", handlers.DeleteDataPlatform)
			})
		})
	})

	r.Route("/me", func(r chi.Router) {
		r.Use(handlers.RestrictAuth)

		r.Get("/subscription", handlers.ListSubscriptionForUserById)
		r.Get("/invoice", handlers.ListInvoiceForUserById)
		r.Get("/user_search_history", handlers.ListUserSearchHistoryForUserById)
		r.Get("/payment_provider_source", handlers.ListPaymentProviderSourceForUserById)
		r.Get("/payment_provider_customer", handlers.ListPaymentProviderCustomerForUserById)
		r.Get("/user", handlers.ListUserForUserById)
		r.Get("/order", handlers.ListOrderForUserById)
		r.Get("/order_credits", handlers.ListOrderCreditsForUserById)
	})

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		log.Fatalf("Error serving HTTP handlers: %v", err)
	}

}
