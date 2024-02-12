package ui

import (
	"github.com/nealwp/callapitter/model"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

var BG_COLOR = tcell.ColorDefault

type AppController interface {
	SendRequest(index int)
	AddRequest()
    EditRequest(index int)
	DeleteRequest(index int)
	GetRequests() ([]model.Request, error)
	UpdateRequest(model.Request)
	GetHosts() ([]model.Host, error)
	SetHosts()
	SelectRequest(index int)
	SetRequests()
	HandleRequestSelected(index int)
    AddHost()
    EditRequestBody(req model.Request)
}

var defaultHeaders = []RequestHeader{
	{Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
}

type AppView struct {
	layout         *tview.Flex
	StatusBar      *StatusBar
	MethodDropdown *MethodDropdown
	HostDropdown   *HostDropdown
	urlInput       *UrlInput
	headersTable   *HeadersTable
	RequestBody    *RequestBodyArea
	RequestList    *RequestList
	responseBox    *ResponseView

	controller AppController
}

func NewAppView() *AppView {
	return &AppView{
		layout:         tview.NewFlex(),
		StatusBar:      NewStatusBar(),
		MethodDropdown: NewMethodDropdown(),
		HostDropdown:   NewHostDropdown(),
		urlInput:       NewUrlInput(),
		headersTable:   NewHeadersTable(),
		RequestBody:    NewRequestBodyArea(),
		RequestList:    NewRequestList(),
		responseBox:    NewResponseView(),
	}
}

func (v *AppView) GetPrimitive() tview.Primitive {

    v.layout.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
            AddItem(v.RequestList.GetPrimitive(), 50, 1, true).
                AddItem(tview.NewFlex().SetDirection(tview.FlexRow).

                AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
                    AddItem(v.MethodDropdown.GetPrimitive(), 15, 1, false).
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

        AddItem(v.StatusBar.GetPrimitive(), 1, 1, false), 
        0, 1, true,
    )

	v.controller.SetHosts()
	v.controller.SetRequests()

	return v.layout
}

func (v *AppView) Bind(controller AppController) {
	v.controller = controller
	v.MethodDropdown.Bind(controller)
	v.urlInput.Bind(controller)
	v.RequestList.Bind(controller)
    v.HostDropdown.Bind(controller)
    v.RequestBody.Bind(controller)
}

func (v *AppView) SetStatus(status string) {
	v.StatusBar.SetStatus(status)
}

func (v *AppView) SetRequests(requests []model.Request) {
	v.RequestList.SetContent(requests)
}

func (v *AppView) SetHosts(hosts []model.Host) {
	v.HostDropdown.SetHosts(hosts)
}

func (v *AppView) SetResponse(body string) {
	v.responseBox.SetContent(body)
}

func (v *AppView) SetSelectedRequest(index int) {
	v.RequestList.SetSelectedRequest(index)
}

func (v *AppView) GetSelectedHost() string {
	return v.HostDropdown.GetSelectedHost()
}

func (v *AppView) RequestSelected(req model.Request) {
	v.MethodDropdown.SetCurrentOption(req)
	v.urlInput.SetText(req)
	v.headersTable.DisplayHeaders(defaultHeaders)
	v.RequestBody.SetRequest(req)
	v.responseBox.SetContent(req.LastResponse.String)
}

func (v *AppView) GetStatusBar() tview.Primitive {
    return v.StatusBar.GetInputField()
}

func (v *AppView) OnStatusInputSubmit(cb func(value string)) {
    v.StatusBar.OnInput(cb)
}
