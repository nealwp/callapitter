package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HostDropdown struct {
    view *tview.DropDown
}

var hosts = []string{"https://jsonplaceholder.typicode.com"}

func NewHostDropdown() *HostDropdown {

    title := "Host"

    view := tview.NewDropDown()
    view.SetOptions(hosts, nil)
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
                nextOption := (currentOption + 1) % len(hosts)
                view.SetCurrentOption(nextOption)
                return nil
            case 'k':
                prevOption := (currentOption - 1 + len(hosts)) % len(hosts)
                view.SetCurrentOption(prevOption)
                return nil
            }
        } else if event.Key() == tcell.KeyEnter {
            // set it here
            return nil
        }
        
        return event
    })

    return &HostDropdown{ view: view }
}

func (h *HostDropdown) GetPrimitive() tview.Primitive {
    return h.view
}

func (h *HostDropdown) GetSelectedHost() string {
    _, host := h.view.GetCurrentOption()
    return host
}
