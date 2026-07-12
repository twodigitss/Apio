package ui

import (
	"net/http"
	"os"

	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/internal/core/parser/lexer"
	"github.com/twodigitss/apio/internal/core/parser/models"
)

// Ensure Model implements tea.Model
var _ tea.Model = Model{}

type Model struct {
	files          []os.DirEntry
	multipleFiles  bool
	selectingFile  bool
	fileCursor     int
	currentFile    []byte
	requests       []models.Tokens
	currentRequest models.Tokens

	response     http.Response
	responseBody string

	cursor   int
	loading  bool
	Width    int
	Height   int
	spinner  spinner.Model
	viewport viewport.Model
	showHelp bool
}

func InitialModel(dir []os.DirEntry, file []byte) Model {
	tokens, _ := lexer.FileToArrTokens(file)
	s := spinner.New()
	s.Spinner = spinner.Dot
	// s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#cfe4ef"))

	var currentRequest models.Tokens
	if len(tokens) > 0 {
		currentRequest = tokens[0]
	}

	// ponytail: configure viewport pager with custom keymap to avoid conflict with main menu navigation
	vp := viewport.New()
	vp.FillHeight = true
	vp.SoftWrap = true
	vp.KeyMap = viewport.DefaultKeyMap()
	vp.KeyMap.Down.SetKeys("ctrl+j")
	vp.KeyMap.Up.SetKeys("ctrl+k")
	vp.MouseWheelEnabled = true
	if len(tokens) > 0 {
		vp.SetContent(currentRequest.Print())
	}

	multiplefiles := false
	if len(dir) > 1 {
		multiplefiles = true
	}

	return Model{
		files:          dir,
		currentFile:    file,
		multipleFiles:  multiplefiles,
		selectingFile:  multiplefiles,
		fileCursor:     0,
		requests:       tokens,
		currentRequest: currentRequest,
		cursor:         0,
		response:       http.Response{},
		loading:        false,
		spinner:        s,
		viewport:       vp,
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// return nil
	return m.spinner.Tick
}
