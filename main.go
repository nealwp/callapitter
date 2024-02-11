package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/controller"
	"github.com/nealwp/callapitter/model"
	"github.com/nealwp/callapitter/ui"
	"github.com/rivo/tview"
)

func main() {

	model := model.NewAppModel("./callapitter.db")
    view := ui.NewAppView()
	controller := controller.NewAppController()

	components := view.GetFocusableComponents()

    app := tview.NewApplication()

    controller.Bind(app, model, view)
    view.Bind(controller)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			focusNext(app, components)
			return nil
		case tcell.KeyBacktab:
			focusPrevious(app, components)
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

	app.SetRoot(view.GetPrimitive(), true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func focusNext(app *tview.Application, components []tview.Primitive) {

	currentFocus := app.GetFocus()
	for i, component := range components {
		if component == currentFocus {
			nextIndex := (i + 1) % len(components)
			app.SetFocus(components[nextIndex])
			break
		}
	}
}

func focusPrevious(app *tview.Application, components []tview.Primitive) {

	currentFocus := app.GetFocus()
	for i, component := range components {
		if component == currentFocus {
			prevIndex := (i - 1 + len(components)) % len(components)
			app.SetFocus(components[prevIndex])
			break
		}
	}
}
