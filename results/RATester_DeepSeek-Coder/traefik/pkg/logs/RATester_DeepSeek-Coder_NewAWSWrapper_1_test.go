package logs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/rs/zerolog"
)

func TestNewAWSWrapper_1(t *testing.T) {
	type args struct {
		logger zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want aws.LoggerFunc
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

			if got := NewAWSWrapper(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAWSWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
