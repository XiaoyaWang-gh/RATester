package logger

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestDefaultLoggerInstance_1(t *testing.T) {
	type args struct {
		c    fiber.Ctx
		data *Data
		cfg  Config
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

			if err := defaultLoggerInstance(tt.args.c, tt.args.data, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("defaultLoggerInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
