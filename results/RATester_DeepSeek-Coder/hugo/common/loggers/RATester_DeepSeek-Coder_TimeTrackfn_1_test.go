package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestTimeTrackfn_1(t *testing.T) {
	type args struct {
		fn func() (logg.LevelLogger, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := TimeTrackfn(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("TimeTrackfn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
