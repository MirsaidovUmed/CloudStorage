package router

import (
	"CloudStorage/internal/transport/http/handlers"
	"CloudStorage/internal/transport/http/middleware"
	"CloudStorage/pkg/http"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler, mw middleware.MiddlewareInterface) *mux.Router {
	router := http.NewRouter()

	router.HandleFunc("/api/registration", handlers.Registration).Methods("POST")
	return router
}
