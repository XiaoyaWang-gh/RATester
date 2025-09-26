package logs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewNoLevelHook_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	minLevel := zerolog.InfoLevel
	level := zerolog.InfoLevel
	want := &NoLevelHook{minLevel: minLevel, level: level}
	got := NewNoLevelHook(minLevel, level)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("NewNoLevelHook() = %v, want %v", got, want)
	}
}
