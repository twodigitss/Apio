package ui

import (
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/internal/core/parser"
	"github.com/twodigitss/apio/internal/core/parser/models"
)

// Ensure Model implements tea.Model
var _ tea.Model = Model{}

type Model struct {
	files        []os.DirEntry
	currentFile  []byte
	requests     []models.Tokens
	cursor       int
	selected     map[int]struct{}
	loading      bool
}

func InitialModel(dir []os.DirEntry, file []byte) Model {
	tokens, _ := core.FileToArrTokens(file)

	return Model{
		files: dir,
		currentFile: file,
		requests: tokens,
		cursor: 0,
		selected: make(map[int]struct{}),
		loading: false,
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

