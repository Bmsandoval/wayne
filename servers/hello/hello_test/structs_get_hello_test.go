package hello_test

import "github.com/bmsandoval/wayne/db/models"

type GetHelloTestData struct {
	BaseTestData
	MockGetHello *MockGetHello
}

type MockGetHello struct {
	OutGreetings []models.Greetings
	OutError   error
}
