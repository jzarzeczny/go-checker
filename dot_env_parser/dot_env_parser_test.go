package dotenvparser

import (
	"os"
	"reflect"
	"testing"
)

func TestParseEnvVariables(t *testing.T) {
	// Create a temporary .env file for testing
	envContent := []byte("KEY1=value1\nKEY2=value2\n")
	err := os.WriteFile(".env", envContent, 0644)
	if err != nil {
		t.Fatal("Failed to create .env file:", err)
	}
	defer os.Remove(".env")

	// Test parsing the .env file
	expectedEnv := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
	}
	envVariables, err := ParseEnvVariables()
	if err != nil {
		t.Errorf("Error parsing .env file: %v", err)
	}
	if !reflect.DeepEqual(envVariables, expectedEnv) {
		t.Errorf("Parsed environment variables do not match expected:\nGot: %v\nExpected: %v", envVariables, expectedEnv)
	}
}

func TestParseEnvVariables_FileNotFound(t *testing.T) {
	// Test handling when .env file is not found
	_, err := ParseEnvVariables()
	if err == nil {
		t.Error(".env file not found error not returned")
	}
}

func TestParseEnvVariables_ReadError(t *testing.T) {
	// Create a temporary .env file
	envContent := []byte("KEY1=value1\nKEY2=value2\n")
	err := os.WriteFile(".env", envContent, 0644)
	if err != nil {
		t.Fatal("Failed to create .env file:", err)
	}
	defer os.Remove(".env")

	// Make the .env file unreadable
	err = os.Chmod(".env", 0000)
	if err != nil {
		t.Fatal("Failed to make .env file unreadable:", err)
	}
	defer os.Chmod(".env", 0644)

	// Test handling when there is an error reading the .env file
	_, err = ParseEnvVariables()
	if err == nil {
		t.Error("No error returned for .env file with read error")
	}
}

// Add more test cases as needed to cover different scenarios
