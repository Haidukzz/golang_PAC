package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Log para verificar se o código está sendo executado até aqui
	log.Println("Iniciando o servidor...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Caso a variável de ambiente 'PORT' não esteja definida, use 8080 como padrão
	}

	// Log para mostrar qual porta será utilizada
	log.Printf("Escutando na porta: %s\n", port)

	// Configura a rota e inicia o servidor
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %s", err) // Log de erro se falhar ao iniciar
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando a rota principal...")
	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("Erro ao renderizar template: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
