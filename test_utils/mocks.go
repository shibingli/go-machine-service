package test_utils

import (
	"github.com/rancherio/go-rancher/client"
)

type MockMachineHostClient struct {
	MachineHost *client.MachineHost
}

func (c *MockMachineHostClient) ById(id string) (*client.MachineHost, error) {
	if c.MachineHost == nil {
		host := &client.MachineHost{
			ExternalId:       "ext-" + id,
			Kind:             "machineHost",
			Driver:           "VirtualBox",
			VirtualboxConfig: client.VirtualboxConfig{},
		}
		host.Id = id
		return host, nil
	}

	return c.MachineHost, nil
}

func (c *MockMachineHostClient) Create(container *client.MachineHost) (*client.MachineHost, error) {
	return nil, nil
}

func (c *MockMachineHostClient) Update(existing *client.MachineHost, updates interface{}) (*client.MachineHost, error) {
	return nil, nil
}

func (c *MockMachineHostClient) List(opts *client.ListOpts) (*client.MachineHostCollection, error) {
	return nil, nil
}

func (c *MockMachineHostClient) Delete(container *client.MachineHost) error {
	return nil
}

type MockRegistrationTokenClient struct{}

func (c *MockRegistrationTokenClient) Create(container *client.RegistrationToken) (*client.RegistrationToken, error) {
	resp := &client.RegistrationToken{}
	return resp, nil
}

func (c *MockRegistrationTokenClient) Update(existing *client.RegistrationToken,
	updates interface{}) (*client.RegistrationToken, error) {
	resp := &client.RegistrationToken{}
	return resp, nil
}

func (c *MockRegistrationTokenClient) List(opts *client.ListOpts) (*client.RegistrationTokenCollection, error) {
	regToken := client.RegistrationToken{}
	regToken.Links = map[string]string{"registrationUrl": "http://1.2.3.4/v1"}

	tokens := []client.RegistrationToken{regToken}
	resp := &client.RegistrationTokenCollection{
		Data: tokens,
	}

	return resp, nil
}

func (c *MockRegistrationTokenClient) ById(id string) (*client.RegistrationToken, error) {
	resp := &client.RegistrationToken{}
	return resp, nil
}

func (c *MockRegistrationTokenClient) Delete(container *client.RegistrationToken) error {
	return nil
}