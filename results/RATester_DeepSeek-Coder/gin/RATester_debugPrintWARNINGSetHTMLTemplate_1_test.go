package gin

import (
	"fmt"
	"testing"
)

func TestDebugPrintWARNINGSetHTMLTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	debugPrintWARNINGSetHTMLTemplate()

	// You can add more assertions here if you want to check the output
}
