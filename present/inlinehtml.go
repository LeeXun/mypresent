package present

import (
	"html/template"
)

func init() {
	Register("inlinehtml", parseInlineHTML)
}

func parseInlineHTML(ctx *Context, fileName string, lineno int, text string) (Elem, error) {
	return InlineHTML{template.HTML(text)}, nil
}

type InlineHTML struct {
	HTML template.HTML
}

func (i InlineHTML) TemplateName() string { return "inlinehtml" }
