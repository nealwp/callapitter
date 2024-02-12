package main

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/controller"
	"github.com/nealwp/callapitter/model"
	"github.com/nealwp/callapitter/ui"
	"github.com/rivo/tview"
)

func main() {

    dbPath, err := getDatabasePath()
    if err != nil {
        panic(err)
    }

	model := model.NewAppModel(dbPath)
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
            case 'p':
                app.SetFocus(view.ResponseBox.GetPrimitive())
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

func getDatabasePath() (string, error) {
    home := os.Getenv("HOME")
    xdgLocalShare := filepath.Join(home, ".local", "share")

    appDir := filepath.Join(xdgLocalShare, "callapitter")

    if _, err := os.Stat(appDir); os.IsNotExist(err) {
        if err := os.MkdirAll(appDir, 0755); err != nil {
            return "", err
        }
    }

    dbPath := filepath.Join(appDir, "callapitter.db")
    return dbPath, nil
}

