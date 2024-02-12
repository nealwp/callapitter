package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
)

type RequestList struct {
	view    *tview.List
	handler AppController
}

func NewRequestList() *RequestList {

	title := "Requests (r)"

	view := tview.NewList()
	view.ShowSecondaryText(false)
    view.SetSelectedTextColor(tcell.ColorLime)
    view.SetSelectedBackgroundColor(DEFAULT)
	view.SetBorder(true)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)
	view.SetBorderPadding(1, 1, 1, 1)
    
	r := &RequestList{view: view}

	r.setKeybindings()
	r.setChangedFunc()
	return r
}

func (r *RequestList) GetPrimitive() tview.Primitive {
	return r.view
}

func (r *RequestList) Bind(handler AppController) {
	r.handler = handler
}

func (r *RequestList) SetContent(requests []model.Request) {
	r.view.Clear()
	for _, req := range requests {
		r.view.AddItem(fmt.Sprintf("%-4s", req.Method)+"  "+req.Endpoint, "", 0, nil)
	}
}

func (r *RequestList) setChangedFunc() {
	r.view.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		r.handler.HandleRequestSelected(index)
	})
}

func (r *RequestList) GetSelectedRequest() int {
	return r.view.GetCurrentItem()
}

func (r *RequestList) SetSelectedRequest(index int) {
	r.view.SetCurrentItem(index)
}

func (r *RequestList) setKeybindings() {

	keybinds := func(event *tcell.EventKey) *tcell.EventKey {
		index := r.view.GetCurrentItem()
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case 'D':
				r.handler.DeleteRequest(index)
			case 'j':
				r.handler.SelectRequest(index + 1)
			case 'k':
				r.handler.SelectRequest(index - 1)
            case 'R':
                r.handler.EditRequest(index)
			case '%':
				r.handler.AddRequest()
			}
			return event
		} else if event.Key() == tcell.KeyEnter {
			r.handler.SendRequest(index)
		}
		return event
	}

	r.view.SetInputCapture(keybinds)
}
