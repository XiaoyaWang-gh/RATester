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

	ctrl := &AuthController{}
	ctrl.GetAuthKey()

	// Assert that the status is 1
	if ctrl.Data["json"].(map[string]interface{})["status"] != 1 {
		t.Errorf("Expected status to be 1, but got %v", ctrl.Data["json"].(map[string]interface{})["status"])
	}

	// Assert that the crypt_auth_key is not empty
	if ctrl.Data["json"].(map[string]interface{})["crypt_auth_key"] == "" {
		t.Errorf("Expected crypt_auth_key to be not empty, but got empty")
	}

	// Assert that the crypt_type is "aes cbc"
	if ctrl.Data["json"].(map[string]interface{})["crypt_type"] != "aes cbc" {
		t.Errorf("Expected crypt_type to be 'aes cbc', but got %v", ctrl.Data["json"].(map[string]interface{})["crypt_type"])
	}
}
