package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursoGo := Curso{Nome: "GoExpert", CargaHoraria: 80}
		cursoTs := Curso{Nome: "TypeScript", CargaHoraria: 150}

		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, Cursos{
			cursoGo,
			cursoTs,
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
