package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/ui"
	"github.com/rivo/tview"
)

var BG_COLOR = tcell.ColorDefault

var app = tview.NewApplication()


type Header struct {
    Key string
    Value string
}

type HttpRequest struct {
    Method string
    Endpoint string
    Headers []Header
    Body string
    LastResponse string
}

var defaultHeaders = []Header {
    {"Authorization", "Bearer 12345ABCDEFG"},
}

var requests = []HttpRequest {
    {"GET", "/api/test/hello", defaultHeaders, "", ""},
    {"GET", "/api/test/health", defaultHeaders, "", ""},
    {"POST", "/api/test/user", defaultHeaders, "{\"name\": \"foo\", \"age\": 99}", ""},
    {"GET", "/api/test/users", defaultHeaders, "", ""},
    {"GET", "/api/test/some/really/long/address", defaultHeaders, "", ""},
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

    headersTable := tview.NewFlex()
    headersTable.SetBackgroundColor(BG_COLOR)

    displayHeaders := func(headers []Header) {
        headersTable.Clear()
        for _, h := range(headers) {
            key := tview.NewInputField().SetFieldBackgroundColor(BG_COLOR)
            key.SetBackgroundColor(BG_COLOR)
            key.SetText(h.Key)
            value := tview.NewInputField().SetFieldBackgroundColor(BG_COLOR)
            value.SetBackgroundColor(BG_COLOR)
            value.SetText(h.Value)
            row := tview.NewFlex().SetDirection(tview.FlexColumn)
            row.SetBackgroundColor(BG_COLOR)
            row.AddItem(key, 20, 1, false)
            row.AddItem(value, 20, 1, false)
            headersTable.AddItem(row, 0, 1, false)
        }
    }

    reqBody := tview.NewTextArea()
    reqBody.SetBackgroundColor(BG_COLOR)
    reqBody.SetTextStyle(tcell.StyleDefault.Background(BG_COLOR))

    pages := tview.NewPages()
    pages.SetBackgroundColor(BG_COLOR)
    pages.SetBorder(true)

    pages.AddPage("headers", headersTable, true, false)
    pages.AddPage("body", reqBody, true, false)

    headersTab := tview.NewButton("Headers [C-d]")
    headersTab.SetStyle(tcell.StyleDefault.Background(BG_COLOR))
    headersTab.SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorGray))

    selectHeadersTab := func() {
        pages.SwitchToPage("headers")
    }

    headersTab.SetSelectedFunc(selectHeadersTab)

    bodyTab := tview.NewButton("Body [C-b]")
    bodyTab.SetStyle(tcell.StyleDefault.Background(BG_COLOR))
    bodyTab.SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorGray))

    selectBodyTab := func() {
        pages.SwitchToPage("body")
    }

    bodyTab.SetSelectedFunc(selectBodyTab)

    tabs := tview.NewFlex().AddItem(headersTab, 14, 1, false).AddItem(bodyTab, 14, 1, false)

    pages.SwitchToPage("headers")

    reqList := tview.NewList()

    for _, r := range requests {
        reqList.AddItem(fmt.Sprintf("%-4s", r.Method) + "  " + r.Endpoint, "", 0, nil)
    }


    reqList.ShowSecondaryText(false)
    reqList.SetBorder(true)
    reqList.SetBackgroundColor(BG_COLOR)
    reqList.SetTitle("Requests [C-r]")
    reqList.SetTitleAlign(tview.AlignLeft)
    reqList.SetBorderPadding(1,1,1,1)

    reqListInputCapture := func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyRune {
            idx := reqList.GetCurrentItem()
            switch event.Rune() {
            case 'j':
                reqList.SetCurrentItem(idx+1)
            case 'k':
                reqList.SetCurrentItem(idx-1)
            }
            return event
        }
        return event
    }

    reqList.SetInputCapture(reqListInputCapture)

    resBox := tview.NewTextView()
    resBox.SetTitle("Response")
    resBox.SetTitleAlign(tview.AlignLeft)
    resBox.SetBackgroundColor(BG_COLOR)
    resBox.SetBorder(true)

    reqList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
        selected := requests[index]
        methodDropdown.SetCurrentOption(selected.Method)
        urlInput.SetText(selected.Endpoint)
        displayHeaders(selected.Headers)
        reqBody.SetText(prettyPrintJSON(selected.Body), false)
        resBox.SetText(prettyPrintJSON(selected.LastResponse))
    })

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
    AddItem(reqList, 50, 1, true).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
    AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
    AddItem(methodDropdown.GetPrimitive(), 15, 1, false).
    AddItem(hostDropdown.GetPrimitive(), 45, 1, false).
    AddItem(urlInput.GetPrimitive(), 0, 1, false).
    AddItem(sendBtn, 12, 1, false),
    3, 1, false).
    AddItem(tabs, 1, 0, false).
    AddItem(pages, 0, 5, false).
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
            app.SetFocus(reqList) 
        case tcell.KeyCtrlH:
            app.SetFocus(hostDropdown.GetPrimitive())
        case tcell.KeyCtrlB:
            selectBodyTab()
        case tcell.KeyCtrlD:
            selectHeadersTab()
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
