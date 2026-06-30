package forms

import "github.com/flosch/pongo2/v6"

var templateSet = pongo2.NewSet(
	"templates",
	pongo2.MustNewLocalFileSystemLoader("forms/templates"),
)

func renderTPL(name string, ctx pongo2.Context) (string, error) {
	tpl, err := templateSet.FromFile(name)
	if err != nil {
		return "", err
	}

	out, err := tpl.Execute(ctx)
	if err != nil {
		return "", err
	}

	return out, nil
}
