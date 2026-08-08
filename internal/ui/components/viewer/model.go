package viewer

import (
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/viewport"
	"charm.land/lipgloss/v2"
)

type Model struct {
	Spinner  spinner.Model
	Loading  bool
	Viewport viewport.Model
}

func New(initialContent string, color string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	vp := viewport.New()
	vp.FillHeight = true
	vp.SoftWrap = true
	vp.KeyMap = viewport.DefaultKeyMap()
	vp.KeyMap.Down.SetKeys("ctrl+j")
	vp.KeyMap.Up.SetKeys("ctrl+k")
	vp.MouseWheelEnabled = true
	if initialContent != "" {
		vp.SetContent(initialContent)
	}

	return Model{
		Spinner:  s,
		Viewport: vp,
		Loading:  false,
	}
}

func (m *Model) SetColor(color string) {
	m.Spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
}
