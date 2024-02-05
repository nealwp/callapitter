package controller

import (
    "net/http"
    "io"
    "strings"

	"github.com/nealwp/callapitter/model"
	"github.com/nealwp/callapitter/ui"
)

type AppController struct {
    model *model.AppModel
    view *ui.AppView
}

func NewAppController() *AppController {
    return &AppController{}
}

func (c *AppController) SetView(view *ui.AppView) {
    c.view = view
}

func (c *AppController) SetModel(model *model.AppModel) {
    c.model = model
}

func (c *AppController) SendRequest(i int) {

    host := c.view.GetSelectedHost()
    req := c.model.Request.GetRequest(i)

    res, err := sendRequest(req, host)    

    if err != nil {
        c.view.SetStatus(err.Error())  
        return
    }

    c.view.SetResponse(res.Body)
    c.view.SetStatus(res.Status)

    //l.requests[index].LastResponse = sql.NullString{String: res.Body, Valid: true}
}

func (c *AppController) CreateRequest() {
    
    req := model.Request{Method: "GET", Endpoint: "/"}

    err := c.model.Request.InsertRequest(req) 
    if err != nil {
        c.view.SetStatus(err.Error())
        return 
    }

    requests, err := c.GetRequests()
    if err != nil {
        c.view.SetStatus(err.Error())
        return
    }

    c.view.SetRequests(requests)
}

func (c *AppController) SetRequests() {
    requests, err := c.GetRequests()

    if err != nil {
        c.view.SetStatus(err.Error())
    }

    c.view.SetRequests(requests)
}

func (c *AppController) HandleRequestSelected(index int) {
    req := c.model.Request.GetRequest(index)
    c.view.RequestSelected(req)
}

func (c *AppController) SetHosts() {
    hosts, err := c.GetHosts()

    if err != nil {
        c.view.SetStatus(err.Error())
    }

    c.view.SetHosts(hosts)
}

func (c *AppController) DeleteRequest(index int) {

    err := c.model.Request.DeleteRequest(index)

    if err != nil {
        c.view.SetStatus(err.Error())     
        return
    }

    requests, err := c.GetRequests()

    if err != nil {
        c.view.SetStatus(err.Error())
        return
    }

    c.view.SetRequests(requests)
}

func (c *AppController) SelectRequest(index int) {
   c.view.SetSelectedRequest(index)
}

func (c *AppController) GetRequests() ([]model.Request, error) {
    requests, err := c.model.Request.GetRequests()
    if err != nil {
        return nil, err
    }
    return requests, nil
}

func (c *AppController) UpdateRequest(req model.Request) {
    err := c.model.Request.UpdateRequest(req)

    if err != nil {
        panic(err)
    }

    requests, err := c.GetRequests()

    if err != nil {
        panic(err)
    }

    c.view.SetRequests(requests)
}

func (c *AppController) CreateHost(host model.Host) error {

    err := c.model.Host.InsertHost(host)
    if err != nil {
        panic(err)
    }

    hosts, err := c.GetHosts()
    if err != nil {
        return err
    }

    c.view.SetHosts(hosts)
    return nil

}

func (c *AppController) GetHosts() ([]model.Host, error) {
    hosts, err := c.model.Host.GetHosts()
    if err != nil {
        return nil, err
    }
    return hosts, nil
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
