package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Obtenha a porta a partir da variável de ambiente 'PORT'
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Caso a variável de ambiente 'PORT' não esteja definida, use 8080 como padrão
	}

	// Inicia o servidor na porta determinada pela variável de ambiente 'PORT'
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":"+port, nil)) // Log de erro se falhar ao iniciar o servidor
}

func index(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
