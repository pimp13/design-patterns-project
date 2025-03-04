package user

import (
	"github.com/gorilla/mux"
	"github.com/pimp13/go-react-project/types"
	"github.com/pimp13/go-react-project/utils"
	"log"
	"net/http"
)

type Handler struct {
	// Dependencies
	store types.UserStore
}

// Constructor
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// Register the user routes
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
	// Check if user is already registered
	// ...

	// Hash the password
	// ...

	// Create new user in database
	// ...

	// Return JSON response
	response := map[string]string{"message": "User registered successfully"}
	if err := utils.WriteJSON(res, http.StatusCreated, response); err != nil {
		log.Fatal(err)
	}
}
