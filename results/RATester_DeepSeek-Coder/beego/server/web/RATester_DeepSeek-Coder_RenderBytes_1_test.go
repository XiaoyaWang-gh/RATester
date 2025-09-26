package web

import (
	"fmt"
	"testing"
)

func TestRenderBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Data:   make(map[interface{}]interface{}),
		Layout: "layout",
		LayoutSections: map[string]string{
			"section1": "section1",
			"section2": "section2",
		},
	}

	ctrl.Data["key1"] = "value1"
	ctrl.Data["key2"] = "value2"

	buf, err := ctrl.RenderBytes()
	if err != nil {
		t.Errorf("Error in RenderBytes: %v", err)
	}

	if len(buf) == 0 {
		t.Errorf("Expected non-empty buffer, got empty")
	}

	expected := "value1"
	actual := ctrl.Data["LayoutContent"]
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = "value1"
	actual = ctrl.Data["section1"]
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = "value2"
	actual = ctrl.Data["section2"]
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
