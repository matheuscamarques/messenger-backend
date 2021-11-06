package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/matheuscamarques/messenger-backend/src/models"
)

type UserController struct {
	Model *models.User
}

var userController UserController = UserController{
	Model: &models.User{},
}

func (controller *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := controller.Model.GetByID(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// transform user in json
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
