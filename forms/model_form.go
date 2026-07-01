package forms

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var camelRe = regexp.MustCompile(`([a-z0-9])([A-Z])`)

type fieldMeta struct {
	Label, Help, Placeholder, Autocomplete, Widget string
	MaxLength, MinLength                           int
	Required, Exclude, NoPrefill                   bool
}

func parseFormTag(tag string) fieldMeta {
	m := fieldMeta{}
	for _, p := range strings.Split(tag, ";") {
		p = strings.TrimSpace(p)
		switch {
		case p == "":
			continue
		case p == "-":
			m.Exclude = true
		case p == "required":
			m.Required = true
		case p == "exclude_prefill":
			m.NoPrefill = true
		default:
			kv := strings.SplitN(p, "=", 2)
			if len(kv) != 2 {
				continue
			}
			switch kv[0] {
			case "label":
				m.Label = kv[1]
			case "help":
				m.Help = kv[1]
			case "placeholder":
				m.Placeholder = kv[1]
			case "autocomplete":
				m.Autocomplete = kv[1]
			case "widget":
				m.Widget = kv[1]
			case "maxlength":
				m.MaxLength, _ = strconv.Atoi(kv[1])
			case "minlength":
				m.MinLength, _ = strconv.Atoi(kv[1])
			}
		}
	}
	return m
}

func defaultWidget(fieldName string, t reflect.Type) string {
	n := strings.ToLower(fieldName)
	switch {
	case strings.Contains(n, "password"):
		return "password"
	case strings.Contains(n, "email"):
		return "email"
	case strings.Contains(n, "phone"):
		return "tel"
	}
	if t.Kind() >= reflect.Int && t.Kind() <= reflect.Uint64 {
		return "number"
	}
	if t.Kind() == reflect.Bool {
		return "checkbox"
	}
	return "text"
}

var hiddenFields = map[string]bool{"ID": true, "CreatedAt": true, "UpdatedAt": true, "DeletedAt": true}

func NewModelFrom(model interface{}, inputClass string) Form {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var fields []Field
	var walk func(reflect.Type, reflect.Value)
	walk = func(t reflect.Type, v reflect.Value) {
		for i := 0; i < t.NumField(); i++ {
			sf := t.Field(i)
			if sf.Anonymous && sf.Type.Kind() == reflect.Struct {
				walk(sf.Type, v.Field(i))
				continue
			}
			if sf.PkgPath != "" {
				continue
			}

			tag, hasTag := sf.Tag.Lookup("form")
			meta := parseFormTag(tag)
			if meta.Exclude || (hiddenFields[sf.Name] && !hasTag) {
				continue
			}

			name := Slugify(sf)
			label := meta.Label
			if label == "" {
				label = Humanize(sf.Name)
			}
			widget := meta.Widget
			if widget == "" {
				widget = defaultWidget(sf.Name, sf.Type)
			}

			attrs := map[string]any{
				"id": name, "name": name, "class": inputClass,
			}
			if meta.Placeholder != "" {
				attrs["placeholder"] = meta.Placeholder
			}
			if meta.Autocomplete != "" {
				attrs["autocomplete"] = meta.Autocomplete
			}
			if meta.MaxLength > 0 {
				attrs["maxlength"] = meta.MaxLength
			}
			if meta.MinLength > 0 {
				attrs["minlength"] = meta.MinLength
			}

			var value string
			if !meta.NoPrefill {
				if fv := v.Field(i); fv.IsValid() && !fv.IsZero() {
					value = fmt.Sprintf("%v", fv.Interface())
				}
			}

			var field Field
			switch widget {
			case "password":
				field = NewPasswordField(name, label, meta.Help, meta.Required, attrs, value)
			case "email":
				field = NewEmailField(name, label, meta.Help, meta.Required, attrs, value)
			case "tel":
				field = NewTelField(name, label, meta.Help, meta.Required, attrs, value)
			case "url":
				field = NewURLField(name, label, meta.Help, meta.Required, attrs, value)
			case "number":
				field = NewNumberField(name, label, meta.Help, meta.Required, attrs, value)
			case "date":
				field = NewDateField(name, label, meta.Help, meta.Required, attrs, value)
			// case "checkbox":
			// 	checked := v.Field(i).Kind() == reflect.Bool && v.Field(i).Bool()
			// 	field = NewCheckboxField(name, label, meta.Help, checked, attrs)
			// case "textarea":
			// 	field = TextareaField{
			// 		Name: name, Label: label, HelpText: meta.Help, Required: meta.Required,
			// 		Errors: []string{}, Attrs: attrs, Value: value,
			// 	}
			default:
				attrs["type"] = "text"
				field = CharField{
					Name: name, Label: label, Required: meta.Required,
					HelpText: meta.Help, Errors: []string{}, Attrs: attrs, Value: value,
				}
			}
			fields = append(fields, field)

		}
	}
	walk(v.Type(), v)
	return Form{Fields: fields}
}

func validateRaw(label string, required bool, attrs map[string]any, raw string) []string {
	var errs []string
	if required && raw == "" {
		errs = append(errs, label+" is required.")
	}
	if ml, has := attrs["minlength"].(int); has && raw != "" && len(raw) < ml {
		errs = append(errs, fmt.Sprintf("%s must be at least %d characters.", label, ml))
	}
	if ml, has := attrs["maxlength"].(int); has && len(raw) > ml {
		errs = append(errs, fmt.Sprintf("%s must be at most %d characters.", label, ml))
	}
	if t, _ := attrs["type"].(string); t == "email" && raw != "" && !strings.Contains(raw, "@") {
		errs = append(errs, label+" must be a valid email address.")
	}
	return errs
}

func setStructField(fv reflect.Value, raw string) {
	if !fv.CanSet() {
		return
	}
	switch fv.Kind() {
	case reflect.String:
		fv.SetString(raw)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if n, err := strconv.ParseInt(raw, 10, 64); err == nil {
			fv.SetInt(n)
		}
	case reflect.Bool:
		fv.SetBool(raw == "on" || raw == "true")
	}
}

func (f *Form) Bind(c *fiber.Ctx, model interface{}) bool {
	v := reflect.ValueOf(model).Elem()
	dest := map[string]reflect.Value{}
	var walk func(reflect.Type, reflect.Value)
	walk = func(t reflect.Type, v reflect.Value) {
		for i := 0; i < t.NumField(); i++ {
			sf := t.Field(i)
			if sf.Anonymous && sf.Type.Kind() == reflect.Struct {
				walk(sf.Type, v.Field(i))
				continue
			}
			if sf.PkgPath == "" {
				dest[Slugify(sf)] = v.Field(i)
			}
		}
	}
	walk(v.Type(), v)

	ok := true
	for i, field := range f.Fields {
		switch cf := field.(type) {

		case CharField:
			raw := strings.TrimSpace(c.FormValue(cf.Name))
			errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
			if len(errs) > 0 {
				cf.Errors = errs
				f.Fields[i] = cf
				ok = false
				continue
			}
			cf.Value = raw
			cf.Attrs["value"] = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}
		case PasswordField:
			raw := strings.TrimSpace(c.FormValue(cf.Name))
			errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
			if len(errs) > 0 {
				cf.Errors = errs
				f.Fields[i] = cf
				ok = false
				continue
			}
			cf.Value = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}

		case EmailField:
			raw := strings.TrimSpace(c.FormValue(cf.Name))
			errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
			if len(errs) > 0 {
				cf.Errors = errs
				f.Fields[i] = cf
				ok = false
				continue
			}
			cf.Value = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}

		case TelField:
			raw := strings.TrimSpace(c.FormValue(cf.Name))
			errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
			if len(errs) > 0 {
				cf.Errors = errs
				f.Fields[i] = cf
				ok = false
				continue
			}
			cf.Value = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}

		case URLField:
			raw := strings.TrimSpace(c.FormValue(cf.Name))
			errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
			if len(errs) > 0 {
				cf.Errors = errs
				f.Fields[i] = cf
				ok = false
				continue
			}
			cf.Value = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}

		case NumberField:
			raw := strings.TrimSpace(c.FormValue(cf.Name))
			errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
			if raw != "" {
				if _, err := strconv.ParseFloat(raw, 64); err != nil {
					errs = append(errs, cf.Label+" must be a number.")
				}
			}
			if len(errs) > 0 {
				cf.Errors = errs
				f.Fields[i] = cf
				ok = false
				continue
			}
			cf.Value = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}

		case DateField:
			raw := strings.TrimSpace(c.FormValue(cf.Name))
			errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
			if len(errs) > 0 {
				cf.Errors = errs
				f.Fields[i] = cf
				ok = false
				continue
			}
			cf.Value = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}

		case HiddenField:
			raw := c.FormValue(cf.Name)
			cf.Value = raw
			f.Fields[i] = cf
			if fv, exists := dest[cf.Name]; exists {
				setStructField(fv, raw)
			}
		/*
			case TextareaField:
				raw := strings.TrimSpace(c.FormValue(cf.Name))
				errs := validateRaw(cf.Label, cf.Required, cf.Attrs, raw)
				if len(errs) > 0 {
					cf.Errors = errs
					f.Fields[i] = cf
					ok = false
					continue
				}
				cf.Value = raw
				f.Fields[i] = cf
				if fv, exists := dest[cf.Name]; exists {
					setStructField(fv, raw)
				}

			case CheckboxField:
				raw := c.FormValue(cf.Name)
				checked := raw == "on" || raw == "true"
				cf.Checked = checked
				f.Fields[i] = cf
				if fv, exists := dest[cf.Name]; exists && fv.Kind() == reflect.Bool {
					fv.SetBool(checked)
				}

			case SelectField:
				raw := strings.TrimSpace(c.FormValue(cf.Name))
				if cf.Required && raw == "" {
					cf.Errors = []string{cf.Label + " is required."}
					f.Fields[i] = cf
					ok = false
					continue
				}
				for oi := range cf.Options {
					cf.Options[oi].Selected = cf.Options[oi].Value == raw
				}
				f.Fields[i] = cf
				if fv, exists := dest[cf.Name]; exists {
					setStructField(fv, raw)
				}
		*/

		default:
			continue
		}
	}
	return ok
}
