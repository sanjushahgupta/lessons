package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sanjushahgupta/lessons/sample-rest-di/domain"
)

// SetupUserRoutes sets up routing for user service
func (r *REST) SetupUserRoutes(router *httprouter.Router) *httprouter.Router {
	router.GET("/user", r.GetAllUser)
	router.GET("/user/:id", r.GetUser)
	router.POST("/user", r.CreateUser)

	return router
}

// CreateUser creates user
// if required fields are not fulfilled, 401 response is returned
func (r *REST) CreateUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Body == nil {
		respond(w, http.StatusBadRequest, response{
			Error:   errNilBody,
			Message: msgNilBody,
		})
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respond(w, http.StatusInternalServerError, response{
			Error:   err,
			Message: err.Error(),
		})
		return
	}
	defer req.Body.Close()

	var user domain.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		respond(w, http.StatusInternalServerError, response{
			Error:   err,
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		respond(w, http.StatusBadRequest, response{
			Error:   err,
			Message: err.Error(),
		})
		return
	}

	id, err := r.logic.CreateUser(user)
	if err != nil {
		respond(w, http.StatusBadRequest, response{
			Error:   err,
			Message: err.Error(),
		})
		return
	}

	respond(w, http.StatusOK, response{
		Message: msgUserCreated(id),
	})
}

// GetAllUser gets all user
func (r *REST) GetAllUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	users, err := r.logic.GetAllUser()
	if err != nil {
		respond(w, http.StatusInternalServerError, response{
			Error:   err,
			Message: err.Error(),
		})
		return
	}

	respond(w, http.StatusOK, response{
		Data: users,
	})
}

// GetUser gets user
func (r *REST) GetUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	strid := ps.ByName("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		respond(w, http.StatusInternalServerError, response{
			Error:   err,
			Message: err.Error(),
		})
		return
	}

	// TODO if user does not exist, it currently returns 500, but should return 200
	// DETAILS returns 200 when User does not exist
	user, err := r.logic.GetUser(id)
	if err != nil {
		respond(w, http.StatusInternalServerError, response{
			Error:   err,
			Message: err.Error(),
		})
		return
	}

	respond(w, http.StatusOK, response{
		Data: user,
	})
}
