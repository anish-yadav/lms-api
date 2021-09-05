package handler

import (
	"encoding/json"
	"github.com/anish-yadav/lms-api/internal/constants"
	webresponse "github.com/anish-yadav/lms-api/internal/pkg/webservice/response"
	"github.com/anish-yadav/lms-api/internal/user/usecase"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

type userHttpHandler struct {
	manager usecase.HttpManager
}

func NewHttpHandler(manager usecase.HttpManager) *userHttpHandler {
	return &userHttpHandler{manager}
}

func (u *userHttpHandler) HandlePostStudent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	var userRequest UserRequest
	err = json.Unmarshal(data, &userRequest)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(userRequest)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	id, err := u.manager.CreateUser(userRequest.toUser())
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			webresponse.RespondWithError(w, http.StatusConflict, constants.Conflict)
			return
		}
		webresponse.RespondWithError(w, http.StatusInternalServerError, constants.IntervalServerError)
		return
	}
	response := UserRequestResponse{id}
	webresponse.RespondWithSuccess(w, http.StatusCreated, response)
	return
}

func (u *userHttpHandler) HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	var userRequest ResetPasswordRequest
	err = json.Unmarshal(data, &userRequest)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(userRequest)
	if err != nil {
		webresponse.RespondWithError(w, http.StatusBadRequest, constants.BadRequest)
		return
	}
	err = u.manager.ResetPassword(userRequest.ID, userRequest.OldPassword, userRequest.NewPassword)
	if err != nil {
		if err.Error() == constants.ItemNotFound {
			webresponse.RespondWithError(w, http.StatusNotFound, constants.ItemNotFound)
			return
		}
		webresponse.RespondWithError(w, http.StatusInternalServerError, constants.IntervalServerError)
		return
	}
	response := UserRequestResponse{userRequest.ID}
	webresponse.RespondWithSuccess(w, http.StatusCreated, response)
	return
}
