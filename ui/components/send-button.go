package ui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

type SendButton struct {
    view *tview.Button
}

func NewSendButton() *SendButton {
    view := tview.NewButton("Send [⏎]")
    view.SetBorder(true)
    view.SetStyle(tcell.StyleDefault.Background(BG_COLOR))
    view.SetLabelColorActivated(tcell.ColorDarkBlue)
    view.SetActivatedStyle(tcell.StyleDefault.Background(BG_COLOR))

    sendRequest := func() {
        // exampleText := "{\"hello\": \"this\", \"is\": \"my\", \"example\": \"json\", \"number\": 1}"
        // this should actually save the response to the req struct
        // hmmmm
        //resBox.SetContent(exampleText)
        //statusBar.SetStatus("req sent")
    }

    view.SetSelectedFunc(sendRequest)

    return &SendButton{view: view}
}

func (b *SendButton) GetPrimitive() tview.Primitive {
    return b.view
}


