package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
)

type MethodDropdown struct {
	view    *tview.DropDown
	handler AppController
	request model.Request
}

var methods = []string{"GET", "POST", "PUT", "DELETE"}

func NewMethodDropdown() *MethodDropdown {

	title := "Method (m)"

	view := tview.NewDropDown()
	view.SetOptions(methods, nil)
	view.SetCurrentOption(0)
	view.SetFieldBackgroundColor(BG_COLOR)
    view.SetFieldTextColor(tcell.ColorOrange)
	view.SetTitle(title)
	view.SetTitleAlign(tview.AlignLeft)
	view.SetBorder(true)
	view.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

	m := &MethodDropdown{view: view}
	m.setKeyBindings()
	return m
}

func (m *MethodDropdown) GetPrimitive() tview.Primitive {
	return m.view
}

func (m *MethodDropdown) Bind(handler AppController) {
	m.handler = handler
}

func (m *MethodDropdown) SetCurrentOption(req model.Request) {
	m.request = req
	index := findMethodIndex(req.Method)
	m.view.SetCurrentOption(index)
}

func (m *MethodDropdown) setKeyBindings() {

	keybinds := func(event *tcell.EventKey) *tcell.EventKey {
		index, _ := m.view.GetCurrentOption()

		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case 'j':
				nextOption := (index + 1) % len(methods)
				m.view.SetCurrentOption(nextOption)
				m.setRequestMethod()
				return nil
			case 'k':
				prevOption := (index - 1 + len(methods)) % len(methods)
				m.view.SetCurrentOption(prevOption)
				m.setRequestMethod()
				return nil
			}
		}

		return event
	}

	m.view.SetInputCapture(keybinds)
}

func (m *MethodDropdown) setRequestMethod() {
	index, _ := m.view.GetCurrentOption()
	m.request.Method = methods[index]
	m.handler.UpdateRequest(m.request)
}

func findMethodIndex(method string) int {
	for i, v := range methods {
		if v == method {
			return i
		}
	}
	return -1
}
