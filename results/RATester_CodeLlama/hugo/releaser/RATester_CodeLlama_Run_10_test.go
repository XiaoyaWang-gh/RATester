package releaser

import (
	"fmt"
	"testing"
)

func TestRun_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &ReleaseHandler{
		branchVersion: "0.90.0",
		step:          1,
	}

	if err := r.Run(); err != nil {
		t.Fatal(err)
	}
}
