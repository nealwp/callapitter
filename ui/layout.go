package ui

import (
	"database/sql"
	"io"
	"net/http"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/nealwp/callapitter/model"
	"github.com/nealwp/callapitter/ui/components"
	"github.com/rivo/tview"
)

type AppController interface {
    SendRequest(req model.Request)
    CreateRequest(req model.Request) error
    DeleteRequest(req model.Request) error
    GetRequests() ([]model.Request, error)
    UpdateRequest(model.Request)
}

var defaultHeaders = []ui.Header{
	{Key: "Authorization", Value: "Bearer 12345ABCDEFG"},
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
    requests []model.Request

    controller AppController
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
		selected := l.requests[index]
		l.methodDropdown.SetCurrentOption(selected)
		l.urlInput.SetText(selected.Endpoint)
		l.headersTable.DisplayHeaders(defaultHeaders)
		l.reqBody.SetText(selected.Body.String)
		l.resBox.SetContent(selected.LastResponse.String)
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

    requests, err := l.controller.GetRequests()
    if err != nil {
        l.statusBar.SetStatus(err.Error())
    }

    l.requests = requests
    l.reqList.SetContent(l.requests)

    l.reqList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        index := l.reqList.GetSelectedRequest()
        if event.Key() == tcell.KeyRune {
            switch event.Rune() {
            case 'D':
                err := l.controller.DeleteRequest(l.requests[index])
                if err != nil {
                    l.statusBar.SetStatus(err.Error())
                }
            case 'j': 
                l.reqList.SetSelectedRequest(index+1)
            case 'k':
                l.reqList.SetSelectedRequest(index-1)
            case '%':
                newRequest := model.Request{Method: "GET", Endpoint: "/"}
                err := l.controller.CreateRequest(newRequest)
                if err != nil {
                    l.statusBar.SetStatus(err.Error())
                }
            }
            return event
        } else if event.Key() == tcell.KeyEnter {
            host := l.hostDropdown.GetSelectedHost()
            res, err := sendRequest(l.requests[index], host)    
            if err != nil {
                panic(err)
            }
            l.resBox.SetContent(res.Body)
            l.statusBar.SetStatus(res.Status)
            l.requests[index].LastResponse = sql.NullString{String: res.Body, Valid: true}

        }
        return event
    })

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

func (l *AppLayout) SetController(controller AppController) {
    l.controller = controller
    l.methodDropdown.OnChange(controller)
}

func (l *AppLayout) SetRequests(requests []model.Request) {
    l.requests = requests
    l.reqList.SetContent(l.requests)
}

type HttpResponse struct {
	Body   string
	Status string
}

func sendRequest(req model.Request, host string) (HttpResponse, error) {
    client := &http.Client{}

	url := host + req.Endpoint

    request, err := http.NewRequest(req.Method, url, strings.NewReader(req.Body.String))
    if err != nil {
        panic(err)
    }

	// do header things later

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return HttpResponse{}, err
	}

	body := string(bodyBytes)
	status := response.Status

	return HttpResponse{Body: body, Status: status}, nil
}
