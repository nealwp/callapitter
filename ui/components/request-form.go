package ui

import (
	"github.com/rivo/tview"
)

type RequestForm struct {
	view *tview.Form
}

func NewRequestForm() *RequestForm {

	title := "Add Request"

	view := tview.NewForm()
    view.AddInputField("URL", "", 20, nil, nil)
    view.AddDropDown("Method", []string{"GET", "DELETE", "POST", "PUT"}, 0, nil)
	view.SetBorder(true)
	view.SetBackgroundColor(BG_COLOR)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)
	view.SetBorderPadding(1, 1, 1, 1)

	return &RequestForm{view: view}
}

func (r *RequestForm) GetPrimitive() tview.Primitive {
	return r.view
}
