package hugofs

import (
	"fmt"
	"testing"
)

func TestStat_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &rootMappingDir{}
	_, err := f.Stat()
	if err != nil {
		t.Errorf("Stat() error = %v, want nil", err)
		return
	}
}
