package sidebar

import (
	"github.com/twodigitss/apio/internal/core/parser/models"
)

type Model struct {
	Requests []models.Tokens
	Cursor   int
	Width    int
	Height   int
}

func New(requests []models.Tokens) Model {
	return Model{
		Requests: requests,
		Cursor:   0,
	}
}
