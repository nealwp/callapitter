package ui

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type RequestBodyArea struct {
   view *tview.TextArea
}

func NewRequestBodyArea() *RequestBodyArea {
    view := tview.NewTextArea()
    view.SetBackgroundColor(BG_COLOR)
    view.SetBorder(true)
    view.SetTitle("Body")
    view.SetTitleAlign(tview.AlignLeft)
    view.SetTextStyle(tcell.StyleDefault.Background(BG_COLOR))

    return &RequestBodyArea{ view: view }
}

func (r *RequestBodyArea) GetPrimitive() tview.Primitive {
    return r.view
}

// TODO: make this not editable, open $EDITOR to modify instead 
func (r *RequestBodyArea) SetText(text string) {
    pretty := PrettyPrintJSON(text)
    r.view.SetText(pretty, false)
}
