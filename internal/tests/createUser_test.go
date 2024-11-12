// package endpoint_test

// import (
// 	"CaseGo/internal/endpoint"
// 	"CaseGo/internal/models"
// 	"CaseGo/internal/service"
// 	"encoding/json"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"gorm.io/gorm"
// )

// // MockService is a mock implementation of the Service interface
// type MockService struct {
// 	mock.Mock
// }

// // CreateUser implements endpoint.Service.
// func (m *MockService) CreateUser(int, string) models.UserModel {
// 	panic("unimplemented")
// }

// // GetCases implements endpoint.Service.
// func (m *MockService) GetCases() []service.Cases {
// 	panic("unimplemented")
// }

// // GetInventory implements endpoint.Service.
// func (m *MockService) GetInventory(int) []service.Inventory {
// 	panic("unimplemented")
// }

// // GetWeapons implements endpoint.Service.
// func (m *MockService) GetWeapons(int) []service.Weapons {
// 	panic("unimplemented")
// }

// func (m *MockService) GetUsers(id int) models.UserModel {
// 	args := m.Called(id)
// 	return args.Get(0).(models.UserModel)
// }

// func TestGetUsers(t *testing.T) {
// 	app := fiber.New()

// 	mockService := new(MockService)
// 	e := endpoint.New(mockService)
// 	app.Post("/getUser", e.GetUsers)

// 	// Define the expected user model
// 	expectedUser := models.UserModel{
// 		ID:         1,
// 		TelegramID: 19234,
// 		Name:       "igor",
// 		Coins:      100,
// 		CreatedAt:  time.Date(2024, 11, 12, 10, 32, 9, 519838000, time.Local),
// 		UpdatedAt:  time.Date(2024, 11, 12, 10, 32, 9, 519838000, time.Local),
// 		DeletedAt:  gorm.DeletedAt{},
// 	}

// 	// Set up the mock to return the expected user
// 	mockService.On("GetUsers", 19234).Return(expectedUser)

// 	userRequest := endpoint.UserRequest{
// 		Id: 1924,
// 	}
// 	requestBody, _ := json.Marshal(userRequest)
// 	req := httptest.NewRequest("POST", "/getUser", strings.NewReader(string(requestBody)))
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, _ := app.Test(req)

// 	// Check the response
// 	var response struct {
// 		Data models.UserModel `json:"data"`
// 	}
// 	err := json.NewDecoder(resp.Body).Decode(&response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUser, response.Data)
// }
