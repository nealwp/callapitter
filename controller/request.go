package controller

import "github.com/nealwp/callapitter/model"

func (c *AppController) SendRequest(i int) {

	host := c.view.GetSelectedHost()
	req := c.model.Request.GetRequest(i)

	res, err := sendHttpRequest(req, host)

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
