package nomad

import (
	"context"
	"fmt"
	"testing"
)

func TestGetNomadServiceDataWithEmptyServices_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	ctx := context.Background()
	items, err := p.getNomadServiceDataWithEmptyServices(ctx)
	if err != nil {
		t.Errorf("getNomadServiceDataWithEmptyServices() error = %v", err)
		return
	}
	if len(items) != 0 {
		t.Errorf("getNomadServiceDataWithEmptyServices() = %v, want %v", items, 0)
	}
}
