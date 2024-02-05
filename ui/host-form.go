package ui

import (
	"github.com/rivo/tview"
)

type HostForm struct {
	view *tview.Form
}

func NewRequestForm() *HostForm {

	title := "Add Host"

	view := tview.NewForm()
	view.AddInputField("Hostname", "", 20, nil, nil)
	view.SetBorder(true)
	view.SetBackgroundColor(BG_COLOR)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)
	view.SetBorderPadding(1, 1, 1, 1)

	return &HostForm{view: view}
}

func (r *HostForm) GetPrimitive() tview.Primitive {
	return r.view
}
