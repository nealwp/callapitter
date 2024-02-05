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
    SelectRequest(index int)
}

var defaultHeaders = []ui.Header{
	{Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
}

type AppView struct {
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

    requests []model.Request
    hosts []model.Host

    controller AppController
}

func NewAppLayout() *AppView {

    l := &AppView{
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
		selected := l.requests[index]
		l.methodDropdown.SetCurrentOption(selected)
		l.urlInput.SetText(selected)
		l.headersTable.DisplayHeaders(defaultHeaders)
		l.reqBody.SetText(selected.Body.String)
		l.resBox.SetContent(selected.LastResponse.String)
	})

	return l
}

func (l *AppView) GetPrimitive() tview.Primitive {

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

    hosts, err := l.controller.GetHosts()
    if err != nil {
        l.statusBar.SetStatus(err.Error())
    }
    l.hosts = hosts
    l.hostDropdown.SetHosts(l.hosts)

    requests, err := l.controller.GetRequests()
    if err != nil {
        l.statusBar.SetStatus(err.Error())
    }

    l.requests = requests
    l.reqList.SetContent(l.requests)

	return l.view
}

func (l *AppView) GetFocusableComponents() []tview.Primitive {
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

func (l *AppView) SetController(controller AppController) {
    l.controller = controller
    l.methodDropdown.OnChange(controller)
    l.urlInput.OnChange(controller)
    l.reqList.SetHandler(controller)
}

func (v *AppView) SetStatus(status string) {
    v.statusBar.SetStatus(status)
}

func (l *AppView) SetRequests(requests []model.Request) {
    l.requests = requests
    l.reqList.SetContent(l.requests)
}

func (l *AppView) SetHosts(hosts []model.Host) {
    l.hosts = hosts
    l.hostDropdown.SetHosts(l.hosts)
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
