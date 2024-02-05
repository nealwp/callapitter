package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
)

type UrlInput struct {
	view    *tview.InputField
	handler AppController
	request model.Request
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

	u := &UrlInput{view: view}
	u.setDoneFunc()
	u.setInputCapture()

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

func (u *UrlInput) setInputCapture() {
	u.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})
}

func (u *UrlInput) setDoneFunc() {
	u.view.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			u.request.Endpoint = u.view.GetText()
			u.handler.UpdateRequest(u.request)
		}
	})
}
