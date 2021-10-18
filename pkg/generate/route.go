package generate

import (
	"fmt"
	"os"
	"text/template"
)

type RouterTemplateData struct {
	Name     string
	Validate bool
	Routes   []string
}

func GenRoute(Name string, Routes []string, validate bool) {
	fileName := "routes/" + Name + "-routes.ts"
	if _, err := os.Stat(fileName); os.IsExist(err) {
		fmt.Fprint(os.Stderr, "File exist")
		os.Exit(1)
	}
	f, _ := os.Create(fileName)
	t, e := template.ParseFiles("templates/route.ts")
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	e = t.Execute(f, RouterTemplateData{
		Name:     Name,
		Validate: validate,
		Routes:   Routes,
	})
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
