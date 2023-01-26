package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Logger        *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
