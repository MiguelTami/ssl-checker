package main

import (
	"context"
	"sync"
	"github.com/MiguelTami/ssl-checker/pkg/ssl"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
	cancelScan context.CancelFunc
	mu sync.Mutex
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AnalyzeDomain(domain string) (*ssl.SSLResult, error) {
	a.mu.Lock()
	ctx, cancel := context.WithCancel(context.Background())
	a.cancelScan = cancel
	a.mu.Unlock()

	progressChan := make(chan string)

	go func() {
		for msg := range progressChan {
			runtime.EventsEmit(a.ctx, "scanProgress", msg)
		}
	}()
	scanner := ssl.NewScanner()
	result, err := scanner.Analyze(ctx,domain, progressChan)
	close(progressChan)

	a.mu.Lock()
	a.cancelScan = nil
	a.mu.Unlock()
	
	return result, err
}

func (a *App) CancelAnalysis() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.cancelScan != nil {
		a.cancelScan() 
		
		runtime.EventsEmit(a.ctx, "scan-progress", "Cancelando operación...")
	}
}
