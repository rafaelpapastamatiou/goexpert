package main

import (
	"os"
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

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		curso,
		curso2,
		curso3,
	})

	if err != nil {
		panic(err)
	}
}
