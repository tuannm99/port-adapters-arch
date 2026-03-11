package edgeconfig

import "context"

type RenderInput struct {
	Domain   string
	Upstream string
}

type Renderer interface {
	Render(ctx context.Context, in RenderInput) (string, error)
}
