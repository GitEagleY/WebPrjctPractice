package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, templateName string) error {
	var err error
	templateCache, err := cacheTemplate()
	if err != nil {
		log.Fatal(err)
		return err
	}
	//make template to render
	templateToRender, ok := templateCache[templateName]

	if !ok {
		log.Fatal(ok)

	}
	//render template
	templateToRender = templateCache[templateName]
	templateToRender.Execute(w, nil)
	return nil
}
func cacheTemplate() (map[string]*template.Template, error) {
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
