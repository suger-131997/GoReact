package pkgs

import (
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type EntryPointGenerator struct {
	dir      string
	rootTmpl *template.Template
}

func NewEntryPointGenerator(entryPointDir string) *EntryPointGenerator {
	return &EntryPointGenerator{
		dir:      entryPointDir,
		rootTmpl: template.Must(template.New("root").Parse(rootTemplate)),
	}
}

func (r *EntryPointGenerator) Generate(entryPoint string) error {
	p := filepath.Join(r.dir, entryPoint)

	err := os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	err = r.rootTmpl.Execute(f, map[string]interface{}{
		"EntryPoint": entryPoint,
	})
	if err != nil {
		return err
	}

	return nil
}

const rootTemplate = `
import { StrictMode } from "react";
import { createRoot } from 'react-dom/client'
import App from '~/{{ .EntryPoint }}'


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App {...(window.APP_PROPS || {})}/>
  </StrictMode>
)
`
