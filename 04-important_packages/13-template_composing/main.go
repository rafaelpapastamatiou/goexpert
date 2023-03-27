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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	cursoGo := Curso{Nome: "GoExpert", CargaHoraria: 80}
	cursoTs := Curso{Nome: "TypeScript", CargaHoraria: 150}
	cursoJava := Curso{Nome: "Java", CargaHoraria: 110}

	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		cursoGo,
		cursoTs,
		cursoJava,
	})

	if err != nil {
		panic(err)
	}
}
