package controllers

import (
	"fmt"
	"testing"
)

func TestIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &LoginController{}
	ctrl.Index()

	// Check if the TplName is set correctly
	if ctrl.TplName != "login/index.html" {
		t.Errorf("Expected TplName to be 'login/index.html', got '%s'", ctrl.TplName)
	}
}
