package ui

import (
	"fmt"
	"io"
	"log"
	"net/http"

	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/internal/core/runner"
	"github.com/twodigitss/apio/internal/ui/data"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case data.RunResponseMsg:
		m.loading = false
		if msg.Err != nil {
			// FIX: do something with this
			log.Fatal("Error running block:", msg.Err)
		}
		m.response = msg.Response
		m.responseBody = msg.Body

		// ponytail: update viewport content with response details
		m.viewport.SetContent(fmt.Sprintf("Status Code: %d \n\n%s\n", m.response.StatusCode, m.responseBody))
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
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

			//in view.go i use statuscode=0 to unrender the response
			//but since i set the response to the interface, statuscode
			//is already zero. dont forget
			m.response = http.Response{}
			m.responseBody = ""
			m.currentRequest = m.requests[m.cursor]

			// ponytail: show the current selected request's source in the viewport
			m.viewport.SetContent(m.currentRequest.Print())
			m.viewport.GotoTop()

		case "down", "j":
			if m.cursor < len(m.requests)-1 {
				m.cursor++
			}

			m.response = http.Response{}
			m.responseBody = ""
			m.currentRequest = m.requests[m.cursor]

			// ponytail: show the current selected request's source in the viewport
			m.viewport.SetContent(m.currentRequest.Print())
			m.viewport.GotoTop()

		case "enter":
			m.loading = true
			req := m.requests[m.cursor]

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
