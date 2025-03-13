package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pimp13/go-react-project/config"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	// Get user payload and parse json payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	// Validation the payload
	if err := utils.Validate.Struct(&payload); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			utils.WriteError(res, http.StatusBadRequest, validationErrors)
		} else {
			utils.WriteError(res, http.StatusInternalServerError, err)
		}
		return
	}

	// Check user by email exists
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(res, http.StatusNotFound, err)
		return
	}

	// Check user password is match
	if !auth.ComparePassword(u.Password, []byte(payload.Password)) {
		utils.WriteError(res, http.StatusNotFound, fmt.Errorf("not found, email or password invalid"))
		return
	}

	// Login user and create jwt token
	secret := []byte(config.Envs.JWTKey)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	// Return response json and token JWT
	if err := utils.WriteJSON(res, http.StatusOK, map[string]string{
		"token":   token,
		"message": "User Logged successfully!",
	}); err != nil {
		log.Fatal(err)
		return
	}
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
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			utils.WriteError(res, http.StatusBadRequest, validationErrors)
		} else {
			utils.WriteError(res, http.StatusInternalServerError, err)
		}
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
		ID:        uuid.New(),
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
