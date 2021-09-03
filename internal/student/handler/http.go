package handler

import (
	"github.com/anish-yadav/lms-api/internal/constants"
	webresponse "github.com/anish-yadav/lms-api/internal/pkg/webservice/response"
	"github.com/anish-yadav/lms-api/internal/student/manager"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type studentHttpHandler struct {
	manager manager.HttpManager
}

func NewHttpHandler(manager manager.HttpManager) *studentHttpHandler {
	return &studentHttpHandler{manager}
}

func (s *studentHttpHandler) HandleGetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := s.manager.GetStudentByID(id)
	if err != nil {
		if err.Error() == constants.StudentNotFound {
			webresponse.RespondWithError(w, http.StatusNotFound, constants.StudentNotFound)
			return
		}

		webresponse.RespondWithError(w, http.StatusInternalServerError, constants.IntervalServerError)
		return
	}

	var student StudentResponse
	bin, err := bson.Marshal(data)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusInternalServerError, constants.IntervalServerError)
		return
	}

	err = bson.Unmarshal(bin, &student)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusInternalServerError, constants.IntervalServerError)
		return
	}



}
