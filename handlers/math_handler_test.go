package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/handlers"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/routes"
)

func setupHandlerTestEnv() *gin.Engine {
	r := gin.Default()
	routes.SetRoutes(r)

	return r
}

func TestHandleRandomAdd(t *testing.T) {
	router := setupHandlerTestEnv()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/random/add", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, but got %d", w.Code)
	}

	var response map[string]handlers.CompResult
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	addResult, exists := response["addResult"]
	if !exists {
		t.Error("Expected addResult in response")
		return
	}

	if addResult.Result != addResult.InputComponents.ComponentOne+addResult.InputComponents.ComponentTwo {
		t.Errorf("Addition result is incorrect: got %d, expected %d", addResult.Result, addResult.InputComponents.ComponentOne+addResult.InputComponents.ComponentTwo)
	}
}

func TestHandleRandomSub(t *testing.T) {
	router := setupHandlerTestEnv()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/random/sub", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, but got %d", w.Code)
	}

	var response map[string]handlers.CompResult
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	subResult, exists := response["subResult"]
	if !exists {
		t.Error("Expected subResult in response")
		return
	}

	if subResult.Result != subResult.InputComponents.ComponentOne-subResult.InputComponents.ComponentTwo {
		t.Errorf("Subtraction result is incorrect: got %d, expected %d", subResult.Result, subResult.InputComponents.ComponentOne-subResult.InputComponents.ComponentTwo)
	}
}

func TestHandleAdd(t *testing.T) {
	router := setupHandlerTestEnv()

	body := `{
		"component_one": 5,
		"component_two": 10
	}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add", bytes.NewBufferString(body))
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, but got %d", w.Code)
	}

	var response map[string]handlers.CompResult
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	addResult, exists := response["addResult"]
	if !exists {
		t.Error("Expected addResult in response")
		return
	}

	if addResult.Result != 15 {
		t.Errorf("Addition result is incorrect: got %d, expected 15", addResult.Result)
	}
}

func TestHandleSub(t *testing.T) {
	router := setupHandlerTestEnv()

	body := `{
		"component_one": 10,
		"component_two": 5
	}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sub", bytes.NewBufferString(body))
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, but got %d", w.Code)
	}

	var response map[string]handlers.CompResult
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	subResult, exists := response["subResult"]
	if !exists {
		t.Error("Expected subResult in response")
		return
	}

	if subResult.Result != 5 {
		t.Errorf("Subtraction result is incorrect: got %d, expected 5", subResult.Result)
	}
}
