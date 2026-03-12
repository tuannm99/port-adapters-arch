package dto

import "github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/domain/website"

type CreateWebsiteRequest struct {
	Domain   string `json:"domain"`
	Upstream string `json:"upstream"`
}

type WebsiteResponse struct {
	ID       string `json:"id"`
	Domain   string `json:"domain"`
	Upstream string `json:"upstream"`
}

func ToWebsiteResponse(w website.Website) WebsiteResponse {
	return WebsiteResponse{
		ID:       w.ID,
		Domain:   w.Domain,
		Upstream: w.Upstream,
	}
}

func ToWebsiteResponses(items []website.Website) []WebsiteResponse {
	out := make([]WebsiteResponse, len(items))

	for i := range items {
		out[i] = ToWebsiteResponse(items[i])
	}

	return out
}

// website configs
type WebsiteConfigResponse struct {
	WebsiteID string `json:"website_id"`
	Domain    string `json:"domain"`
	Upstream  string `json:"upstream"`
	Config    string `json:"config"`
}
