package router

import (
	"CloudStorage/internal/transport/http/handlers"
	"CloudStorage/internal/transport/http/middleware"
	"CloudStorage/pkg/http"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler, mw middleware.MiddlewareInterface) *mux.Router {
	router := http.NewRouter()
	privateRouter := router.NewRoute().Subrouter()
	privateRouter.Use(mw.JWT)

	router.HandleFunc("/api/registration", handlers.Registration).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	privateRouter.HandleFunc("/user/list", handlers.GetUserList).Methods("GET")
	privateRouter.HandleFunc("/user/get/{id}", handlers.GetUserByID).Methods("GET")
	privateRouter.HandleFunc("/user/update", handlers.UpdateUser).Methods("PUT")
	privateRouter.HandleFunc("/admin/users/delete/{id}", handlers.DeleteUser).Methods("DELETE")

	privateRouter.HandleFunc("/files/upload", handlers.UploadFile).Methods("POST")
	privateRouter.HandleFunc("/files/list", handlers.GetFileList).Methods("GET")
	return router
}
