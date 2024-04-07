package templating

import (
	"embed"
	"go-with-test/reading_files"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	template *template.Template
}

func NewPostRenderer() (PostRenderer, error) {
	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return PostRenderer{}, err
	}
	return PostRenderer{template: tmpl}, nil
}

func (r PostRenderer) Render(w io.Writer, post reading_files.Post) error {
	if err := r.template.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}
	return nil
}
