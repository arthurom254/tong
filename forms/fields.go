package forms

type Field interface {
	GetValue() any
}

type BaseField struct {
	Name     string
	Widget   Widget
	Errors   []string
	Required bool
	Label    string
	HelpText string
	Disabled bool
	Attrs    map[string]any `json:"attr"`
}

func (b BaseField) GetLabel() string {
	if b.Label != "" {
		return b.Label
	}
	if b.Name != "" {
		return Humanize(b.Name)
	}
	return ""
}

// text

// type TextField struct {
// 	BaseField
// }

// password

type PasswordField struct {
	BaseField
	Value any `validators:"required"`
}

// email

type EmailField struct {
	BaseField
}

// number
type NumberField struct {
	BaseField
}

// tel
type TelField struct {
	BaseField
}

// url

type URLField struct {
	BaseField
}

// search
type SearchField struct {
	BaseField
}

// date

type DateField struct {
	BaseField
}

// time

type TimeField struct {
	BaseField
}

// datetime-local

type DateTimeLocalField struct {
	BaseField
}

// month

type MonthField struct {
	BaseField
}

// week

type WeekField struct {
	BaseField
}

// color

type ColorField struct {
	BaseField
}

// range

type RangeField struct {
	BaseField
}

// checkbox

type CheckboxField struct {
	BaseField
}

// radio

type RadioField struct {
	BaseField
}

// file

type FileField struct {
	BaseField
}

// hidden

type HiddenField struct {
	BaseField
}

// image
type ImageField struct {
	BaseField
}

// textarea
type TextareaField struct {
	BaseField
}

// select
type SelectField struct {
	BaseField
}

// option

type OptionField struct {
	BaseField
}

// optgroup
type OptionGroupField struct {
	BaseField
}

// datalist

// progress
// legend

func Start(f Field) {

}
