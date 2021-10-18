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
	f, e := os.Create(fileName)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	t, e := template.ParseFiles("templates/route.ts")
	if e != nil {
		f.Close()
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	e = t.Execute(f, RouterTemplateData{
		Name:     Name,
		Validate: validate,
		Routes:   Routes,
	})
	f.Close()
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
