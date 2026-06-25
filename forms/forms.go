package forms

import "fmt"

type Form struct {
	Fields []Widget
	Errors map[string]any
	Data   *any
}

func (f Form) AsP() string {
	return "<p>Paragrap</p>"
}

func (f Form) AsDiv() string {
	return "<div></div>"
}

func (f Form) isBound() bool {
	return f.Data == nil
}

func (f *Form) Clean() error {
	if !f.isBound() {
		return nil
	}
	for i := 0; i < len(f.Fields); i++ {
		fmt.Println("Validated Form")
	}
	return nil
}
