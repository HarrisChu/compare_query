package pkg

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed tpl
var templates embed.FS

type QueryFactory struct {
	template *template.Template
}

func NewQueryFactory(templateName string) (*QueryFactory, error) {
	files, err := templates.ReadDir("tpl")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.Name() == templateName {
			content, err := templates.ReadFile("tpl/" + templateName)
			if err != nil {
				return nil, err
			}
			//first line
			s := strings.Split(string(content), "\n")[0]
			tmpl, err := template.New(templateName).Parse(s)
			if err != nil {
				return nil, err
			}
			return &QueryFactory{
				template: tmpl,
			}, nil
		}
	}
	return nil, fmt.Errorf("cannot find the template %s", templateName)
}

func (qf *QueryFactory) NewQuery(vid string) (string, error) {
	var bs bytes.Buffer
	params := struct{ Vid string }{Vid: vid}
	if err := qf.template.Execute(&bs, params); err != nil {
		return "", err
	}
	return bs.String(), nil
}
