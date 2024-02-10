package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	text := "Template start, Action \"{{.}}\",Template end\n"
	executeTemplate(text, 123)
	executeTemplate(text, "nikolas")
	executeTemplate(text, true)

	text = "start {{if .}}Dot is true!{{end}} end\n"
	executeTemplate(text, true)
	executeTemplate(text, false)

	text = "Before loop {{.}}\n{{range .}}In loop {{end}}\nAfter loop {{.}}\n"
	executeTemplate(text, []string{"1", "2", "3"})

	text = "Part type, Name: {{.Name}} - Count: {{.Count}}\n"
	executeTemplate(text, Part{Name: "Nikolas", Count: 100})
}
