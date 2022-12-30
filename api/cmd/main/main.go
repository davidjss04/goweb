package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/davidjss04/api/pkg/routes"
)

func main() {
	port := "3000"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))
	r.Group(routes.RegisterMailRoutes)
	http.Handle("/", r)
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

	// utils.CreateJson("/home/davidjss/workspace/zincsearch/enron_mail_20110402/maildir", "export.json")
	// utils.LoadMails("/home/davidjss/workspace/zincsearch/indexer/tree", "export.json")
	// utils.LoadMails("/home/davidjss/workspace/zincsearch/enron_mail_20110402/maildir", "export.json")
}
