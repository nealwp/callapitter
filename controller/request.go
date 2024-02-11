package controller

import (
	"os"
	"os/exec"

	"github.com/nealwp/callapitter/model"
)

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

func (c *AppController) EditRequestBody(req model.Request) {

    tmpfile, err := os.CreateTemp("", "tmp-*.json")
    if err != nil {
        c.view.SetStatus(err.Error())
		return
    }

    defer os.Remove(tmpfile.Name())

    _, err = tmpfile.Write([]byte(req.Body.String))
    if err != nil {
        c.view.SetStatus(err.Error())
        tmpfile.Close()
		return
    }

    tmpfile.Close()

    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "nvim"
    }

    success := true

    c.app.Suspend(func() {
        cmd := exec.Command(editor, tmpfile.Name())
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        err := cmd.Run()
        if err != nil {
           success = false 
        }
    })

    if !success {
        return 
    }

    updatedContent, err := os.ReadFile(tmpfile.Name())

    if err != nil {
        panic(err)
    }

    req.Body.String = string(updatedContent)

    c.UpdateRequest(req)
}
