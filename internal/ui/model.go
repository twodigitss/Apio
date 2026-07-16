package ui

import (
	"net/http"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/internal/core/parser/lexer"
	"github.com/twodigitss/apio/internal/core/parser/models"
	"github.com/twodigitss/apio/internal/ui/components/fileselection"
	"github.com/twodigitss/apio/internal/ui/components/sidebar"
	"github.com/twodigitss/apio/internal/ui/components/viewer"
)

// Ensure Model implements tea.Model
var _ tea.Model = Model{}

type Model struct {
	selectingFile  bool
	currentRequest models.Tokens

	response     http.Response
	responseBody string

	Width    int
	Height   int
	showHelp bool

	sidebar       sidebar.Model
	fileSelection fileselection.Model
	viewer        viewer.Model
}

func New(dir []os.DirEntry, file []byte) Model {
	tokens, _ := lexer.FileToArrTokens(file)

	var currentRequest models.Tokens
	if len(tokens) > 0 {
		currentRequest = tokens[0]
	}

	initialContent := ""
	if len(tokens) > 0 {
		initialContent = currentRequest.Print()
	}

	multipleFiles := len(dir) > 1

	return Model{
		selectingFile:  multipleFiles,
		currentRequest: currentRequest,
		response:       http.Response{},
		sidebar:        sidebar.New(tokens),
		fileSelection:  fileselection.New(dir),
		viewer:         viewer.New(initialContent),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.sidebar.Init(),
		m.fileSelection.Init(),
		m.viewer.Init(),
	)
}
