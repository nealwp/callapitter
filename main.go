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


    app := tview.NewApplication()

    controller.Bind(app, model, view)
    view.Bind(controller)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlC:
			return nil
		case tcell.KeyRune:
            if app.GetFocus() == view.StatusBar.Input {
                return event 
            }
            switch event.Rune() {
            case 'b':
                app.SetFocus(view.RequestBody.GetPrimitive())
                return nil
            case 'h':
                app.SetFocus(view.HostDropdown.GetPrimitive())
                return nil
            case 'm':
                app.SetFocus(view.MethodDropdown.GetPrimitive())
                return nil
            case 'q':
                app.Stop()
            case 'r':
                app.SetFocus(view.RequestList.GetPrimitive())
                return nil
            }
		}
		return event
	})

	app.SetRoot(view.GetPrimitive(), true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

