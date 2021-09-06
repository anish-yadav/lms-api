package router

import (
	"github.com/anish-yadav/lms-api/internal/auth"
	studentHandler "github.com/anish-yadav/lms-api/internal/student/handler"
	studentManager "github.com/anish-yadav/lms-api/internal/student/manager"
	userHttpHandler "github.com/anish-yadav/lms-api/internal/user/handler"
	userUsecase "github.com/anish-yadav/lms-api/internal/user/usecase"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	studentHttpManager := studentManager.NewHttpManager()
	studentHttpHandler := studentHandler.NewHttpHandler(studentHttpManager)

	userManager := userUsecase.NewHttpManager()
	userHandler := userHttpHandler.NewHttpHandler(userManager)
	// routes
	v1 := router.PathPrefix("/api/v1").Subrouter()
	// jwt
	v1.Use(auth.Middleware)
	gets := v1.Methods(http.MethodGet).Subrouter()
	posts := v1.Methods(http.MethodPost).Subrouter()
	del := v1.Methods(http.MethodDelete).Subrouter()

	gets.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("success")); err != nil {
			log.Errorf("failed to send health reposnse")
		}
	})

	gets.HandleFunc("/students/{id}", studentHttpHandler.HandleGetStudent)

	posts.HandleFunc("/users", userHandler.HandlePostStudent)
	posts.HandleFunc("/users/password-reset", userHandler.HandleChangePassword)

	del.HandleFunc("/users/{id}", userHandler.HandleUserDelete)
	return router
}
