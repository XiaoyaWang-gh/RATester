package crd

import (
	"fmt"
	"testing"
)

func TestMakeServiceKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rule := "rule"
	ingressName := "ingressName"
	key, err := makeServiceKey(rule, ingressName)
	if err != nil {
		t.Errorf("makeServiceKey() error = %v", err)
		return
	}
	if key != "ingressName-rule-%.10x" {
		t.Errorf("makeServiceKey() key = %v, want %v", key, "ingressName-rule-%.10x")
	}
}
