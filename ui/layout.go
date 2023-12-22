package ui

import (
    "github.com/rivo/tview"
)

type AppLayout struct {
    view *tview.Flex
}

var defaultHeaders = []Header {
    {Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
}

var requests = []HttpRequest {
    {Method: "GET", Endpoint: "/api/test/hello", Headers: defaultHeaders, Body: "", LastResponse: ""},
    {Method: "GET", Endpoint: "/api/test/health", Headers: defaultHeaders, Body: "", LastResponse: ""},
    {Method: "POST", Endpoint: "/api/test/user", Headers: defaultHeaders, Body: "{\"name\": \"foo\", \"age\": 99}", LastResponse: ""},
    {Method: "GET", Endpoint: "/api/test/users", Headers: defaultHeaders, Body: "", LastResponse: ""},
    {Method: "GET", Endpoint: "/api/test/some/really/long/address", Headers: defaultHeaders, Body: "", LastResponse: ""},
}

func NewAppLayout() *AppLayout {

    statusBar := NewStatusBar() 
    methodDropdown := NewMethodDropdown()
    hostDropdown := NewHostDropdown()
    urlInput := NewUrlInput() 
    headersTable := NewHeadersTable() 
    reqBody := NewRequestBodyArea()

    reqList := NewRequestList() 
    reqList.SetContent(requests)

    resBox := NewResponseView() 
    sendBtn := NewSendButton() 

    view := tview.NewFlex().
    AddItem(reqList.GetPrimitive(), 50, 1, true).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
            AddItem(methodDropdown.GetPrimitive(), 15, 1, false).
            AddItem(hostDropdown.GetPrimitive(), 45, 1, false).
            AddItem(urlInput.GetPrimitive(), 0, 1, false).
            AddItem(sendBtn.GetPrimitive(), 12, 1, false),
            3, 1, false).
        AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
            AddItem(reqBody.GetPrimitive(), 0, 5, false).
            AddItem(headersTable.GetPrimitive(), 0, 5, false),
            0, 5, false).
        AddItem(resBox.GetPrimitive(), 0, 5, false).
        AddItem(statusBar.GetPrimitive(), 3, 1, false), 
        0, 2, false,
    )

    return &AppLayout{view: view}
}

func (l *AppLayout) GetPrimitive() tview.Primitive {
    return l.view
}
