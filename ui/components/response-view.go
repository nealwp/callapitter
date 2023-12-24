package ui

import (
    "github.com/rivo/tview"
)

type ResponseView struct {
    view *tview.TextView
}

func NewResponseView() *ResponseView {

    title := "Response"

    view := tview.NewTextView()
    view.SetTitle(title)
    view.SetTitleAlign(tview.AlignLeft)
    view.SetBackgroundColor(BG_COLOR)
    view.SetBorder(true)
   
    return &ResponseView{ view: view }
}

func (r *ResponseView) GetPrimitive() tview.Primitive {
    return r.view
}

func (r *ResponseView) SetContent(text string) {
    pretty := PrettyPrintJSON(text)
    r.view.SetText(pretty)
}
