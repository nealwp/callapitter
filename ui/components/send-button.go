package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: what if we had no button instead...
type SendButton struct {
	view *tview.Button
}

func NewSendButton() *SendButton {
	label := "Send"

	view := tview.NewButton(label)
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
