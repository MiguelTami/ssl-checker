package main

import (
	"context"
	"github.com/MiguelTami/ssl-checker/pkg/ssl"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// AnalyzeDomain returns a greeting for the given name
func (a *App) AnalyzeDomain(domain string) (*ssl.SSLResult, error) {
	scanner := ssl.NewScanner()
	return scanner.Analyze(domain, nil)
}
