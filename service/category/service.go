package category

import (
	"github.com/gorilla/mux"
	"github.com/pimp13/go-react-project/types"
	"github.com/pimp13/go-react-project/utils"
	"net/http"
)

type Handler struct {
	store types.CategoryStore
}

func NewHandler(store types.CategoryStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/categories", h.HandleGetAll).Methods("GET")
	router.HandleFunc("/categories", h.HandleCreate).Methods("POST")
}

func (h *Handler) HandleGetAll(res http.ResponseWriter, req *http.Request) {
	panic("get all categories handler")
}
func (h *Handler) HandleCreate(res http.ResponseWriter, req *http.Request) {
	// Get payload and parse json
	var payload *types.CreateCategoryPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	// Validation payload
}
