package proxy

import (
	"fmt"
	"testing"
)

func TestResetReqMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if resetReqMethod("ET") != "GET" {
		t.Errorf("resetReqMethod(\"ET\") = %v, want \"GET\"", resetReqMethod("ET"))
	}
	if resetReqMethod("OST") != "POST" {
		t.Errorf("resetReqMethod(\"OST\") = %v, want \"POST\"", resetReqMethod("OST"))
	}
	if resetReqMethod("GET") != "GET" {
		t.Errorf("resetReqMethod(\"GET\") = %v, want \"GET\"", resetReqMethod("GET"))
	}
	if resetReqMethod("POST") != "POST" {
		t.Errorf("resetReqMethod(\"POST\") = %v, want \"POST\"", resetReqMethod("POST"))
	}
}
