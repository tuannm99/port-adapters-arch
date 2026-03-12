package edgeconfig_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/application/edgeconfig"
	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/application/edgeconfig/mocks"
)

func TestEdgeService_BuildWebsiteConfig(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(r *mocks.MockRenderer)
		wantErr bool
	}{
		{
			"success",
			func(r *mocks.MockRenderer) {
				r.EXPECT().
					Render(context.Background(), edgeconfig.RenderInput{
						Domain:   "demo.local",
						Upstream: "http://demo:8080",
					}).
					Return("config", nil)
			},
			false,
		},
		{
			"renderer error",
			func(r *mocks.MockRenderer) {
				r.EXPECT().
					Render(context.Background(), edgeconfig.RenderInput{
						Domain:   "demo.local",
						Upstream: "http://demo:8080",
					}).
					Return("", errors.New("render error"))
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer := mocks.NewMockRenderer(t)
			tt.mock(renderer)

			svc := edgeconfig.NewEdgeService(renderer)

			_, err := svc.BuildWebsiteConfig(
				context.Background(),
				"demo.local",
				"http://demo:8080",
			)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
