package ssl

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 1. Prueba de Lógica Pura: Limpieza de Dominios
func TestCleanDomain(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"google.com", "google.com"},
		{"https://google.com", "google.com"},
		{"http://yahoo.com/", "yahoo.com"},
		{"   facebook.com   ", "facebook.com"},
	}

	for _, table := range tables {
		result := cleanDomain(table.input)
		if result != table.expected {
			t.Errorf("Entrada: %s, Esperado: %s, Obtenido: %s", table.input, table.expected, result)
		}
	}
}

// 2. Prueba de Integración Simulada (Mocking)
func TestAnalyze_Success(t *testing.T) {
	// A. Crear un servidor falso que simula ser SSL Labs
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("host") != "test.com" {
			t.Errorf("El scanner no envió el host correcto")
		}

		fakeResponse := SSLResult{
			Host:   "test.com",
			Status: "READY",
			Endpoints: []Endpoint{
				{IPAddress: "1.1.1.1", Grade: "A+"},
			},
		}
		json.NewEncoder(w).Encode(fakeResponse)
	}))
	defer mockServer.Close()

	// B. Configurar nuestro scanner para usar el servidor falso
	scanner := NewScanner()
	scanner.BaseURL = mockServer.URL

	// C. Ejecutar el análisis
	result, err := scanner.Analyze(context.Background(), "test.com", nil)

	// D. Validaciones (Asserts)
	if err != nil {
		t.Fatalf("No se esperaba error, pero ocurrió: %v", err)
	}
	if result.Status != "READY" {
		t.Errorf("Se esperaba status READY, se obtuvo %s", result.Status)
	}
	if result.Endpoints[0].Grade != "A+" {
		t.Errorf("Se esperaba grado A+, se obtuvo %s", result.Endpoints[0].Grade)
	}
}