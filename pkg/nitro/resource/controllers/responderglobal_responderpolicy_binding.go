package controllers

import (
	"net/http"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type ResponderGlobalResponderPolicyBindingController struct {
	client *nitro.Client
}

func NewResponderGlobalResponderPolicyBindingController(client *nitro.Client) *ResponderGlobalResponderPolicyBindingController {
	c := ResponderGlobalResponderPolicyBindingController{
		client: client,
	}
	return &c
}

func (c *ResponderGlobalResponderPolicyBindingController) Add(name string, bindType string, priority string, gotoPriorityExpression string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method: http.MethodPost,
		Data: []config.ResponderGlobalResponderPolicyBinding{
			config.NewResponderGlobalResponderPolicyBindingAddRequest(name, bindType, priority, gotoPriorityExpression),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) AddRequest(r nitro.Request[config.ResponderGlobalResponderPolicyBinding]) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) Count() (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) Delete(name string, bindType string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method: http.MethodDelete,
		Arguments: map[string]string{
			"policyname": name,
			"type":       bindType,
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) Get(name string, attributes []string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) List(filter map[string]string, attributes []string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.executeNitroRequest(c.client, &r)
}
