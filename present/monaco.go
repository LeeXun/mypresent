package present

import (
	"bytes"
	"html/template"
	"strings"
)

func init() {
	Register("monaco", parseMonaco)
}

type Monaco struct {
}

type manacoTemplateData struct {
}

func (m Monaco) TemplateName() string { return "monaco" }

var monacoTemplate = template.Must(template.New("monaco").Funcs(template.FuncMap{
	"trimSpace":    strings.TrimSpace,
	"leadingSpace": leadingSpaceRE.FindString,
}).Parse(myCodeT))

const myCodeT = `
<section class="try">
    <div class="container">
    <h3>Editor</h3>
        <div class="editor row">
            <div class="span3">                 
                <p>Colorizers are implemented using <a href="monarch.html"
                    target="_blank">Monarch</a>.</p>
            </div>
            <div class="span9">
                <div class="row">
                    <div class="span4">
                        <label class="control-label">Language</label>
                        <select class="language-picker"></select>
                    </div>
                    <div class="span4">
                        <label class="control-label">Theme</label>
                        <select class="theme-picker">
                            <option>Visual Studio</option>
                            <option>Visual Studio Dark</option>
                            <option>High Contrast Dark</option>
                        </select>
                    </div>
                </div>
                <div class="editor-frame">
                    <div class="loading editor" style="display: none;">
                        <div class="progress progress-striped active">
                            <div class="bar"></div>
                        </div>
                    </div>
                    <div id="editor"></div>
                </div>
            </div>
        </div>   
   </div>
</section>
`

func parseMonaco(ctx *Context, sourceFile string, sourceLine int, cmd string) (Elem, error) {
	data := &manacoTemplateData{}

	var buf bytes.Buffer
	if err := monacoTemplate.Execute(&buf, data); err != nil {
		return nil, err
	}
	return Monaco{}, nil
}
