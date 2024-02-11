package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatusBar struct {
	view *tview.Flex
    input *tview.InputField
    label *tview.TextView
}

func NewStatusBar() *StatusBar {
    label := tview.NewTextView()
    label.SetBorderPadding(0, 0, 1, 0)
    input := tview.NewInputField()
    view := tview.NewFlex().SetDirection(tview.FlexColumn).
        AddItem(label, 0, 1, false)
    
    view.SetBackgroundColor(BG_COLOR)
    label.SetBackgroundColor(BG_COLOR)
	input.SetFieldBackgroundColor(BG_COLOR)

    return &StatusBar{view: view, label: label, input: input}
}

func (s *StatusBar) GetPrimitive() tview.Primitive {
	return s.view
}

func (s *StatusBar) GetInputField() tview.Primitive {
    return s.input
}

func (s *StatusBar) SetStatus(msg string) {
	s.label.SetText(msg)
}

func (s *StatusBar) OnInput(callback func(value string)) {
    s.view.AddItem(s.input, 0, 8, false)
    s.input.SetDoneFunc(func (key tcell.Key) {
        if key == tcell.KeyEnter {
            callback(s.input.GetText())        
            s.input.SetText("")
            s.view.RemoveItem(s.input)
        } else if key == tcell.KeyEscape {
            s.input.SetText("")
            callback("")        
            s.view.RemoveItem(s.input)
        }
    })
}
