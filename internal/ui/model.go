package ui

import (
	"net/http"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/internal/core/config"
	"github.com/twodigitss/apio/internal/core/parser/lexer"
	"github.com/twodigitss/apio/internal/core/parser/models"
	"github.com/twodigitss/apio/internal/ui/components/fileselection"
	"github.com/twodigitss/apio/internal/ui/components/sidebar"
	"github.com/twodigitss/apio/internal/ui/components/viewer"
	data "github.com/twodigitss/apio/internal/ui/data"
)

// Ensure Model implements tea.Model
var _ tea.Model = Model{}

// ponytail: cached once — config.Default() reads disk on every call
var cfg = config.Default()

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
		initialContent = currentRequest.PrintV2(cfg.UI.Glyphs)
	}

	return Model{
		selectingFile:  len(dir) > 1,
		currentRequest: currentRequest,
		response:       http.Response{},
		sidebar:        sidebar.New(tokens),
		fileSelection:  fileselection.New(dir),
		viewer:         viewer.New(initialContent, data.GetColorByHttpMethod(currentRequest.Method)),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.sidebar.Init(),
		m.fileSelection.Init(),
		m.viewer.Init(),
	)
}
