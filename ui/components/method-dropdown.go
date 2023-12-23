package ui

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type MethodDropdown struct {
    view *tview.DropDown
}

var methods = []string{"GET", "POST", "PUT", "DELETE"}

func NewMethodDropdown() *MethodDropdown {

    title := "Method"

    view := tview.NewDropDown()
    view.SetOptions(methods, nil)
    view.SetCurrentOption(0)
    view.SetFieldBackgroundColor(BG_COLOR)
    view.SetFieldTextColor(BG_COLOR)
    view.SetTitle(title)
    view.SetTitleAlign(tview.AlignLeft)
    view.SetBackgroundColor(BG_COLOR)
    view.SetBorder(true)
    view.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))


    view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        currentOption, _ := view.GetCurrentOption()

        if event.Key() == tcell.KeyRune {
            switch event.Rune() {
            case 'j':
                nextOption := (currentOption + 1) % len(methods)
                view.SetCurrentOption(nextOption)
                return nil
            case 'k':
                prevOption := (currentOption - 1 + len(methods)) % len(methods)
                view.SetCurrentOption(prevOption)
                return nil
            }
        } else if event.Key() == tcell.KeyEnter {
            // set it here
            return nil
        }
        
        return event
    })

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
    for i, v := range methods {
        if v == method {
            return i
        }
    }
    return -1
}
