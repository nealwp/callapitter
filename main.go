package main

import (
	"fmt"
    "bytes"
    "encoding/json"

	"github.com/gdamore/tcell/v2"
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
}

var defaultHeaders = []Header {
    {"Authorization", "Bearer 12345ABCDEFG"},
}

var requests = []HttpRequest {
    {"GET", "/api/test/hello", defaultHeaders, ""},
    {"GET", "/api/test/health", defaultHeaders, ""},
    {"POST", "/api/test/user", defaultHeaders, "{\"name\": \"foo\", \"age\": 99}"},
    {"GET", "/api/test/users", defaultHeaders, ""},
    {"GET", "/api/test/some/really/long/address", defaultHeaders, ""},
}

func prettyPrintJSON(inputJSON string) (string) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(inputJSON), "", "    "); err != nil {
        return inputJSON  
    }
    return prettyJSON.String()
}

func main() {


    statusBar := tview.NewTextView()
    statusBar.SetBorder(true)
    statusBar.SetBackgroundColor(BG_COLOR)
    statusBar.SetTitle("Status")
    statusBar.SetTitleAlign(tview.AlignLeft)

    setStatus := func(msg string) {
        statusBar.SetText(msg)
    }

    reqMethods := []string{"GET", "POST", "PUT", "DELETE"}

    findMethodIndex := func(method string) int {
        for i, v := range reqMethods {
            if v == method {
                return i
            }
        }
        return -1
    }

    methodDropdown := tview.NewDropDown().SetOptions(reqMethods, nil).SetCurrentOption(0).SetFieldBackgroundColor(BG_COLOR).SetFieldTextColor(BG_COLOR)
    methodDropdown.SetTitle("Method [C-e]")
    methodDropdown.SetTitleAlign(tview.AlignLeft)
    methodDropdown.SetBackgroundColor(BG_COLOR)
    methodDropdown.SetBorder(true)
    methodDropdown.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

    hosts := []string{"http://localhost:8000"}

    hostDropdown := tview.NewDropDown().SetOptions(hosts, nil).SetCurrentOption(0).SetFieldBackgroundColor(BG_COLOR).SetFieldTextColor(BG_COLOR)
    hostDropdown.SetTitle("Host [C-h]")
    hostDropdown.SetTitleAlign(tview.AlignLeft)
    hostDropdown.SetBackgroundColor(BG_COLOR)
    hostDropdown.SetBorder(true)
    hostDropdown.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

    urlInput := tview.NewInputField().SetFieldBackgroundColor(BG_COLOR).SetFieldTextColor(BG_COLOR)
    urlInput.SetBackgroundColor(BG_COLOR)
    urlInput.SetBorder(true)
    urlInput.SetTitle("URL [C-u]")
    urlInput.SetTitleAlign(tview.AlignLeft)

    urlInputCapture := func(event *tcell.EventKey) *tcell.EventKey {
        return event
    }

    urlInput.SetInputCapture(urlInputCapture)

    sendBtn := tview.NewButton("Send [⏎]")
    sendBtn.SetBorder(true)
    sendBtn.SetStyle(tcell.StyleDefault.Background(BG_COLOR))
    sendBtn.SetLabelColorActivated(tcell.ColorDarkBlue)
    sendBtn.SetActivatedStyle(tcell.StyleDefault.Background(BG_COLOR))

    sendRequest := func() {
       setStatus("req sent")
    }

    sendBtn.SetSelectedFunc(sendRequest)

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

    reqList := tview.NewList()

    for _, r := range requests {
        reqList.AddItem(fmt.Sprintf("%-4s", r.Method) + "  " + r.Endpoint, "", 0, nil)
    }

    reqList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
        selected := requests[index]

        methodIdx := findMethodIndex(selected.Method)
        methodDropdown.SetCurrentOption(methodIdx)
        urlInput.SetText(selected.Endpoint)
        body := prettyPrintJSON(selected.Body)
        displayHeaders(selected.Headers)
        reqBody.SetText(body, false)
    })

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

    resBox := tview.NewBox()
    resBox.SetTitle("Response")
    resBox.SetTitleAlign(tview.AlignLeft)
    resBox.SetBackgroundColor(BG_COLOR)
    resBox.SetBorder(true)

    layout := tview.NewFlex().
        AddItem(reqList, 50, 1, true).
        AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
            AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
                AddItem(methodDropdown, 15, 1, false).
                AddItem(hostDropdown, 25, 1, false).
                AddItem(urlInput, 0, 1, false).
                AddItem(sendBtn, 12, 1, false),
            3, 1, false).
            AddItem(tabs, 1, 0, false).
            AddItem(pages, 0, 5, false).
            AddItem(resBox, 0, 5, false).
            AddItem(statusBar, 3, 1, false), 
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
            app.SetFocus(hostDropdown)
        case tcell.KeyCtrlB:
            selectBodyTab()
        case tcell.KeyCtrlD:
            selectHeadersTab()
        case tcell.KeyCtrlU:
            app.SetFocus(urlInput)
        case tcell.KeyCtrlE:
            app.SetFocus(methodDropdown)
        case tcell.KeyRune:
            if event.Rune() == 'q' && !urlInput.HasFocus() {
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
