package main

import (
	"bytes"
	"encoding/json"

	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/ui"
	"github.com/rivo/tview"
)

var BG_COLOR = tcell.ColorDefault

var app = tview.NewApplication()

var defaultHeaders = []ui.Header {
    {Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
}

var requests = []ui.HttpRequest {
    {Method: "GET", Endpoint: "/api/test/hello", Headers: defaultHeaders, Body: "", LastResponse: ""},
    {Method: "GET", Endpoint: "/api/test/health", Headers: defaultHeaders, Body: "", LastResponse: ""},
    {Method: "POST", Endpoint: "/api/test/user", Headers: defaultHeaders, Body: "{\"name\": \"foo\", \"age\": 99}", LastResponse: ""},
    {Method: "GET", Endpoint: "/api/test/users", Headers: defaultHeaders, Body: "", LastResponse: ""},
    {Method: "GET", Endpoint: "/api/test/some/really/long/address", Headers: defaultHeaders, Body: "", LastResponse: ""},
}

func prettyPrintJSON(inputJSON string) (string) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(inputJSON), "", "    "); err != nil {
        return inputJSON  
    }
    return prettyJSON.String()
}

func main() {

    statusBar := ui.NewStatusBar() 
    methodDropdown := ui.NewMethodDropdown()
    hostDropdown := ui.NewHostDropdown()
    urlInput := ui.NewUrlInput() 
    headersTable := ui.NewHeadersTable() 
    reqBody := ui.NewRequestBodyArea()

    reqList := ui.NewRequestList() 
    reqList.SetContent(requests)

    resBox := tview.NewTextView()
    resBox.SetTitle("Response")
    resBox.SetTitleAlign(tview.AlignLeft)
    resBox.SetBackgroundColor(BG_COLOR)
    resBox.SetBorder(true)

    //reqList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
    //    selected := requests[index]
    //    methodDropdown.SetCurrentOption(selected.Method)
    //    urlInput.SetText(selected.Endpoint)
    //    headersTable.DisplayHeaders(selected.Headers)
    //    reqBody.SetText(selected.Body)
    //    resBox.SetText(prettyPrintJSON(selected.LastResponse))
    //})

    sendRequest := func() {
        exampleText := "{\"hello\": \"this\", \"is\": \"my\", \"example\": \"json\", \"number\": 1}"
        // this should actually save the response to the req struct
        resBox.SetText(prettyPrintJSON(exampleText))
        statusBar.SetStatus("req sent")
    }

    sendBtn := tview.NewButton("Send [⏎]")
    sendBtn.SetBorder(true)
    sendBtn.SetStyle(tcell.StyleDefault.Background(BG_COLOR))
    sendBtn.SetLabelColorActivated(tcell.ColorDarkBlue)
    sendBtn.SetActivatedStyle(tcell.StyleDefault.Background(BG_COLOR))

    sendBtn.SetSelectedFunc(sendRequest)

    layout := tview.NewFlex().
    AddItem(reqList.GetPrimitive(), 50, 1, true).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
            AddItem(methodDropdown.GetPrimitive(), 15, 1, false).
            AddItem(hostDropdown.GetPrimitive(), 45, 1, false).
            AddItem(urlInput.GetPrimitive(), 0, 1, false).
            AddItem(sendBtn, 12, 1, false),
            3, 1, false).
        AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
            AddItem(reqBody.GetPrimitive(), 0, 5, false).
            AddItem(headersTable.GetPrimitive(), 0, 5, false),
            0, 5, false).
        AddItem(resBox, 0, 5, false).
        AddItem(statusBar.GetPrimitive(), 3, 1, false), 
        0, 2, false,
)


    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        case tcell.KeyCtrlC:
            return nil
        case tcell.KeyEnter:
            sendRequest()
            return nil
        case tcell.KeyCtrlR:
            app.SetFocus(reqList.GetPrimitive()) 
        case tcell.KeyCtrlH:
            app.SetFocus(hostDropdown.GetPrimitive())
        case tcell.KeyCtrlB:
            app.SetFocus(reqBody.GetPrimitive())
        case tcell.KeyCtrlD:
            app.SetFocus(headersTable.GetPrimitive())
        case tcell.KeyCtrlU:
            app.SetFocus(urlInput.GetPrimitive())
        case tcell.KeyCtrlE:
            app.SetFocus(methodDropdown.GetPrimitive())
        case tcell.KeyRune:
            if event.Rune() == 'q' && !urlInput.GetPrimitive().HasFocus() {
                app.Stop()
            }
        }
        return event
    })

    app.EnableMouse(true)
    app.SetRoot(layout, true)

    if err := app.Run(); err != nil {
        panic(err)
    }
}
