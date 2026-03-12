package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/adapters/http/dto"
	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/adapters/http/httpx"
	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/application/edgeconfig"
	websiteapp "github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/application/website"
)

type WebsiteHandler struct {
	websiteService    websiteapp.WebsiteUseCase
	edgeConfigService edgeconfig.EdgeConfigUseCase
}

func NewWebsiteHandler(
	websiteService websiteapp.WebsiteUseCase,
	edgeConfigService edgeconfig.EdgeConfigUseCase,
) *WebsiteHandler {
	return &WebsiteHandler{
		websiteService:    websiteService,
		edgeConfigService: edgeConfigService,
	}
}

func (h *WebsiteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateWebsiteRequest

	if err := httpx.DecodeJSON(r, &req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid json body")
		return
	}

	result, err := h.websiteService.Create(
		r.Context(),
		websiteapp.CreateInput{
			Domain:   req.Domain,
			Upstream: req.Upstream,
		},
	)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	httpx.JSON(w, http.StatusCreated, dto.ToWebsiteResponse(result))
}

func (h *WebsiteHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.websiteService.List(r.Context())
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, dto.ToWebsiteResponses(items))
}

func (h *WebsiteHandler) GetNginxConfig(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		httpx.Error(w, http.StatusBadRequest, "id is required")
		return
	}

	item, err := h.websiteService.GetByID(r.Context(), id)
	if err != nil {
		httpx.Error(w, http.StatusNotFound, err.Error())
		return
	}

	cfg, err := h.edgeConfigService.BuildWebsiteConfig(
		r.Context(),
		item.Domain,
		item.Upstream,
	)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, dto.WebsiteConfigResponse{
		WebsiteID: item.ID,
		Domain:    item.Domain,
		Upstream:  item.Upstream,
		Config:    cfg,
	})
}
