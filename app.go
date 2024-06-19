package main

import (
	"context"
	"fmt"

	"github.com/stargazer39/file-tagger/tagger"
)

// App struct
type App struct {
	ctx    context.Context
	tagger tagger.Tagger
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.tagger = tagger.NewTagger()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetFileList() (*[]tagger.TaggedFile, error) {
	return a.tagger.ListFiles(".")
}
