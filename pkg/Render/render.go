package render

import (
	"fmt"
	"html/template"
	"net/http"
)

var templateCache = make(map[string]*template.Template) //template cache

func RenderTemplate(w http.ResponseWriter, templateName string) {
	var templateToRender *template.Template
	_, isInMap := templateCache[templateName]
	if !isInMap {
		err := cacheTemplate(templateName)
		if err != nil {
			return
		}
		fmt.Println("cached")
	} else {
		fmt.Println("using cached")
	}
	templateToRender = templateCache[templateName]
	templateToRender.Execute(w, nil)
}

func cacheTemplate(templateName string) error {
	tmplts := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.tmpl",
	}

	templateParsed, err := template.ParseFiles(tmplts...)
	if err != nil {
		return err
	}
	templateCache[templateName] = templateParsed
	return nil
}
