package fiber

import (
	"fmt"
	"testing"
)

func TestExecuteOnGroupHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &Hooks{}
	group := &Group{}

	// Act
	err := h.executeOnGroupHooks(*group)

	// Assert
	if err != nil {
		t.Errorf("executeOnGroupHooks() error = %v", err)
		return
	}
}
