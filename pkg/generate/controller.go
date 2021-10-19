package generate

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type ControllerTemplateData struct {
	Name      string
	NameTitle string
	Body      bool
}

type ControllerIndexTemplateData struct {
	Actions []string
}

func GenController(name string, routes []string, body bool) {
	t, e := template.ParseFiles("templates/controllers/controller.ts")
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	os.Mkdir("controllers", os.ModePerm)
	// Generate index.ts
	fileDir := fmt.Sprintf("controllers/%v-controllers", name)
	if _, err := os.Stat(fileDir); os.IsExist(err) {
		fmt.Fprint(os.Stderr, "Dir exist")
		os.Exit(1)
	}
	os.Mkdir(fileDir, os.ModePerm)
	indexControllerWriter, _ := os.Create(fileDir + "/index.ts")
	indexControllerTemplate, _ := template.ParseFiles("templates/controllers/index.ts")
	e = indexControllerTemplate.Execute(indexControllerWriter, ControllerIndexTemplateData{
		Actions: routes,
	})
	indexControllerWriter.Close()
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}

	for _, controller := range routes {
		filePath := fmt.Sprintf("%v/%v.ts", fileDir, controller)
		if _, err := os.Stat(filePath); os.IsExist(err) {
			fmt.Fprint(os.Stderr, "File exist")
			os.Exit(1)
		}
		f, e := os.Create(filePath)
		if e != nil {
			f.Close()
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		}
		e = t.Execute(f, ControllerTemplateData{
			Name:      controller,
			NameTitle: strings.Title(controller),
			Body:      body,
		})
		f.Close()
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		}
	}

}
