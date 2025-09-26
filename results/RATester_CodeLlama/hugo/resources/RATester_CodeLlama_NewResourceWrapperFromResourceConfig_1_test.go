package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestNewResourceWrapperFromResourceConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var rc *pagemeta.ResourceConfig
	var r *Spec
	var err error
	var resource resource.Resource

	// Act
	resource, err = r.NewResourceWrapperFromResourceConfig(rc)

	// Assert
	if err != nil {
		t.Errorf("NewResourceWrapperFromResourceConfig() error = %v", err)
		return
	}

	if resource == nil {
		t.Errorf("NewResourceWrapperFromResourceConfig() resource = nil")
		return
	}
}
