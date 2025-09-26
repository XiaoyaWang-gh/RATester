package rungroup

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestRun_9(t *testing.T) {
	type testCase struct {
		name    string
		ctx     context.Context
		cfg     Config[int]
		wantErr bool
	}

	tests := []testCase{
		{
			name: "test case 1",
			ctx:  context.Background(),
			cfg: Config[int]{
				NumWorkers: 1,
				Handle: func(ctx context.Context, i int) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			ctx:  context.Background(),
			cfg: Config[int]{
				NumWorkers: 0,
				Handle: func(ctx context.Context, i int) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "test case 3",
			ctx:  context.Background(),
			cfg: Config[int]{
				NumWorkers: 1,
				Handle: func(ctx context.Context, i int) error {
					return errors.New("test error")
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			g := Run(tt.ctx, tt.cfg)
			for i := 0; i < tt.cfg.NumWorkers; i++ {
				g.Enqueue(i)
			}
			err := g.Wait()
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
