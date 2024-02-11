package controller

import (
	"io"
	"net/http"
	"strings"

	"github.com/nealwp/callapitter/model"
	"github.com/nealwp/callapitter/ui"
	"github.com/rivo/tview"
)

type AppController struct {
	model *model.AppModel
	view  *ui.AppView
    app *tview.Application
}

func NewAppController() *AppController {
	return &AppController{}
}

func (c *AppController) Bind(app *tview.Application, model *model.AppModel, view *ui.AppView) {
	c.view = view
    c.model = model
    c.app = app
}

func (c *AppController) AppSync() {
    c.app.Sync()
}

type HttpResponse struct {
	Body   string
	Status string
}

func sendHttpRequest(req model.Request, host string) (HttpResponse, error) {
	client := &http.Client{}

	url := host + req.Endpoint

	request, err := http.NewRequest(req.Method, url, strings.NewReader(req.Body.String))
	if err != nil {
		return HttpResponse{}, err
	}

	// do header things later

	response, err := client.Do(request)
	if err != nil {
		return HttpResponse{}, err
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
