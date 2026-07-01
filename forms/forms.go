package forms

import (
	"fmt"

	"github.com/flosch/pongo2/v6"
)

type Form struct {
	Fields []Field
	Errors []string
	Data   any
}

func (f Form) getContext() pongo2.Context {
	ctx := pongo2.Context{
		"errors":        f.Errors,
		"fields":        f.Fields,
		"hidden_fields": []Field{},
	}
	return ctx
}

func (f Form) render() *pongo2.Value {
	return f.AsDiv()
}

func (f Form) AsDiv() *pongo2.Value {
	ctx := f.getContext()
	tpl, err := renderTPL("div.html", ctx)
	if err != nil {
		panic(err)
	}
	return pongo2.AsSafeValue(tpl)
}
func (f Form) AsP() *pongo2.Value {
	ctx := f.getContext()
	tpl, err := renderTPL("p.html", ctx)
	if err != nil {
		panic(err)
	}
	return pongo2.AsSafeValue(tpl)
}
func (f Form) AsTable() *pongo2.Value {
	ctx := f.getContext()
	tpl, err := renderTPL("table.html", ctx)
	if err != nil {
		panic(err)
	}
	return pongo2.AsSafeValue(tpl)
}

func (f Form) IsBound() bool {
	return f.Data != nil
}

func (f Form) String() *pongo2.Value {
	return pongo2.AsSafeValue(f.AsDiv())
}

func (f *Form) Clean() error {
	if !f.IsBound() {
		return nil
	}
	for i := 0; i < len(f.Fields); i++ {
		fmt.Println("Validated Form")
	}
	return nil
}
