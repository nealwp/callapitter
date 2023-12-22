package ui

import (
    "fmt"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

type HttpRequest struct {
    Method string
    Endpoint string
    Headers []Header
    Body string
    LastResponse string
}

type RequestList struct {
    view *tview.List
}

func NewRequestList() *RequestList {
    view := tview.NewList()
    view.ShowSecondaryText(false)
    view.SetBorder(true)
    view.SetBackgroundColor(BG_COLOR)
    view.SetTitle("Requests [C-r]")
    view.SetTitleAlign(tview.AlignLeft)
    view.SetBorderPadding(1,1,1,1)

    reqListInputCapture := func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyRune {
            idx := view.GetCurrentItem()
            switch event.Rune() {
            case 'j':
                view.SetCurrentItem(idx+1)
            case 'k':
                view.SetCurrentItem(idx-1)
            }
            return event
        }
        return event
    }

    view.SetInputCapture(reqListInputCapture)
    return &RequestList{ view: view }
}

func (r *RequestList) GetPrimitive() tview.Primitive {
    return r.view
}

func (r *RequestList) SetContent(requests []HttpRequest) {
    for _, req := range requests {
        r.view.AddItem(fmt.Sprintf("%-4s", req.Method) + "  " + req.Endpoint, "", 0, nil)
    }
}




