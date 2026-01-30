package render

import (
	"bytes"
	"context"
	"encoding/json"
	"goreact/internal/contextutil"
	"html/template"
)

type PageData struct {
	Title    string
	JS       template.JS
	AppProps template.JS
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="icon" type="image/svg+xml" href="/_assets/vite.svg" />
    <link href="/_assets/output.css" rel="stylesheet">
    <title>{{ .Title }}</title>
</head>
<body>
    <div id="root"></div>
	<script>window.APP_PROPS={{ .AppProps }};</script>
	<script type="module">{{ .JS }}</script>
</body>
</html>
`

type CoreRenderer struct {
	tmpl *template.Template

	jsf func(entryPoint string) func() (string, error)
}

func NewCoreRenderer(jsf func(entryPoint string) func() (string, error)) *CoreRenderer {
	return &CoreRenderer{tmpl: template.Must(template.New("html").Parse(htmlTemplate)), jsf: jsf}
}

type Renderer[T any] struct {
	renderer *CoreRenderer

	js func() (string, error)
}

func NewRenderer[T any](renderer *CoreRenderer, entryPoint string) *Renderer[T] {
	return &Renderer[T]{renderer: renderer, js: renderer.jsf(entryPoint)}
}

func (r *Renderer[T]) Render(ctx context.Context, props T) ([]byte, error) {
	js, err := r.js()
	if err != nil {
		return nil, err
	}

	propsJson, err := json.Marshal(props)
	if err != nil {
		return nil, err
	}

	title, _ := contextutil.TitleFromContext(ctx)

	data := PageData{
		Title:    title,
		JS:       template.JS(js),
		AppProps: template.JS(propsJson),
	}

	var buf bytes.Buffer
	err = r.renderer.tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
