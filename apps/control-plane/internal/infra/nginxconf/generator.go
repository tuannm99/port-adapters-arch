package nginxconf

import (
	"bytes"
	"context"
	"text/template"

	"github.com/tuannm99/edge-platform/apps/control-plane/internal/application/edgeconfig"
)

var _ edgeconfig.Renderer = (*Generator)(nil)

type Generator struct {
	tpl *template.Template
}

func NewGenerator() (*Generator, error) {
	tpl, err := template.New("server").Parse(serverTemplate)
	if err != nil {
		return nil, err
	}

	return &Generator{tpl: tpl}, nil
}

func (g *Generator) Render(_ context.Context, in edgeconfig.RenderInput) (string, error) {
	var buf bytes.Buffer

	if err := g.tpl.Execute(&buf, in); err != nil {
		return "", err
	}

	return buf.String(), nil
}
