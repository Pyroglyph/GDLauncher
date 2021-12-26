package main

import (
	"context"
	"fmt"
	"time"
)

// App application struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (b *App) Greet(name string) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("Hello 2 %s!", name)
}

// Greet returns a greeting for the given name
func (b *App) LaunchInstance(name string) string {
	time.Sleep(time.Second * 4)
	return fmt.Sprint("Launch later 4", "Hi")
}