package ui

import (
	"github.com/rivo/tview"
)

type StatusBar struct {
	view *tview.InputField
}

func NewStatusBar() *StatusBar {
	view := tview.NewInputField()
	view.SetBorder(true)
    view.SetDisabled(true)
    view.SetBackgroundColor(BG_COLOR)
	view.SetFieldBackgroundColor(BG_COLOR)
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
