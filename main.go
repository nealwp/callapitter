package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/ui"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()
	layout := ui.NewAppLayout()

	focusables := layout.GetFocusableComponents()

	focusNext := func() {
		currentFocus := app.GetFocus()
		for i, component := range focusables {
			if component == currentFocus {
				nextIndex := (i + 1) % len(focusables)
				app.SetFocus(focusables[nextIndex])
				break
			}
		}
	}

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			focusNext()
			return nil
		case tcell.KeyCtrlC:
			return nil
		case tcell.KeyRune:
			if event.Rune() == 'q' {
				app.Stop()
			}
		}
		return event
	})

	app.EnableMouse(true)
	app.SetRoot(layout.GetPrimitive(), true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
