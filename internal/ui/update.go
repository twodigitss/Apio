package ui

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/atotto/clipboard"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/core/config"
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

	if !m.showHelp && !m.selectingFile {
		m.sidebar, cmd = m.sidebar.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	if m.selectingFile {
		m.fileSelection, cmd = m.fileSelection.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	switch msg := msg.(type) {

	case data.RunResponseMsg:
		m.viewer.Loading = false

		if msg.Err != nil {
			m.response = http.Response{}
			m.responseBody = fmt.Sprintf("Error: %v", msg.Err)
			m.viewer.Viewport.SetContent(m.responseBody)
			m.viewer.Viewport.GotoTop()
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

		//some colors
		var Response string = "{}"
		if m.responseBody != "" {
			Response = m.responseBody
		}

		m.viewer.Viewport.SetContent(
			fmt.Sprintf("%s: %s \n%s: %s\n\n%s: %s\n\n%s:\n%s",
				lipgloss.NewStyle().Render(" Status"),
				lipgloss.NewStyle().Foreground(lipgloss.Color(data.ColorResponse(m.response.StatusCode))).Render(strings.TrimSpace(m.response.Status)),
				lipgloss.NewStyle().Bold(true).Render("󰿘 Protocol"),
				lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.SUBTEXT)).Render(strings.TrimSpace(m.response.Proto)),
				lipgloss.NewStyle().Bold(true).Render("Payload"),
				lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.SUBTEXT)).Render(strings.TrimSpace(Response)),
				lipgloss.NewStyle().Bold(true).Render("󰓹 Headers"),
				prettyHeaders(m.response.Header),
			),
		)
		m.viewer.Viewport.GotoTop()
		return m, nil

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

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
		m.viewer.Viewport.SetWidth(vpWidth)
		m.viewer.Viewport.SetHeight(vpHeight)

	case tea.KeyPressMsg:

		if m.showHelp {
			if msg.String() == "?" || msg.String() == "h" || msg.String() == "esc" || msg.String() == "q" {
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
			case "enter":
				selectedFileEntry := m.fileSelection.Files[m.fileSelection.FileCursor]
				fileBytes, err := finder.ReadFile(selectedFileEntry)
				if err == nil {
					tokens, err := lexer.FileToArrTokens(fileBytes)
					if err == nil {
						m.sidebar.Requests = tokens
						m.sidebar.Cursor = 0
						if len(tokens) > 0 {
							m.currentRequest = tokens[0]
							m.viewer.Viewport.SetContent(m.currentRequest.PrintV2())
						} else {
							m.currentRequest = models.Tokens{}
							m.viewer.Viewport.SetContent("")
						}
						m.viewer.Viewport.GotoTop()
					}
				}
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
			if len(m.fileSelection.Files) > 1 {
				m.selectingFile = true
			}
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k", "down", "j":
			if len(m.sidebar.Requests) > 0 {
				m.response = http.Response{}
				m.responseBody = ""
				m.currentRequest = m.sidebar.Requests[m.sidebar.Cursor]

				m.viewer.Viewport.SetContent(m.currentRequest.PrintV2())
				m.viewer.Viewport.GotoTop()
			}

		case "y":
			_ = clipboard.WriteAll(m.responseBody)

		case "enter":
			m.viewer.Loading = true
			req := m.sidebar.Requests[m.sidebar.Cursor]

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
			selectedFileEntry := m.fileSelection.Files[m.fileSelection.FileCursor]
			fileBytes, err := finder.ReadFile(selectedFileEntry)
			if err != nil {
				return m, nil
			}

			reloadedRequests, err := lexer.FileToArrTokens(fileBytes)
			if err != nil {
				return m, nil
			}

			m.sidebar.Requests = reloadedRequests

			if m.sidebar.Cursor >= len(m.sidebar.Requests) {
				m.sidebar.Cursor = len(m.sidebar.Requests) - 1
			}
			if m.sidebar.Cursor < 0 {
				m.sidebar.Cursor = 0
			}

			m.currentRequest = m.sidebar.Requests[m.sidebar.Cursor]
			m.viewer.Viewport.SetContent(m.currentRequest.PrintV2())
			m.viewer.Viewport.GotoTop()

		case "c":
			m.response.Body = nil
			m.response.StatusCode = 0
			m.responseBody = ""

			m.viewer.Viewport.SetContent(m.currentRequest.PrintV2())
			m.viewer.Viewport.GotoTop()
		}

	}

	m.viewer, cmd = m.viewer.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	if len(cmds) == 0 {
		return m, nil
	}

	return m, tea.Batch(cmds...)
}

func prettyHeaders(h http.Header) string {
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		for _, v := range h[k] {
			fmt.Fprintf(&sb, " • %s: %s\n",
				lipgloss.NewStyle().Render(k),
				lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.INFRATEXT)).Italic(true).Render(v),
			)
		}
	}
	return sb.String()
}
