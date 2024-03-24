package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jzarzeczny/go-checker/interfaces"
)

func TestGetWebsiteStatus(t *testing.T) {
	// Define sample URL list
	urlList := []interfaces.URLData{
		{Name: "Google", URL: "https://www.google.com"},
		{Name: "GitHub", URL: "https://www.github.com"},
	}

	token := "TEST_TOKEN"

	// Create a request with a dummy request body (not used in this handler)
	req, err := http.NewRequest("GET", "/status", bytes.NewBufferString(""))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the sample URL list
	GetWebsiteStatus(rr, req, urlList, token)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the Content-Type header
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Add more assertions as needed to validate the response body
	// For example, you can parse the JSON response body and check its structure or content
	// Example: Check if the response body contains valid JSON data
	// var result []interfaces.Result
	// err = json.Unmarshal(rr.Body.Bytes(), &result)
	// if err != nil {
	// 	t.Errorf("error unmarshalling JSON: %v", err)
	// }
}
