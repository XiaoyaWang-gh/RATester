package docker

import (
	"fmt"
	"testing"
)

func TestGetStringMultipleStrict_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	labels := map[string]string{
		"label1": "value1",
		"label2": "value2",
		"label3": "value3",
	}
	labelNames := []string{"label1", "label2", "label3"}
	foundLabels, err := getStringMultipleStrict(labels, labelNames...)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if len(foundLabels) != 3 {
		t.Errorf("foundLabels length is not 3")
	}
	if foundLabels["label1"] != "value1" {
		t.Errorf("foundLabels[label1] is not value1")
	}
	if foundLabels["label2"] != "value2" {
		t.Errorf("foundLabels[label2] is not value2")
	}
	if foundLabels["label3"] != "value3" {
		t.Errorf("foundLabels[label3] is not value3")
	}
}
