package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	"github.com/davidjss04/api/pkg/routes"
)

func main() {
	port := "3000"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Group(routes.RegisterMailRoutes)
	http.Handle("/", r)
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

	// utils.CreateJson("/home/davidjss/workspace/zincsearch/enron_mail_20110402/maildir", "export.json")
	// utils.LoadMails("/home/davidjss/workspace/zincsearch/indexer/tree", "export.json")
	// utils.LoadMails("/home/davidjss/workspace/zincsearch/enron_mail_20110402/maildir", "export.json")
}
