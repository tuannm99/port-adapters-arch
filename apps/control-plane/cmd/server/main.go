package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tuannm99/edge-platform/apps/control-plane/internal/adapters/http/handler"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/adapters/http/router"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/application/edgeconfig"
	websiteapp "github.com/tuannm99/edge-platform/apps/control-plane/internal/application/website"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/infra/nginxconf"
	"github.com/tuannm99/edge-platform/apps/control-plane/internal/infra/store/memory"
)

func main() {
	store := memory.NewWebsiteStore()
	websiteService := websiteapp.NewWebsiteService(store)

	gen, err := nginxconf.NewGenerator()
	if err != nil {
		log.Fatal(err)
	}
	edgeConfigService := edgeconfig.NewEdgeService(gen)

	websiteHandler := handler.NewWebsiteHandler(websiteService, edgeConfigService)
	httpHandler := router.New(websiteHandler)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           httpHandler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("control-plane listening on :8080")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
