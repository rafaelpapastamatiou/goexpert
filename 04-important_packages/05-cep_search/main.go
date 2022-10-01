package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Response error: %v\n", err)
		}

		defer res.Body.Close()

		content, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when reading response body: %v\n", err)
		}

		var data ViaCep

		err = json.Unmarshal(content, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when parsing response body to ViaCep struct: %v\n", err)
		}

		file, err := os.Create("city.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when creating file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when writing to file: %v\n", err)
		}

		println("File created successfully!")
	}
}
