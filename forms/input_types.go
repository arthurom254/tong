package forms

func newCharField(inputType, name, label, help string, required bool, attrs map[string]any, value string) CharField {
	if attrs == nil {
		attrs = map[string]any{}
	}
	attrs["type"] = inputType
	return CharField{
		Name: name, Label: label, HelpText: help, Required: required,
		Errors: []string{}, Attrs: attrs, InputType: inputType, Value: value,
	}
}

type PasswordField struct{ CharField }

func NewPasswordField(name, label, help string, required bool, attrs map[string]any, value string) PasswordField {
	return PasswordField{newCharField("password", name, label, help, required, attrs, value)}
}

type EmailField struct{ CharField }

func NewEmailField(name, label, help string, required bool, attrs map[string]any, value string) EmailField {
	return EmailField{newCharField("email", name, label, help, required, attrs, value)}
}

type NumberField struct{ CharField }

func NewNumberField(name, label, help string, required bool, attrs map[string]any, value string) NumberField {
	return NumberField{newCharField("number", name, label, help, required, attrs, value)}
}

type TelField struct{ CharField }

func NewTelField(name, label, help string, required bool, attrs map[string]any, value string) TelField {
	return TelField{newCharField("tel", name, label, help, required, attrs, value)}
}

type URLField struct{ CharField }

func NewURLField(name, label, help string, required bool, attrs map[string]any, value string) URLField {
	return URLField{newCharField("url", name, label, help, required, attrs, value)}
}

type SearchField struct{ CharField }

func NewSearchField(name, label, help string, required bool, attrs map[string]any, value string) SearchField {
	return SearchField{newCharField("search", name, label, help, required, attrs, value)}
}

type DateField struct{ CharField }

func NewDateField(name, label, help string, required bool, attrs map[string]any, value string) DateField {
	return DateField{newCharField("date", name, label, help, required, attrs, value)}
}

type TimeField struct{ CharField }

func NewTimeField(name, label, help string, required bool, attrs map[string]any, value string) TimeField {
	return TimeField{newCharField("time", name, label, help, required, attrs, value)}
}

type DateTimeLocalField struct{ CharField }

func NewDateTimeLocalField(name, label, help string, required bool, attrs map[string]any, value string) DateTimeLocalField {
	return DateTimeLocalField{newCharField("datetime-local", name, label, help, required, attrs, value)}
}

type MonthField struct{ CharField }

func NewMonthField(name, label, help string, required bool, attrs map[string]any, value string) MonthField {
	return MonthField{newCharField("month", name, label, help, required, attrs, value)}
}

type WeekField struct{ CharField }

func NewWeekField(name, label, help string, required bool, attrs map[string]any, value string) WeekField {
	return WeekField{newCharField("week", name, label, help, required, attrs, value)}
}

type ColorField struct{ CharField }

func NewColorField(name, label, help string, required bool, attrs map[string]any, value string) ColorField {
	return ColorField{newCharField("color", name, label, help, required, attrs, value)}
}

type RangeField struct{ CharField }

func NewRangeField(name, label, help string, required bool, attrs map[string]any, value string) RangeField {
	return RangeField{newCharField("range", name, label, help, required, attrs, value)}
}

type HiddenField struct{ CharField }

func NewHiddenField(name string, value string) HiddenField {
	f := newCharField("hidden", name, "", "", false, map[string]any{}, value)
	f.UseFieldset = false
	return HiddenField{f}
}

type ImageField struct{ CharField }

func NewImageField(name, label, help string, required bool, attrs map[string]any) ImageField {
	f := newCharField("file", name, label, help, required, attrs, "")
	if f.Attrs == nil {
		f.Attrs = map[string]any{}
	}
	f.Attrs["accept"] = "image/*"
	return ImageField{f}
}

type FileField struct{ CharField }

func NewFileField(name, label, help string, required bool, attrs map[string]any) FileField {
	return FileField{newCharField("file", name, label, help, required, attrs, "")}
}
