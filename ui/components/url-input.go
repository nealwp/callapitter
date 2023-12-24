package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UrlInput struct {
	view *tview.InputField
}

func NewUrlInput() *UrlInput {
	title := "URL"

	view := tview.NewInputField()
	view.SetFieldBackgroundColor(BG_COLOR)
	view.SetFieldTextColor(BG_COLOR)
	view.SetBackgroundColor(BG_COLOR)
	view.SetBorder(true)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)

	urlInputCapture := func(event *tcell.EventKey) *tcell.EventKey {
		return event
	}

	view.SetInputCapture(urlInputCapture)

	return &UrlInput{view: view}
}

func (u *UrlInput) GetPrimitive() tview.Primitive {
	return u.view
}

func (u *UrlInput) SetText(text string) {
	u.view.SetText(text)
}
