package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestDebug_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	l := OxyWrapper{
		logger: zerolog.New(os.Stdout),
	}
	s := "test"
	i := []interface{}{"test"}

	// when
	l.Debug(s, i...)

	// then
	// TODO
}
