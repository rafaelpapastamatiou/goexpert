package main

import (
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	curso := Curso{Nome: "GoExpert", CargaHoraria: 80}
	curso2 := Curso{Nome: "Java", CargaHoraria: 40}
	curso3 := Curso{Nome: "NodeJS", CargaHoraria: 120}

	t := template.New("template.html")

	t.Funcs(template.FuncMap{
		"ToUpper": ToUpper,
	})

	t = template.Must(t.ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		curso,
		curso2,
		curso3,
	})

	if err != nil {
		panic(err)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
