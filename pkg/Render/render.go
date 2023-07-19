package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	config "github.com/GitEagleY/WebPrjctPractice/pkg/config"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func RenderTemplate(w http.ResponseWriter, templateName string) error {

	templateCache := app.TemplateCache

	//make template to render
	templateToRender, ok := templateCache[templateName]

	if !ok {
		log.Fatal(ok)

	}
	//render template
	templateToRender.Execute(w, nil)
	return nil
}
func CacheTemplate() (map[string]*template.Template, error) {
	Cache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return Cache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		tmplts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return Cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return Cache, err
		}

		if len(matches) > 0 {
			tmplts, err = tmplts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return Cache, err
			}
		}

		Cache[name] = tmplts
	}

	return Cache, nil
}
