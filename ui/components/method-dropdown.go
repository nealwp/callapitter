package ui

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type MethodDropdown struct {
    view *tview.DropDown
}

var reqMethods = []string{"GET", "POST", "PUT", "DELETE"}

func NewMethodDropdown() *MethodDropdown {

    view := tview.NewDropDown()
    view.SetOptions(reqMethods, nil)
    view.SetCurrentOption(0)
    view.SetFieldBackgroundColor(BG_COLOR)
    view.SetFieldTextColor(BG_COLOR)
    view.SetTitle("Method [C-e]")
    view.SetTitleAlign(tview.AlignLeft)
    view.SetBackgroundColor(BG_COLOR)
    view.SetBorder(true)
    view.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

    return &MethodDropdown{ view: view }
}

func (m *MethodDropdown) GetPrimitive() tview.Primitive {
    return m.view
}

func (m *MethodDropdown) SetCurrentOption(method string) {
    index := findMethodIndex(method)
    m.view.SetCurrentOption(index)
}

func findMethodIndex(method string) int {
    for i, v := range reqMethods {
        if v == method {
            return i
        }
    }
    return -1
}
