package mem

import (
	"fmt"
	"testing"
	"time"
)

func TestClearUselessInfo_1(t *testing.T) {
	type args struct {
		continuousOfflineDuration time.Duration
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "Test case 1",
			args: args{
				continuousOfflineDuration: time.Hour,
			},
			want:  1,
			want1: 2,
		},
		{
			name: "Test case 2",
			args: args{
				continuousOfflineDuration: time.Minute,
			},
			want:  0,
			want1: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &serverMetrics{
				info: &ServerStatistics{
					ProxyStatistics: map[string]*ProxyStatistics{
						"proxy1": {
							LastCloseTime: time.Now().Add(-time.Hour),
						},
						"proxy2": {
							LastCloseTime: time.Now().Add(-time.Minute),
						},
					},
				},
			}
			got, got1 := m.clearUselessInfo(tt.args.continuousOfflineDuration)
			if got != tt.want {
				t.Errorf("clearUselessInfo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("clearUselessInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
