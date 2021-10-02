package handler

import (
	"encoding/json"
	"github.com/anish-yadav/lms-api/internal/class/usecase"
	"github.com/anish-yadav/lms-api/internal/constants"
	webresponse "github.com/anish-yadav/lms-api/internal/pkg/webservice/response"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

type classHttpHandler struct {
	manager usecase.HttpManager
}

func NewHttpHandler(manager usecase.HttpManager) *classHttpHandler {
	return &classHttpHandler{manager}
}

func (u *classHttpHandler) HandleCreateClass(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	var classRequest CreateClassRequest
	err = json.Unmarshal(data, &classRequest)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(classRequest)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	id, err := u.manager.CreateClass(classRequest.toClass())
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			webresponse.RespondWithError(w, http.StatusConflict, constants.Conflict)
			return
		}
		webresponse.RespondWithError(w, http.StatusInternalServerError, constants.IntervalServerError)
		return
	}
	response := ClassRequestResponse{id}
	webresponse.RespondWithSuccess(w, http.StatusCreated, response)
	return
}

func (u *classHttpHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {

	classes, err := u.manager.GetAllClass()
	if err != nil {
		if err.Error() == constants.ItemNotFound {
			webresponse.RespondWithError(w, http.StatusNotFound, constants.ItemNotFound)
			return
		}
		webresponse.RespondWithError(w, http.StatusInternalServerError, constants.IntervalServerError)
		return
	}

	response := &ClassResponse{classes}
	webresponse.RespondWithSuccess(w, http.StatusOK, response.classes)
	return
}
