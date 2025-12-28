package main

import (
	"github.com/MiguelTami/ssl-checker/internal/ssl"
	"flag"
	"fmt"
	"os"
)

func main() {
	var domain string
	flag.StringVar(&domain, "d", "", "El dominio a analizar")
	flag.Parse()

	if domain == "" {
		if len(os.Args) > 1 {
			domain = os.Args[1]
		} else {
			fmt.Println("Uso: go run main.go -d google.com")
			return
		}
	}

	// Canal para recibir actualizaciones de estado 
	progressChan := make(chan string)
	
	// Goroutine para imprimir los mensajes del canal
	go func() {
		for msg := range progressChan {
			fmt.Printf("\r>> Estado: %-50s", msg) 
		}
	}()

	scanner := ssl.NewScanner()
	fmt.Printf("--- Iniciando análisis para: %s ---\n", domain)

	result, err := scanner.Analyze(domain, progressChan)
	close(progressChan) 
	fmt.Println()       

	if err != nil {
		fmt.Printf("\n❌ Error Crítico: %v\n", err)
		os.Exit(1)
	}

	printReport(result)
}

func printReport(result *ssl.SSLResult) {
	fmt.Println("\n-----------------------------------")
	fmt.Printf("RESULTADO FINAL: %s\n", result.Host)
	for _, ep := range result.Endpoints {
		fmt.Printf("IP: %s -> Grado: %s\n", ep.IPAddress, ep.Grade)
	}
}