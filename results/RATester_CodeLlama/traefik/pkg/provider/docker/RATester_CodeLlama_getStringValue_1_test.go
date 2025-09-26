package docker

import (
	"fmt"
	"testing"
)

func TestGetStringValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	labels := map[string]string{"label1": "value1"}
	labelName := "label1"
	defaultValue := "defaultValue"
	expected := "value1"
	actual := getStringValue(labels, labelName, defaultValue)
	if actual != expected {
		t.Errorf("Expected %s, but actual is %s", expected, actual)
	}
}
