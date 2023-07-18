package config

import "html/template"

//holds app config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
