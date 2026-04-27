package pkgs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"html/template"
)

type renderCreatorContextKey struct{}

func WithRenderCreatorForDev(ctx context.Context, htmlTemplate, viteServer, workdir string) (context.Context, error) {
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, renderCreatorContextKey{}, newDevRendererFunc(tmpl, viteServer, workdir)), nil
}

func RenderCreatorFromContext(ctx context.Context) (func(ctx context.Context, entryPoint string) (Renderer, error), error) {
	value := ctx.Value(renderCreatorContextKey{})
	if value == nil {
		return nil, errors.New("no render creator found in context")
	}
	renderCreator, ok := value.(func(ctx context.Context, entryPoint string) (Renderer, error))
	if !ok {
		return nil, errors.New("invalid render creator type")
	}

	return renderCreator, nil
}

type Renderer interface {
	Render(ctx context.Context, props any) ([]byte, error)
}

type devRenderer struct {
	entryPoint   string
	htmlTemplate *template.Template

	viteServer string
	workdir    string
}

type devRendererData struct {
	AppProps   template.JS
	Dev        bool
	ViteServer string
	Workdir    string
	EntryPoint string
}

func newDevRendererFunc(htmlTemplate *template.Template, viteServer, workdir string) func(ctx context.Context, entryPoint string) (Renderer, error) {
	return func(ctx context.Context, entryPoint string) (Renderer, error) {
		generator, err := EntryPointGeneratorFromContext(ctx)
		if err != nil {
			return nil, err
		}

		err = generator.Generate(entryPoint)
		if err != nil {
			return nil, err
		}

		return &devRenderer{
			entryPoint:   entryPoint,
			htmlTemplate: htmlTemplate,
			viteServer:   viteServer,
			workdir:      workdir,
		}, nil
	}
}

func (d *devRenderer) Render(ctx context.Context, props any) ([]byte, error) {
	propsJson, err := json.Marshal(props)
	if err != nil {
		return nil, err
	}

	data := devRendererData{
		AppProps:   template.JS(propsJson),
		Dev:        true,
		ViteServer: d.viteServer,
		Workdir:    d.workdir,
		EntryPoint: d.entryPoint,
	}

	var buf bytes.Buffer
	if err := d.htmlTemplate.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
