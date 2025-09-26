package page

import (
	"fmt"
	"testing"
)

func TestRelRefFrom_2(t *testing.T) {
	type args struct {
		argsm  map[string]any
		source any
	}
	tests := []struct {
		name    string
		p       *nopPage
		args    args
		want    string
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

			got, err := tt.p.RelRefFrom(tt.args.argsm, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("nopPage.RelRefFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("nopPage.RelRefFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
