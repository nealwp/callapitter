package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HostDropdown struct {
    view *tview.DropDown
}

var hosts = []string{"http://localhost:8000", "https://jsonplaceholder.typicode.com"}

func NewHostDropdown() *HostDropdown {

    view := tview.NewDropDown()
    view.SetOptions(hosts, nil)
    view.SetCurrentOption(0)
    view.SetFieldBackgroundColor(BG_COLOR)
    view.SetFieldTextColor(BG_COLOR)
    view.SetTitle("Host [C-h]")
    view.SetTitleAlign(tview.AlignLeft)
    view.SetBackgroundColor(BG_COLOR)
    view.SetBorder(true)
    view.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

    return &HostDropdown{ view: view }
}

func (h *HostDropdown) GetPrimitive() tview.Primitive {
    return h.view
}
