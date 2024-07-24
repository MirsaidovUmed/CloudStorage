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
	privateRouter.HandleFunc("/admin/users/list", handlers.AdminGetUserList).Methods("GET")
	privateRouter.HandleFunc("/admin/users/update/{id}", handlers.AdminUpdateUserById).Methods("PUT")
	privateRouter.HandleFunc("/admin/users/get/{id}", handlers.AdminGetUserByID).Methods("GET")

	privateRouter.HandleFunc("/files/upload", handlers.UploadFile).Methods("POST")
	privateRouter.HandleFunc("/files/list", handlers.GetFileList).Methods("GET")
	privateRouter.HandleFunc("/files/get/{id}", handlers.GetFileById).Methods("GET")
	privateRouter.HandleFunc("/files/remove/{id}", handlers.DeleteFile).Methods("DELETE")
	privateRouter.HandleFunc("/files/rename/{id}", handlers.RenameFile).Methods("PUT")

	privateRouter.HandleFunc("/directories/create", handlers.CreateDirectory).Methods("POST")
	privateRouter.HandleFunc("/directories/rename", handlers.RenameDirectory).Methods("POST")
	privateRouter.HandleFunc("/directories/get/{id}", handlers.GetDirectoryById).Methods("Get")
	privateRouter.HandleFunc("/directories/delete/{id}", handlers.DeleteDirectory).Methods("DELETE")

	privateRouter.HandleFunc("/files/share/{id}", handlers.GetFileAccessUsers).Methods("GET")
	privateRouter.HandleFunc("/files/share/{id}/{user_id}", handlers.ShareFile).Methods("PUT")
	return router
}
