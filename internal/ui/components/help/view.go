package help

import "charm.land/lipgloss/v2"

func View() string {
	content := `
  Navigation
  j / ↓     Move down
  k / ↑     Move up
  
  Viewport
  ctrl+j    Scroll down
  ctrl+k    Scroll up
  
  Actions
  enter     Execute request
  r         Reload file
  c         Clear response
  f         Select file
  h / ?     Toggle help
  q         Quit
  y         Copy response body
`
	return lipgloss.NewStyle().
		Padding(1, 3).
		Render(content)
}
