package fiber

import (
	"fmt"
	"testing"
)

func TestAcquireRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	redirect := AcquireRedirect()

	// Act
	redirect.Back()
	redirect.Message("key")
	redirect.Messages()
	redirect.OldInput("key")
	redirect.OldInputs()
	redirect.Route("name")
	redirect.Status(302)
	redirect.To("location")
	redirect.With("key", "value")
	redirect.WithInput()
	redirect.parseAndClearFlashMessages()
	redirect.processFlashMessages()
	redirect.release()

	// Assert
	if redirect.c == nil {
		t.Errorf("Expected redirect.c to be not nil")
	}
	if redirect.messages == nil {
		t.Errorf("Expected redirect.messages to be not nil")
	}
	if redirect.status != 302 {
		t.Errorf("Expected redirect.status to be 302")
	}
}
