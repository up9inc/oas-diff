package reporter

import (
	"bytes"
	"text/template"

	"github.com/up9inc/oas-diff/differentiator"
)

type htmlReporter struct {
	changelog *differentiator.ChangelogOutput
}

func NewHTMLReporter(changelog *differentiator.ChangelogOutput) Reporter {
	return &htmlReporter{
		changelog: changelog,
	}
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// TODO: Use the changelog data
func (h *htmlReporter) Build() ([]byte, error) {
	tmpl := template.Must(template.ParseFiles("reporter/template.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
			{Title: "Task 4", Done: false},
		},
	}
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
