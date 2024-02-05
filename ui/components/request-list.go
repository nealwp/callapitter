package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
)

type Handler interface {
    DeleteRequest(index int)
    SelectRequest(index int)
    CreateRequest()
    SendRequest(index int)
}

type RequestList struct {
	view *tview.List
    handler Handler
}

func NewRequestList() *RequestList {

	title := "Requests"

	view := tview.NewList()
	view.ShowSecondaryText(false)
	view.SetBorder(true)
	view.SetBackgroundColor(BG_COLOR)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)
	view.SetBorderPadding(1, 1, 1, 1)

    r := &RequestList{view: view}
    r.setKeybindings()
	return r 
}

func (r *RequestList) GetPrimitive() tview.Primitive {
	return r.view
}

func (r *RequestList) SetHandler(handler Handler) {
    r.handler = handler
}

func (r *RequestList) SetContent(requests []model.Request) {
    r.view.Clear()
	for _, req := range requests {
		r.view.AddItem(fmt.Sprintf("%-4s", req.Method)+"  "+req.Endpoint, "", 0, nil)
	}
}

func (r *RequestList) SetChangedFunc(f func(index int, mainText, secondaryText string, shortcut rune)) {
	r.view.SetChangedFunc(f)
}

func (r *RequestList) SetSelectedFunc(f func(index int, mainText, secondaryText string, shortcut rune)) {
	r.view.SetSelectedFunc(f)
}

func (r *RequestList) SetInputCapture(f func(event *tcell.EventKey) *tcell.EventKey) {
	r.view.SetInputCapture(f)
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
                r.handler.SelectRequest(index+1)
            case 'k':
                r.handler.SelectRequest(index-1)
            case '%':
                r.handler.CreateRequest()
            }
            return event
        } else if event.Key() == tcell.KeyEnter {
            r.handler.SendRequest(index)
        }
        return event
    }

    r.SetInputCapture(keybinds)
}
