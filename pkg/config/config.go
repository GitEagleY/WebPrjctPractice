package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// holds app config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool //for developing mode
	Production    bool
	Session       *scs.SessionManager
}
