package webservice

import (
	"github.com/anish-yadav/lms-api/internal/router"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartServer(port string) {
	r := router.NewRouter()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Infof("webservice.StartServer: server starting at port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Debugf("webservice.StartServer: %s", err.Error())
	}
}
