package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatusBar struct {
	view *tview.Flex
    Input *tview.InputField
    label *tview.TextView
}

func NewStatusBar() *StatusBar {
    label := tview.NewTextView()
    label.SetBorderPadding(0, 0, 1, 0)
    input := tview.NewInputField()
    view := tview.NewFlex().SetDirection(tview.FlexColumn).
        AddItem(label, 0, 1, false)
    
    view.SetBackgroundColor(DEFAULT)
    label.SetBackgroundColor(DEFAULT)
	input.SetFieldBackgroundColor(DEFAULT)

    s := &StatusBar{view: view, label: label, Input: input}
    s.setInputCapture()
    return s
}

func (s *StatusBar) GetPrimitive() tview.Primitive {
	return s.view
}

func (s *StatusBar) GetInputField() tview.Primitive {
    return s.Input
}

func (s *StatusBar) SetStatus(msg string) {
	s.label.SetText(msg)
}

func (s *StatusBar) setInputCapture() {
    keybinds := func(event *tcell.EventKey) *tcell.EventKey {
        return event
    }

    s.Input.SetInputCapture(keybinds)
}

func (s *StatusBar) OnInput(callback func(value string)) {
    s.view.AddItem(s.Input, 0, 8, false)
    s.Input.SetDoneFunc(func (key tcell.Key) {
        if key == tcell.KeyEnter {
            callback(s.Input.GetText())        
            s.Input.SetText("")
            s.view.RemoveItem(s.Input)
        } else if key == tcell.KeyEscape {
            s.Input.SetText("")
            callback("")        
            s.view.RemoveItem(s.Input)
        }
    })
}
