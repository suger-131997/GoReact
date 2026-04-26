package pkgs

import (
	"bytes"
	"context"
	"errors"
	"html/template"
)

type RendererCreator struct {
}

func (r *RendererCreator) Get(ctx context.Context, entryPoint string) Renderer {
	return nil
}

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
	workDir    string
}

type devRendererData struct {
	Dev           bool
	ViteServer    string
	EntryPointDir string
	EntryPoint    string
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
			workDir:      workdir,
		}, nil
	}
}

func (d *devRenderer) Render(ctx context.Context, props any) ([]byte, error) {
	data := devRendererData{
		Dev:           true,
		ViteServer:    d.viteServer,
		EntryPointDir: d.workDir,
		EntryPoint:    d.entryPoint,
	}

	var buf bytes.Buffer
	err := d.htmlTemplate.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
