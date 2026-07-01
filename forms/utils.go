package forms

import (
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Humanize(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	s = regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`).ReplaceAllString(s, "$1 $2")
	s = regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(s, "$1 $2")
	return cases.Title(language.English).String(s)
}

func Slugify(sf reflect.StructField) string {
	if j, ok := sf.Tag.Lookup("json"); ok {
		if name := strings.Split(j, ",")[0]; name != "" && name != "-" {
			return name
		}
	}
	return strings.ToLower(camelRe.ReplaceAllString(sf.Name, "${1}_${2}"))
}
