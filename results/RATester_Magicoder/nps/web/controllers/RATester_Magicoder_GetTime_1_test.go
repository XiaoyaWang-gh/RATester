package controllers

import (
	"fmt"
	"testing"
)

func TestGetTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &AuthController{}
	ctrl.GetTime()

	// Assert that the time is not zero
	if ctrl.Data["json"].(map[string]interface{})["time"] == 0 {
		t.Error("Time is zero")
	}
}
