package identity

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/compare"
)

func TestGetSearchID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var searchID *searchID
	var id Identity
	var isDp bool
	var isPeq bool
	var hasEqer bool
	var maxDepth int
	var seen map[Manager]bool
	var dp IsProbablyDependentProvider
	var peq compare.ProbablyEqer
	var eqer compare.Eqer
	// Act
	searchID = getSearchID()
	id = searchID.id
	isDp = searchID.isDp
	isPeq = searchID.isPeq
	hasEqer = searchID.hasEqer
	maxDepth = searchID.maxDepth
	seen = searchID.seen
	dp = searchID.dp
	peq = searchID.peq
	eqer = searchID.eqer
	// Assert
	if id != nil {
		t.Errorf("Expected id to be nil")
	}
	if isDp != false {
		t.Errorf("Expected isDp to be false")
	}
	if isPeq != false {
		t.Errorf("Expected isPeq to be false")
	}
	if hasEqer != false {
		t.Errorf("Expected hasEqer to be false")
	}
	if maxDepth != 0 {
		t.Errorf("Expected maxDepth to be 0")
	}
	if seen != nil {
		t.Errorf("Expected seen to be nil")
	}
	if dp != nil {
		t.Errorf("Expected dp to be nil")
	}
	if peq != nil {
		t.Errorf("Expected peq to be nil")
	}
	if eqer != nil {
		t.Errorf("Expected eqer to be nil")
	}
}
