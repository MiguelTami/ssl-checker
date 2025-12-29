package ssl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	pollInterval = 10 * time.Second
	maxRetries   = 3 
)

type Scanner struct {
	client *http.Client
	BaseURL string
}

func NewScanner() *Scanner {
	return &Scanner{
		client: &http.Client{Timeout: 30 * time.Second}, 
		BaseURL: "https://api.ssllabs.com/api/v2/analyze",
	}
}

// Analyze gestiona el ciclo de vida completo del análisis (Polling)
// Acepta un canal 'progress' para notificar a la GUI o Consola qué está pasando
func (s *Scanner) Analyze(domain string, progress chan<- string) (*SSLResult, error) {
	domain = cleanDomain(domain)
	
	url := fmt.Sprintf("%s?host=%s&all=done", s.BaseURL, domain)
	failCount := 0

	for {
		if progress != nil {
			progress <- "Consultando estado..."
		}

		result, err := s.fetchOne(url)
		
		if err != nil {
			failCount++
			if failCount >= maxRetries {
				return nil, fmt.Errorf("fallo de red persistente: %v", err)
			}
			if progress != nil {
				progress <- fmt.Sprintf("Error de red (%d/%d), reintentando...", failCount, maxRetries)
			}
			time.Sleep(5 * time.Second)
			continue
		}
		
		failCount = 0

		switch result.Status {
		case "READY":
			if progress != nil {
				progress <- "Análisis finalizado."
			}
			return result, nil
		case "ERROR":
			return nil, errors.New("SSL Labs reportó un error interno al analizar el dominio")
		case "DNS":
			if progress != nil {
				progress <- "Resolviendo DNS..."
			}
		default:
			if progress != nil {
				progress <- "Analizando protocolos y certificados..."
			}
		}

		time.Sleep(pollInterval)
	}
}

func (s *Scanner) fetchOne(url string) (*SSLResult, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Go-SSL-Checker-Pro/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API retornó código: %d", resp.StatusCode)
	}

	var result SSLResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func cleanDomain(d string) string {
	d = strings.TrimSpace(d)
	d = strings.TrimPrefix(d, "https://")
	d = strings.TrimPrefix(d, "http://")
	d = strings.TrimSuffix(d, "/")
	return d
}