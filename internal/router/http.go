package router

import (
	"github.com/anish-yadav/lms-api/internal/auth"
	studentHandler "github.com/anish-yadav/lms-api/internal/student/handler"
	studentManager "github.com/anish-yadav/lms-api/internal/student/manager"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	studentHttpManager := studentManager.NewHttpManager()
	studentHttpHandler := studentHandler.NewHttpHandler(studentHttpManager)
	// routes
	v1 := router.PathPrefix("/api/v1").Subrouter()
	// jwt
	v1.Use(auth.Middleware)

	v1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("success")); err != nil {
			log.Errorf("failed to send health reposnse")
		}
	})

	v1.HandleFunc("/students/{id}", studentHttpHandler.HandleGetStudent)
	return router
}
