package ui

import (
	"github.com/nealwp/callapitter/model"
	"github.com/nealwp/callapitter/ui/components"
	"github.com/rivo/tview"
)

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

var defaultHeaders = []ui.Header{
	{Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
}

type AppView struct {
	view           *tview.Flex
	statusBar      *ui.StatusBar
	methodDropdown *ui.MethodDropdown
	hostDropdown   *ui.HostDropdown
	urlInput       *ui.UrlInput
	headersTable   *ui.HeadersTable
	reqBody        *ui.RequestBodyArea
	reqList        *ui.RequestList
	resBox         *ui.ResponseView
	sendBtn        *ui.SendButton

	controller AppController
}

func NewAppView() *AppView {
	return &AppView{
		view:           tview.NewFlex(),
		statusBar:      ui.NewStatusBar(),
		methodDropdown: ui.NewMethodDropdown(),
		hostDropdown:   ui.NewHostDropdown(),
		urlInput:       ui.NewUrlInput(),
		headersTable:   ui.NewHeadersTable(),
		reqBody:        ui.NewRequestBodyArea(),
		reqList:        ui.NewRequestList(),
		resBox:         ui.NewResponseView(),
		sendBtn:        ui.NewSendButton(),
	}
}

func (v *AppView) GetPrimitive() tview.Primitive {

	v.view.AddItem(v.reqList.GetPrimitive(), 50, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(v.methodDropdown.GetPrimitive(), 15, 1, false).
				AddItem(v.hostDropdown.GetPrimitive(), 45, 1, false).
				AddItem(v.urlInput.GetPrimitive(), 0, 1, false).
				AddItem(v.sendBtn.GetPrimitive(), 12, 1, false),
				3, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(v.reqBody.GetPrimitive(), 0, 5, false).
				AddItem(v.headersTable.GetPrimitive(), 0, 5, false),
				0, 5, false).
			AddItem(v.resBox.GetPrimitive(), 0, 5, false).
			AddItem(v.statusBar.GetPrimitive(), 3, 1, false),
			0, 2, false,
		)

	v.controller.SetHosts()
	v.controller.SetRequests()

	return v.view
}

func (v *AppView) GetFocusableComponents() []tview.Primitive {
	focusables := []tview.Primitive{
		v.reqList.GetPrimitive(),
		v.methodDropdown.GetPrimitive(),
		v.hostDropdown.GetPrimitive(),
		v.urlInput.GetPrimitive(),
		v.reqBody.GetPrimitive(),
		v.headersTable.GetPrimitive(),
		v.resBox.GetPrimitive(),
	}
	return focusables
}

func (v *AppView) SetController(controller AppController) {
	v.controller = controller
	v.methodDropdown.OnChange(controller)
	v.urlInput.OnChange(controller)
	v.reqList.SetHandler(controller)
}

func (v *AppView) SetStatus(status string) {
	v.statusBar.SetStatus(status)
}

func (v *AppView) SetRequests(requests []model.Request) {
	v.reqList.SetContent(requests)
}

func (v *AppView) SetHosts(hosts []model.Host) {
	v.hostDropdown.SetHosts(hosts)
}

func (v *AppView) SetResponse(body string) {
	v.resBox.SetContent(body)
}

func (v *AppView) SetSelectedRequest(index int) {
	v.reqList.SetSelectedRequest(index)
}

func (v *AppView) GetSelectedHost() string {
	return v.hostDropdown.GetSelectedHost()
}

func (v *AppView) RequestSelected(req model.Request) {
	v.methodDropdown.SetCurrentOption(req)
	v.urlInput.SetText(req)
	v.headersTable.DisplayHeaders(defaultHeaders)
	v.reqBody.SetText(req.Body.String)
	v.resBox.SetContent(req.LastResponse.String)
}
