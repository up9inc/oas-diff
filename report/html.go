package report

import (
	"bytes"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func RenderReport() ([]byte, error) {
	tmpl := template.Must(template.ParseFiles("report/template.html"))
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
