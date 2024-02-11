package ui

import (
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

var BG_COLOR = tcell.ColorDefault

type AppController interface {
	SendRequest(index int)
	CreateRequest()
	DeleteRequest(index int)
	GetRequests() ([]model.Request, error)
	UpdateRequest(model.Request)
	GetHosts() ([]model.Host, error)
	SetHosts()
	SelectRequest(index int)
	SetRequests()
	HandleRequestSelected(index int)
    AddHost()
    EditRequestBody(body string)
    AppSync()
}

var defaultHeaders = []RequestHeader{
	{Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
}

type AppView struct {
	layout         *tview.Flex
	statusBar      *StatusBar
	methodDropdown *MethodDropdown
	HostDropdown   *HostDropdown
	urlInput       *UrlInput
	headersTable   *HeadersTable
	RequestBody    *RequestBodyArea
	requestList    *RequestList
	responseBox    *ResponseView

	controller AppController
}

func NewAppView() *AppView {
	return &AppView{
		layout:         tview.NewFlex(),
		statusBar:      NewStatusBar(),
		methodDropdown: NewMethodDropdown(),
		HostDropdown:   NewHostDropdown(),
		urlInput:       NewUrlInput(),
		headersTable:   NewHeadersTable(),
		RequestBody:    NewRequestBodyArea(),
		requestList:    NewRequestList(),
		responseBox:    NewResponseView(),
	}
}

func (v *AppView) GetPrimitive() tview.Primitive {

    v.layout.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
            AddItem(v.requestList.GetPrimitive(), 50, 1, true).
                AddItem(tview.NewFlex().SetDirection(tview.FlexRow).

                AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
                    AddItem(v.methodDropdown.GetPrimitive(), 15, 1, false).
                    AddItem(v.HostDropdown.GetPrimitive(), 45, 1, false).
                    AddItem(v.urlInput.GetPrimitive(), 0, 1, false),
                    3, 1, false).

                AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
                    AddItem(v.RequestBody.GetPrimitive(), 0, 5, false).
                    AddItem(v.headersTable.GetPrimitive(), 0, 5, false),
                    0, 5, false).

                AddItem(v.responseBox.GetPrimitive(), 0, 5, false),
                0, 1, false),
            0, 2, true,
        ).

        AddItem(v.statusBar.GetPrimitive(), 1, 1, false), 
        0, 1, true,
    )

	v.controller.SetHosts()
	v.controller.SetRequests()

	return v.layout
}

func (v *AppView) GetFocusableComponents() []tview.Primitive {
	focusables := []tview.Primitive{
		v.requestList.GetPrimitive(),
		v.methodDropdown.GetPrimitive(),
		v.HostDropdown.GetPrimitive(),
		v.urlInput.GetPrimitive(),
		v.RequestBody.GetPrimitive(),
		v.headersTable.GetPrimitive(),
		v.responseBox.GetPrimitive(),
	}
	return focusables
}

func (v *AppView) Bind(controller AppController) {
	v.controller = controller
	v.methodDropdown.Bind(controller)
	v.urlInput.Bind(controller)
	v.requestList.Bind(controller)
    v.HostDropdown.Bind(controller)
    v.RequestBody.Bind(controller)
}

func (v *AppView) SetStatus(status string) {
	v.statusBar.SetStatus(status)
}

func (v *AppView) SetRequests(requests []model.Request) {
	v.requestList.SetContent(requests)
}

func (v *AppView) SetHosts(hosts []model.Host) {
	v.HostDropdown.SetHosts(hosts)
}

func (v *AppView) SetResponse(body string) {
	v.responseBox.SetContent(body)
}

func (v *AppView) SetSelectedRequest(index int) {
	v.requestList.SetSelectedRequest(index)
}

func (v *AppView) GetSelectedHost() string {
	return v.HostDropdown.GetSelectedHost()
}

func (v *AppView) RequestSelected(req model.Request) {
	v.methodDropdown.SetCurrentOption(req)
	v.urlInput.SetText(req)
	v.headersTable.DisplayHeaders(defaultHeaders)
	v.RequestBody.SetText(req.Body.String)
	v.responseBox.SetContent(req.LastResponse.String)
}

func (v *AppView) GetStatusBar() tview.Primitive {
    return v.statusBar.GetInputField()
}

func (v *AppView) OnStatusInputSubmit(cb func(value string)) {
    v.statusBar.OnInput(cb)
}
