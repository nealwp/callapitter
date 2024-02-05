package ui

import (
	"github.com/rivo/tview"
)

type StatusBar struct {
	view *tview.TextView
}

func NewStatusBar() *StatusBar {
	view := tview.NewTextView()
	view.SetBorder(true)
	view.SetBackgroundColor(BG_COLOR)
	view.SetTitle("Status")
	view.SetTitleAlign(tview.AlignLeft)

	return &StatusBar{view: view}
}

func (s *StatusBar) GetPrimitive() tview.Primitive {
	return s.view
}

func (s *StatusBar) SetStatus(msg string) {
	s.view.SetText(msg)
}
