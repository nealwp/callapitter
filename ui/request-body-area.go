package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
)

type RequestBodyArea struct {
	view *tview.TextView
    handler AppController
    request model.Request
}

func NewRequestBodyArea() *RequestBodyArea {
	view := tview.NewTextView()
	view.SetBorder(true)
	view.SetTitle("Body (b)")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetTextStyle(tcell.StyleDefault.Background(DEFAULT))

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

func (r *RequestBodyArea) SetRequest(req model.Request) {
    r.request = req
	pretty := PrettyPrintJSON(r.request.Body.String)
	r.view.SetText(pretty)
}

func (r *RequestBodyArea) setInputCapture() {
    keybinds := func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyRune {
            switch event.Rune(){
            case 'e': 
                r.handler.EditRequestBody(r.request)
                return nil
            }
        } 
        return nil
    }

    r.view.SetInputCapture(keybinds)
}
