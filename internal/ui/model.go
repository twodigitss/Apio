package ui

import (
	"net/http"
	"os"

	tea "charm.land/bubbletea/v2"
	core "github.com/twodigitss/apio/internal/core/parser"
	"github.com/twodigitss/apio/internal/core/parser/models"
)

// Ensure Model implements tea.Model
var _ tea.Model = Model{}

type Model struct {
	files          []os.DirEntry
	currentFile    []byte
	requests       []models.Tokens
	currentRequest models.Tokens

	response     http.Response
	responseBody string // ponytail: cached response body string to avoid double-reads and nil pointer panics

	cursor int
	// selected int
	loading bool
	Width   int
	Height  int
}

func InitialModel(dir []os.DirEntry, file []byte) Model {
	tokens, _ := core.FileToArrTokens(file)

	return Model{
		files:       dir,
		currentFile: file,
		requests:    tokens,
		cursor:      0,
		response:    http.Response{},
		// selected:    0,
		loading: false,
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
