package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/pimp13/go-react-project/service/auth"
	"github.com/pimp13/go-react-project/types"
	"github.com/pimp13/go-react-project/utils"
)

type Handler struct {
	// Dependencies
	store types.UserStore
}

// NewHandler Constructor
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes Register the user routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// Handlers
func (h *Handler) handleLogin(res http.ResponseWriter, req *http.Request) {
	//
}

func (h *Handler) handleRegister(res http.ResponseWriter, req *http.Request) {
	// Get JSON response
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	// Validate the payload data
	if err := utils.Validate.Struct(&payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(res, http.StatusBadRequest, errors)
		return
	}

	// Check if user is already registered
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		// if existing error user with email input is existing and registered
		utils.WriteError(res, http.StatusConflict,
			fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	// Create new user in database
	err = h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	// Return JSON response
	response := map[string]string{"message": "User registered successfully"}
	if err := utils.WriteJSON(res, http.StatusCreated, response); err != nil {
		log.Fatal(err)
	}
}
