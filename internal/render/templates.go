package render

import "embed"

//go:embed templates/*.tmpl
// TemplateFS holds our embedded templates
var TemplateFS embed.FS