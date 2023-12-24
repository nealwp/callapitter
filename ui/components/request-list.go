package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HttpRequest struct {
	Method       string
	Endpoint     string
	Headers      []Header
	Body         string
	LastResponse string
}

type RequestList struct {
	view *tview.List
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

	return &RequestList{view: view}
}

func (r *RequestList) GetPrimitive() tview.Primitive {
	return r.view
}

func (r *RequestList) SetContent(requests []HttpRequest) {
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
