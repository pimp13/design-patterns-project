package user

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"github.com/pimp13/go-react-project/types"
	// "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// * One test
/*
type MockUserStore struct {
	users map[string]*types.User
}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	if user, exists := m.users[email]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (m *MockUserStore) CreateUser(user *types.User) error {
	if _, exists := m.users[user.Email]; exists {
		return fmt.Errorf("user already exists")
	}
	m.users[user.Email] = user
	return nil
}

func (m *MockUserStore) GetUserByID(id uint) (*types.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func TestHandleRegister(t *testing.T) {
	mockStore := &MockUserStore{users: make(map[string]*types.User)}
	handler := NewHandler(mockStore)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	payload := types.RegisterUserPayload{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "User registered successfully", response["message"])
}
*/
// * Tow test

type MockUserStore struct{}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}
func (m *MockUserStore) GetUserByID(id uint) (*types.User, error) {
	return nil, nil
}
func (m *MockUserStore) CreateUser(user *types.User) error {
	return nil
}

func TestUserServiceHandler(t *testing.T) {
	userStore := &MockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user the payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John2",
			LastName:  "Doe2",
			Email:     "john.doe2@example.com",
			Password:  "password123",
		}
		body, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expexted status code %d, got %d", http.StatusBadRequest, recorder.Code)
		}

	})
}
