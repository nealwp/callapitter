package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
)

type HostDropdown struct {
	view  *tview.DropDown
	hosts []model.Host
    handler AppController
}

func NewHostDropdown() *HostDropdown {

	title := "Host (h)"

	view := tview.NewDropDown()
	view.SetFieldBackgroundColor(DEFAULT)
	view.SetFieldTextColor(DEFAULT)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)
	view.SetBorder(true)
	view.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

	h := &HostDropdown{view: view}
	h.setKeyBindings()
	return h
}

func (h *HostDropdown) GetPrimitive() tview.Primitive {
	return h.view
}

func (h *HostDropdown) Bind(handler AppController) {
    h.handler = handler
}

func (h *HostDropdown) SetHosts(hosts []model.Host) {

	h.hosts = hosts

	var hostnames []string

	for _, h := range hosts {
		hostnames = append(hostnames, h.Name)
	}

	h.view.SetOptions(hostnames, nil)
	h.view.SetCurrentOption(0)
}

func (h *HostDropdown) GetSelectedHost() string {
	_, host := h.view.GetCurrentOption()
	return host
}

func (h *HostDropdown) setKeyBindings() {

	keybinds := func(event *tcell.EventKey) *tcell.EventKey {
		index, _ := h.view.GetCurrentOption()

		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case 'D':
				h.handler.DeleteHost(h.hosts[index])
				return nil
			case 'j':
				nextOption := (index + 1) % len(h.hosts)
				h.view.SetCurrentOption(nextOption)
				return nil
			case 'k':
				prevOption := (index - 1 + len(h.hosts)) % len(h.hosts)
				h.view.SetCurrentOption(prevOption)
				return nil
            case '%':
                h.handler.AddHost()
                return nil
			}
		} else if event.Key() == tcell.KeyEnter {
			// set it here
			return nil
		}

		return event
	}

	h.view.SetInputCapture(keybinds)
}
