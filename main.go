package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/ui"
	"github.com/rivo/tview"
)

var BG_COLOR = tcell.ColorDefault

var app = tview.NewApplication()

func main() {

    layout := ui.NewAppLayout()

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        case tcell.KeyCtrlC:
            return nil
        case tcell.KeyEnter:
            //sendRequest()
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
