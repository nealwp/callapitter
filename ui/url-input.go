package ui

import (
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
)

type UrlInput struct {
	view    *tview.TextView
	handler AppController
	request model.Request
}

func NewUrlInput() *UrlInput {
	title := "URL"

	view := tview.NewTextView()
	view.SetBackgroundColor(BG_COLOR)
	view.SetBorder(true)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)

	u := &UrlInput{view: view}

	return u
}

func (u *UrlInput) GetPrimitive() tview.Primitive {
	return u.view
}

func (u *UrlInput) SetText(req model.Request) {
	u.request = req
	u.view.SetText(req.Endpoint)
}

func (u *UrlInput) Bind(handler AppController) {
	u.handler = handler
}
