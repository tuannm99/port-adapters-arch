package websiteapp_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	websiteapp "github.com/tuannm99/edge-platform/apps/control-plane/internal/application/website"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/application/website/mocks"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/domain/website"
)

////////////////////////////////////////////////////////
// CREATE
////////////////////////////////////////////////////////

func TestWebsiteService_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   websiteapp.CreateInput
		mock    func(r *mocks.MockRepository)
		wantErr bool
	}{
		{
			"missing domain",
			websiteapp.CreateInput{Upstream: "u"},
			nil,
			true,
		},
		{
			"missing upstream",
			websiteapp.CreateInput{Domain: "d"},
			nil,
			true,
		},
		{
			"repo error",
			websiteapp.CreateInput{Domain: "d", Upstream: "u"},
			func(r *mocks.MockRepository) {
				r.EXPECT().
					Create(mock.Anything, mock.Anything).
					Return(errors.New("db error"))
			},
			true,
		},
		{
			"success",
			websiteapp.CreateInput{Domain: "d", Upstream: "u"},
			func(r *mocks.MockRepository) {
				r.EXPECT().
					Create(mock.Anything, mock.Anything).
					Return(nil)
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockRepository(t)

			if tt.mock != nil {
				tt.mock(repo)
			}

			svc := websiteapp.NewWebsiteService(repo)

			_, err := svc.Create(context.Background(), tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

////////////////////////////////////////////////////////
// LIST
////////////////////////////////////////////////////////

func TestWebsiteService_List(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(r *mocks.MockRepository)
		wantErr bool
	}{
		{
			"repo error",
			func(r *mocks.MockRepository) {
				r.EXPECT().
					List(mock.Anything).
					Return(nil, errors.New("db error"))
			},
			true,
		},
		{
			"success",
			func(r *mocks.MockRepository) {
				r.EXPECT().
					List(mock.Anything).
					Return([]website.Website{{ID: "1"}}, nil)
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockRepository(t)
			tt.mock(repo)

			svc := websiteapp.NewWebsiteService(repo)

			_, err := svc.List(context.Background())

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

////////////////////////////////////////////////////////
// GET BY ID
////////////////////////////////////////////////////////

func TestWebsiteService_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(r *mocks.MockRepository)
		wantErr bool
	}{
		{
			"repo error",
			func(r *mocks.MockRepository) {
				r.EXPECT().
					GetByID(mock.Anything, "1").
					Return(website.Website{}, errors.New("not found"))
			},
			true,
		},
		{
			"success",
			func(r *mocks.MockRepository) {
				r.EXPECT().
					GetByID(mock.Anything, "1").
					Return(website.Website{ID: "1"}, nil)
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockRepository(t)
			tt.mock(repo)

			svc := websiteapp.NewWebsiteService(repo)

			_, err := svc.GetByID(context.Background(), "1")

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
