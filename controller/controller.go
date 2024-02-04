package controller

import (
	"fmt"

	"github.com/nealwp/callapitter/model"
	"github.com/nealwp/callapitter/ui"
)

type AppController struct {
    model *model.AppModel
    view *ui.AppLayout
}

func NewAppController() *AppController {
    return &AppController{}
}

func (c *AppController) SetView(view *ui.AppLayout) {
    c.view = view
}

func (c *AppController) SetModel(model *model.AppModel) {
    c.model = model
}

func (c *AppController) SendRequest(req model.Request) {
    fmt.Println("request sent")
}

func (c *AppController) CreateRequest(req model.Request) error {

    err := c.model.Request.InsertRequest(req) 
    if err != nil {
        return err
    }

    requests, err := c.GetRequests()
    if err != nil {
        return err
    }

    c.view.SetRequests(requests)
    return nil
}

func (c *AppController) DeleteRequest(req model.Request) error {
    err := c.model.Request.DeleteRequest(req)

    if err != nil {
        return err
    }

    requests, err := c.GetRequests()
    if err != nil {
        return err
    }

    c.view.SetRequests(requests)

    return nil
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

    requests, err := c.GetRequests()
    if err != nil {
        panic(err)
    }

    c.view.SetRequests(requests)
}
