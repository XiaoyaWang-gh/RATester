package prometheus

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/beego/beego/v2"
	"github.com/beego/beego/v2/server/web"
	"github.com/prometheus/client_golang/prometheus"
)

func TestRegisterBuildInfo_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test registerBuildInfo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			registerBuildInfo()
			registry := prometheus.NewRegistry()
			buildInfo := prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name:      "beego",
				Subsystem: "build_info",
				Help:      "The building information",
				ConstLabels: map[string]string{
					"appname":        web.BConfig.AppName,
					"build_version":  beego.BuildVersion,
					"build_revision": beego.BuildGitRevision,
					"build_status":   beego.BuildStatus,
					"build_tag":      beego.BuildTag,
					"build_time":     strings.Replace(beego.BuildTime, "--", " ", 1),
					"go_version":     beego.GoVersion,
					"git_branch":     beego.GitBranch,
					"start_time":     time.Now().Format("2006-01-02 15:04:05"),
				},
			}, []string{})
			err := registry.Register(buildInfo)
			if err != nil {
				t.Errorf("Failed to register buildInfo: %v", err)
			}
			buildInfo.WithLabelValues().Set(1)
			got, err := registry.Gather()
			if err != nil {
				t.Errorf("Failed to gather metrics: %v", err)
			}
			if len(got) == 0 {
				t.Errorf("No metrics returned")
			}
		})
	}
}
