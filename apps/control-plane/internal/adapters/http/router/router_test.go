package router_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/adapters/http/handler"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/adapters/http/router"
)

func TestNewRouter(t *testing.T) {
	websiteHandler := handler.NewWebsiteHandler(nil, nil)
	h := router.New(websiteHandler)
	assert.Implements(t, new(http.Handler), h)
}
