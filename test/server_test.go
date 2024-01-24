package app_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"p3ld3v.dev/template/app"
	"p3ld3v.dev/template/app/services"
	"p3ld3v.dev/template/app/services/db/sqlc"
	"p3ld3v.dev/template/test/mocks"
)

type MockDependencies struct {
	DbStore *mocks.DbStore
	Logger  *mocks.Logger
}

func TestGetUser(t *testing.T) {
	dbMocks := &mocks.DbStore{}
	depMocks := &app.Dependencies{
		Logger:      &mocks.Logger{},
		DbService:   dbMocks,
		UserService: services.NewUserService(dbMocks, &mocks.Logger{}),
	}
	var id int64 = 1
	dbMocks.On("GetUser", id).Return(&sqlc.User{
		ID:   1,
		Name: "test",
	}, nil)
	router := app.SetupRouter(depMocks)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Code, http.StatusOK)
	assert.Equal(t, recorder.Body.String(), `{"id":1,"name":"test"}`)
}

func TestPostUser(t *testing.T) {
	dbMocks := &mocks.DbStore{}
	depMocks := &app.Dependencies{
		Logger:      &mocks.Logger{},
		DbService:   dbMocks,
		UserService: services.NewUserService(dbMocks, &mocks.Logger{}),
	}
	dbMocks.On("CreateUser", "test").Return(&sqlc.User{
		ID:   1,
		Name: "test",
	}, nil)
	router := app.SetupRouter(depMocks)
	recorder := httptest.NewRecorder()
	// Create a JSON payload for the POST request
	jsonPayload := []byte(`{"name": "test"}`)

	request, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Serve the HTTP request
	router.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Code, http.StatusOK)
	assert.Equal(t, recorder.Body.String(), `{"id":1,"name":"test"}`)
}
