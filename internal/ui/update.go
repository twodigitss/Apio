package ui

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/atotto/clipboard"

	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/internal/core/finder"
	"github.com/twodigitss/apio/internal/core/parser/lexer"
	"github.com/twodigitss/apio/internal/core/parser/models"
	"github.com/twodigitss/apio/internal/core/runner"
	"github.com/twodigitss/apio/internal/ui/data"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// ponytail: update sidebar model first so it can process and capture cursor navigation keypresses
	if !m.showHelp && !m.selectingFile {
		m.sidebar, cmd = m.sidebar.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	switch msg := msg.(type) {

	case data.RunResponseMsg:
		m.loading = false

		if msg.Err != nil {
			m.response = http.Response{}
			m.responseBody = fmt.Sprintf("Error: %v", msg.Err)
			m.viewport.SetContent(m.responseBody)
			m.viewport.GotoTop()
			return m, nil
		}

		contentType := msg.Response.Header.Get("Content-Type")

		if strings.Contains(contentType, "image/") ||
			strings.Contains(contentType, "application/octet-stream") {
			m.responseBody = fmt.Sprintf("[Binary content — %s]\n[%d bytes]",
				contentType, len(msg.Body))
		} else {
			m.responseBody = msg.Body
		}
		m.response = msg.Response
		// m.responseBody = msg.Body

		// ponytail: update viewport content with response details
		m.viewport.SetContent(fmt.Sprintf("Status: %s \nProtocol: %s\n\n%s\n", m.response.Status, m.response.Proto, m.responseBody))
		m.viewport.GotoTop()

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		// ponytail: adjust viewport size to fit the remaining space on the right side
		sidebarWidth := msg.Width / 3
		viewerWidth := msg.Width - sidebarWidth
		vpWidth := viewerWidth - 10
		vpHeight := msg.Height - 6 // Top padding: 2, Footer: 2, Safety margin: 2
		if vpWidth < 0 {
			vpWidth = 0
		}
		if vpHeight < 0 {
			vpHeight = 0
		}
		m.viewport.SetWidth(vpWidth)
		m.viewport.SetHeight(vpHeight)

	case tea.KeyPressMsg:

		if m.showHelp {
			if msg.String() == "?" || msg.String() == "h" || msg.String() == "esc" {
				m.showHelp = false
			}
			return m, nil
		}

		if m.selectingFile {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "esc", "f":
				m.selectingFile = false
				return m, nil
			case "up", "k":
				if m.fileCursor > 0 {
					m.fileCursor--
				}
				return m, nil
			case "down", "j":
				if m.fileCursor < len(m.files)-1 {
					m.fileCursor++
				}
				return m, nil
			case "enter":
				selectedFileEntry := m.files[m.fileCursor]
				fileBytes, err := finder.ReadFile(selectedFileEntry)
				if err == nil {
					tokens, err := lexer.FileToArrTokens(fileBytes)
					if err == nil {
						m.currentFile = fileBytes
						m.requests = tokens

						m.sidebar.Requests = tokens
						m.sidebar.Cursor = 0

						if len(tokens) > 0 {
							m.currentRequest = tokens[0]
							m.viewport.SetContent(m.currentRequest.Print())
						} else {
							m.currentRequest = models.Tokens{}
							m.viewport.SetContent("")
						}
						m.viewport.GotoTop()
					}
				}
				m.currentFile = fileBytes
				m.selectingFile = false
				return m, nil
			}
			return m, nil
		}

		switch msg.String() {
		case "?", "h":
			m.showHelp = !m.showHelp
			return m, nil

		case "f":
			if m.multipleFiles {
				m.selectingFile = true
			}
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k", "down", "j":
			if len(m.requests) > 0 {
				m.response = http.Response{}
				m.responseBody = ""
				m.currentRequest = m.requests[m.sidebar.Cursor]

				// ponytail: show the current selected request's source in the viewport
				m.viewport.SetContent(m.currentRequest.Print())
				m.viewport.GotoTop()
			}

		case "y":
			_ = clipboard.WriteAll(m.responseBody)

		case "enter":
			m.loading = true
			req := m.requests[m.sidebar.Cursor]

			return m, func() tea.Msg {
				res, err := runner.Run(req)
				if err != nil {
					return data.RunResponseMsg{Err: err}
				}
				var bodyBytes []byte
				if res.Body != nil {
					bodyBytes, _ = io.ReadAll(res.Body)
					res.Body.Close()
				}
				return data.RunResponseMsg{
					Response: res,
					Body:     string(bodyBytes),
				}
			}

		case "r":
			selectedFileEntry := m.files[m.fileCursor]
			fileBytes, err := finder.ReadFile(selectedFileEntry)
			if err != nil {
				return m, nil
			}

			reloadedRequests, err := finder.ReloadFiles(fileBytes)
			if err != nil {
				return m, nil
			}

			m.currentFile = fileBytes
			m.requests = reloadedRequests
			m.sidebar.Requests = reloadedRequests

			if m.sidebar.Cursor >= len(m.requests) {
				m.sidebar.Cursor = len(m.requests) - 1
			}
			if m.sidebar.Cursor < 0 {
				m.sidebar.Cursor = 0
			}

			m.currentRequest = m.requests[m.sidebar.Cursor]
			m.viewport.SetContent(m.currentRequest.Print())
			m.viewport.GotoTop()

		case "c":
			m.response.Body = nil
			m.response.StatusCode = 0
			m.responseBody = ""

			// ponytail: reset viewport back to showing the current request
			m.viewport.SetContent(m.currentRequest.Print())
			m.viewport.GotoTop()
		}

	}

	// ponytail: update spinner and viewport models to handle ticks, key bindings, and mouse events
	m.spinner, cmd = m.spinner.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	m.viewport, cmd = m.viewport.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	if len(cmds) == 0 {
		return m, nil
	}

	return m, tea.Batch(cmds...)
}
