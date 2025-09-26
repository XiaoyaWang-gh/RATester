package controllers

import (
	"fmt"
	"testing"
)

func TestGetAuthKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	authController := &AuthController{}
	authController.GetAuthKey()

	// Check the status code
	if authController.Data["json"].(map[string]interface{})["status"] != 1 {
		t.Errorf("Expected status to be 1, got %v", authController.Data["json"].(map[string]interface{})["status"])
	}

	// Check the crypt_auth_key
	if _, ok := authController.Data["json"].(map[string]interface{})["crypt_auth_key"]; !ok {
		t.Errorf("Expected crypt_auth_key to be present")
	}

	// Check the crypt_type
	if authController.Data["json"].(map[string]interface{})["crypt_type"] != "aes cbc" {
		t.Errorf("Expected crypt_type to be 'aes cbc', got %v", authController.Data["json"].(map[string]interface{})["crypt_type"])
	}
}
