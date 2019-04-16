package present

import (
	"html/template"
)

func init() {
	Register("purehtml", parsePureHTML)
}

func parsePureHTML(ctx *Context, fileName string, lineno int, text string) (Elem, error) {
	return PureHTML{template.HTML(text)}, nil
}

type PureHTML struct {
	HTML template.HTML
}

func (s PureHTML) TemplateName() string { return "html" }
