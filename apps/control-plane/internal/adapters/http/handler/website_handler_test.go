package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/adapters/http/handler"
	edgeconfigmocks "github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/application/edgeconfig/mocks"
	websiteappmocks "github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/application/website/mocks"
	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/domain/website"
)

func withChiParam(req *http.Request, key, val string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, val)
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	return req.WithContext(ctx)
}

func newHandler(
	t *testing.T,
) (*handler.WebsiteHandler, *websiteappmocks.MockWebsiteUseCase, *edgeconfigmocks.MockEdgeConfigUseCase) {
	w := websiteappmocks.NewMockWebsiteUseCase(t)
	e := edgeconfigmocks.NewMockEdgeConfigUseCase(t)
	return handler.NewWebsiteHandler(w, e), w, e
}

func TestWebsiteHandler_Create(t *testing.T) {
	tests := []struct {
		name   string
		body   string
		mock   func(w *websiteappmocks.MockWebsiteUseCase)
		status int
	}{
		// invalid json
		{"bad json", "{bad", nil, http.StatusBadRequest},

		// usecase error
		{
			"usecase error",
			`{"domain":"d","upstream":"u"}`,
			func(w *websiteappmocks.MockWebsiteUseCase) {
				w.EXPECT().
					Create(mock.Anything, mock.Anything).
					Return(website.Website{}, errors.New("err"))
			},
			http.StatusBadRequest,
		},

		// success
		{
			"success",
			`{"domain":"d","upstream":"u"}`,
			func(w *websiteappmocks.MockWebsiteUseCase) {
				w.EXPECT().
					Create(mock.Anything, mock.Anything).
					Return(website.Website{ID: "1"}, nil)
			},
			http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h, w, _ := newHandler(t)
			if tt.mock != nil {
				tt.mock(w)
			}

			req := httptest.NewRequest(http.MethodPost, "/websites", strings.NewReader(tt.body))
			rr := httptest.NewRecorder()

			h.Create(rr, req)

			require.Equal(t, tt.status, rr.Code)
		})
	}
}

func TestWebsiteHandler_List(t *testing.T) {
	tests := []struct {
		name   string
		mock   func(w *websiteappmocks.MockWebsiteUseCase)
		status int
	}{
		// usecase error
		{
			"error",
			func(w *websiteappmocks.MockWebsiteUseCase) {
				w.EXPECT().
					List(mock.Anything).
					Return(nil, errors.New("err"))
			},
			http.StatusInternalServerError,
		},

		// success
		{
			"success",
			func(w *websiteappmocks.MockWebsiteUseCase) {
				w.EXPECT().
					List(mock.Anything).
					Return([]website.Website{{ID: "1"}}, nil)
			},
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h, w, _ := newHandler(t)
			tt.mock(w)

			req := httptest.NewRequest(http.MethodGet, "/websites", nil)
			rr := httptest.NewRecorder()

			h.List(rr, req)

			require.Equal(t, tt.status, rr.Code)
		})
	}
}

func TestWebsiteHandler_GetNginxConfig(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		mock   func(w *websiteappmocks.MockWebsiteUseCase, e *edgeconfigmocks.MockEdgeConfigUseCase)
		status int
	}{
		// missing id
		{"missing id", "", nil, http.StatusBadRequest},

		// website not found
		{
			"website error",
			"1",
			func(w *websiteappmocks.MockWebsiteUseCase, _ *edgeconfigmocks.MockEdgeConfigUseCase) {
				w.EXPECT().
					GetByID(mock.Anything, "1").
					Return(website.Website{}, errors.New("err"))
			},
			http.StatusNotFound,
		},

		// config build error
		{
			"config error",
			"1",
			func(w *websiteappmocks.MockWebsiteUseCase, e *edgeconfigmocks.MockEdgeConfigUseCase) {
				w.EXPECT().
					GetByID(mock.Anything, "1").
					Return(website.Website{
						Domain:   "d",
						Upstream: "u",
					}, nil)

				e.EXPECT().
					BuildWebsiteConfig(mock.Anything, "d", "u").
					Return("", errors.New("err"))
			},
			http.StatusInternalServerError,
		},

		// success
		{
			"success",
			"1",
			func(w *websiteappmocks.MockWebsiteUseCase, e *edgeconfigmocks.MockEdgeConfigUseCase) {
				w.EXPECT().
					GetByID(mock.Anything, "1").
					Return(website.Website{
						Domain:   "d",
						Upstream: "u",
					}, nil)

				e.EXPECT().
					BuildWebsiteConfig(mock.Anything, "d", "u").
					Return("cfg", nil)
			},
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h, w, e := newHandler(t)

			if tt.mock != nil {
				tt.mock(w, e)
			}

			req := httptest.NewRequest(http.MethodGet, "/websites/"+tt.id+"/nginx-config", nil)

			if tt.id != "" {
				req = withChiParam(req, "id", tt.id)
			}

			rr := httptest.NewRecorder()

			h.GetNginxConfig(rr, req)

			require.Equal(t, tt.status, rr.Code)
		})
	}
}
