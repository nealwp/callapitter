package ui

import (
	"github.com/nealwp/callapitter/ui/components"
	"github.com/rivo/tview"
)


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

type AppLayout struct {
    view *tview.Flex
    statusBar *ui.StatusBar
    methodDropdown *ui.MethodDropdown
    hostDropdown *ui.HostDropdown
    urlInput *ui.UrlInput
    headersTable *ui.HeadersTable
    reqBody *ui.RequestBodyArea
    reqList *ui.RequestList
    resBox *ui.ResponseView
    sendBtn *ui.SendButton
}

func NewAppLayout() *AppLayout {

    l := &AppLayout{
        view: tview.NewFlex(),
        statusBar: ui.NewStatusBar(),
        methodDropdown: ui.NewMethodDropdown(),
        hostDropdown: ui.NewHostDropdown(),
        urlInput: ui.NewUrlInput(),
        headersTable: ui.NewHeadersTable(),
        reqBody: ui.NewRequestBodyArea(),
        reqList: ui.NewRequestList(),
        resBox: ui.NewResponseView(),
        sendBtn: ui.NewSendButton(),
    }
    
    l.reqList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
        selected := requests[index]
        l.methodDropdown.SetCurrentOption(selected.Method)
        l.urlInput.SetText(selected.Endpoint)
        l.headersTable.DisplayHeaders(selected.Headers)
        l.reqBody.SetText(selected.Body)
        l.resBox.SetContent(selected.LastResponse)
    })

    return l
}

func (l *AppLayout) GetPrimitive() tview.Primitive {

    l.view.AddItem(l.reqList.GetPrimitive(), 50, 1, true).
        AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
            AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
                AddItem(l.methodDropdown.GetPrimitive(), 15, 1, false).
                AddItem(l.hostDropdown.GetPrimitive(), 45, 1, false).
                AddItem(l.urlInput.GetPrimitive(), 0, 1, false).
                AddItem(l.sendBtn.GetPrimitive(), 12, 1, false),
                3, 1, false).
            AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
                AddItem(l.reqBody.GetPrimitive(), 0, 5, false).
                AddItem(l.headersTable.GetPrimitive(), 0, 5, false),
                0, 5, false).
            AddItem(l.resBox.GetPrimitive(), 0, 5, false).
            AddItem(l.statusBar.GetPrimitive(), 3, 1, false), 
            0, 2, false,
        )

    // these will come from db
    l.reqList.SetContent(requests)

    return l.view
}

func (l *AppLayout) GetFocusableComponents() []tview.Primitive {
    focusables := []tview.Primitive{
        l.reqList.GetPrimitive(), 
        l.methodDropdown.GetPrimitive(), 
        l.hostDropdown.GetPrimitive(),
        l.urlInput.GetPrimitive(),
        l.reqBody.GetPrimitive(),
        l.headersTable.GetPrimitive(),
        l.resBox.GetPrimitive(),
    }
    return focusables
}
