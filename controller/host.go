package controller

import "github.com/nealwp/callapitter/model"

func (c *AppController) SetHosts() {
	hosts, err := c.GetHosts()

	if err != nil {
		c.view.SetStatus(err.Error())
	}

	c.view.SetHosts(hosts)
}

func (c *AppController) CreateHost(host model.Host) {

	err := c.model.Host.InsertHost(host)
	if err != nil {
		c.view.SetStatus(err.Error())
		return
	}

	hosts, err := c.GetHosts()
	if err != nil {
		c.view.SetStatus(err.Error())
		return
	}

	c.view.SetHosts(hosts)
}

func (c *AppController) GetHosts() ([]model.Host, error) {
	hosts, err := c.model.Host.GetHosts()
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func (c *AppController) AddHost() {
    c.view.SetStatus("Enter new host name: ")

    cb := func(host string) {
        if host != "" {
            c.CreateHost(model.Host{Name: host})
            c.view.SetStatus("Host " + host + " created")
        } else {
            c.view.SetStatus("")
        }

        c.app.SetFocus(c.view.HostDropdown.GetPrimitive())
    }

    c.view.OnStatusInputSubmit(cb)
    c.app.SetFocus(c.view.GetStatusBar())
}

func (c *AppController) DeleteHost(host model.Host) {
    
}
