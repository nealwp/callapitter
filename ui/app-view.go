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
}

var defaultHeaders = []RequestHeader{
	{Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
}

type AppView struct {
	layout         *tview.Flex
	statusBar      *StatusBar
	methodDropdown *MethodDropdown
	hostDropdown   *HostDropdown
	urlInput       *UrlInput
	headersTable   *HeadersTable
	requestBody    *RequestBodyArea
	requestList    *RequestList
	responseBox    *ResponseView

	controller AppController
}

func NewAppView() *AppView {
	return &AppView{
		layout:         tview.NewFlex(),
		statusBar:      NewStatusBar(),
		methodDropdown: NewMethodDropdown(),
		hostDropdown:   NewHostDropdown(),
		urlInput:       NewUrlInput(),
		headersTable:   NewHeadersTable(),
		requestBody:    NewRequestBodyArea(),
		requestList:    NewRequestList(),
		responseBox:    NewResponseView(),
	}
}

func (v *AppView) GetPrimitive() tview.Primitive {

	v.layout.AddItem(v.requestList.GetPrimitive(), 50, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(v.methodDropdown.GetPrimitive(), 15, 1, false).
				AddItem(v.hostDropdown.GetPrimitive(), 45, 1, false).
				AddItem(v.urlInput.GetPrimitive(), 0, 1, false),
				3, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(v.requestBody.GetPrimitive(), 0, 5, false).
				AddItem(v.headersTable.GetPrimitive(), 0, 5, false),
				0, 5, false).
			AddItem(v.responseBox.GetPrimitive(), 0, 5, false).
			AddItem(v.statusBar.GetPrimitive(), 3, 1, false),
			0, 2, false,
		)

	v.controller.SetHosts()
	v.controller.SetRequests()

	return v.layout
}

func (v *AppView) GetFocusableComponents() []tview.Primitive {
	focusables := []tview.Primitive{
		v.requestList.GetPrimitive(),
		v.methodDropdown.GetPrimitive(),
		v.hostDropdown.GetPrimitive(),
		v.urlInput.GetPrimitive(),
		v.requestBody.GetPrimitive(),
		v.headersTable.GetPrimitive(),
		v.responseBox.GetPrimitive(),
	}
	return focusables
}

func (v *AppView) SetController(controller AppController) {
	v.controller = controller
	v.methodDropdown.Bind(controller)
	v.urlInput.Bind(controller)
	v.requestList.Bind(controller)
}

func (v *AppView) SetStatus(status string) {
	v.statusBar.SetStatus(status)
}

func (v *AppView) SetRequests(requests []model.Request) {
	v.requestList.SetContent(requests)
}

func (v *AppView) SetHosts(hosts []model.Host) {
	v.hostDropdown.SetHosts(hosts)
}

func (v *AppView) SetResponse(body string) {
	v.responseBox.SetContent(body)
}

func (v *AppView) SetSelectedRequest(index int) {
	v.requestList.SetSelectedRequest(index)
}

func (v *AppView) GetSelectedHost() string {
	return v.hostDropdown.GetSelectedHost()
}

func (v *AppView) RequestSelected(req model.Request) {
	v.methodDropdown.SetCurrentOption(req)
	v.urlInput.SetText(req)
	v.headersTable.DisplayHeaders(defaultHeaders)
	v.requestBody.SetText(req.Body.String)
	v.responseBox.SetContent(req.LastResponse.String)
}
