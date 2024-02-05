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
