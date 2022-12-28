package routes

import (
	"github.com/davidjss04/api/pkg/controllers"
	"github.com/go-chi/chi/v5"
)

var RegisterMailRoutes = func(route chi.Router) {
	route.Get("/mail/{mailId}", controllers.GetMailById)
	route.Get("/mails/{phrase}", controllers.SearchMail)
}
