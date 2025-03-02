package handler

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct{}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"users/register.html"})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"users/login.html"})
}

func (h *UserHandler) PostRegister(w http.ResponseWriter, r *http.Request) {}

func (h *UserHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest = &LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	json.NewEncoder(w).Encode(&loginRequest)
}
