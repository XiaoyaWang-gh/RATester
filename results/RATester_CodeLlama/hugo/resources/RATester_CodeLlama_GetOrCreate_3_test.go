package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestGetOrCreate_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ResourceCache{}
	key := "key"
	f := func() (resource.Resource, error) {
		return nil, nil
	}
	_, err := c.GetOrCreate(key, f)
	if err != nil {
		t.Errorf("GetOrCreate() error = %v, want nil", err)
		return
	}
}
