package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type RequestBodyArea struct {
	view *tview.TextArea
    handler AppController
}

func NewRequestBodyArea() *RequestBodyArea {
	view := tview.NewTextArea()
	view.SetBackgroundColor(BG_COLOR)
	view.SetBorder(true)
	view.SetTitle("Body")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetTextStyle(tcell.StyleDefault.Background(BG_COLOR))

    r := &RequestBodyArea{view: view}
    r.setInputCapture()

    return r
}

func (r *RequestBodyArea) GetPrimitive() tview.Primitive {
	return r.view
}

func (r *RequestBodyArea) Bind(handler AppController) {
    r.handler = handler
}

func (r *RequestBodyArea) setInputCapture() {
    keybinds := func(event *tcell.EventKey) *tcell.EventKey {

        if event.Key() == tcell.KeyCtrlE {
            r.handler.EditRequestBody(r.view.GetText())
            r.handler.AppSync()
            return nil
        }

        return event
    }

    r.view.SetInputCapture(keybinds)
}

// TODO: make this not editable, open $EDITOR to modify instead
func (r *RequestBodyArea) SetText(text string) {
    // TODO: should content come in already formatted?
	pretty := PrettyPrintJSON(text)
	r.view.SetText(pretty, false)
}
