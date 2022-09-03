package templator

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/czyt/kilot/cmd/kilot/pkg/formater"
	"text/template"
)

type TemplateWorker struct {
	// templates name
	name string
	// templates
	tpl string
	// formatCode
	formatCode bool
}

func New(name string, tpl string, formatCode bool) *TemplateWorker {
	return &TemplateWorker{
		name:       name,
		tpl:        tpl,
		formatCode: formatCode,
	}
}

func (t *TemplateWorker) ChangeCodeFormatOption(formatCode bool) {
	t.formatCode = formatCode
}

func (t *TemplateWorker) ChangeTemplateOption(tpl string) {
	t.tpl = tpl
}

func (t *TemplateWorker) Execute(ctxValue interface{}) (*bytes.Buffer, error) {
	tpl, err := template.New(t.name).Parse(t.tpl)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("templates parse error:%v", t.tpl))
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, ctxValue); err != nil {
		return nil, err
	}
	if !t.formatCode {
		return buf, nil
	}
	if formatedCode, err := formater.FormatCode(buf.Bytes()); err != nil {
		return nil, err
	} else {
		buf.Reset()
		buf.Write(formatedCode)
	}
	return buf, nil
}
