/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package registry

import (
	"fmt"

	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro"
)

type Environment struct {
	Name               string                   `json:"name" yaml:"name" mapstructure:"name"`                                           // Target environment name, such as "Production"
	Type               string                   `json:"type" yaml:"type" mapstructure:"type"`                                           // Target type: "StandAlone", "HighAvailabilityPair", "Cluster"
	Snip               Node                     `json:"snip" yaml:"snip" mapstructure:"snip"`                                           // Connection details for the shared Address (SNIP) of the environment
	Nodes              []Node                   `json:"nodes" yaml:"nodes" mapstructure:"nodes"`                                        // Connection details for the individual Nodes of each node
	Credentials        nitro.Credentials        `json:"credentials" yaml:"credentials" mapstructure:"credentials"`                      // Credentials
	ConnectionSettings nitro.ConnectionSettings `json:"connectionSettings" yaml:"connectionSettings" mapstructure:"connectionSettings"` // Connections settings
}

func (e *Environment) GetNitroClientForNode(nodeName string) (*nitro.Client, error) {
	var (
		err    error
		client *nitro.Client
	)

	for _, n := range e.GetNodes() {
		if n.Name == nodeName {
			client, err = nitro.NewClient(n.Name, n.Address, e.Credentials, e.ConnectionSettings)
			if err != nil {
				return nil, fmt.Errorf("could not create client for node %s with error %w", nodeName, err)
			}
			return client, nil
		}
	}
	return nil, fmt.Errorf("could not create client for node %s with error: node not found in environment %s", nodeName, e.Name)
}

func (e *Environment) GetNitroClientForSnip() (*nitro.Client, error) {
	// Return the SNIP Node if defined in the environment
	if !e.HasSnip() {
		return nil, fmt.Errorf("no SNIP node defined for environment %s", e.Name)

	}

	client, err := e.GetNitroClientForNode(e.Snip.Name)
	if err != nil {
		return nil, fmt.Errorf("could not create client for SNIP node %s for environment %s with error %w", e.Snip.Name, e.Name, err)
	}
	return client, nil
}

func (e *Environment) GetNodeNames() []string {
	var output []string
	for _, n := range e.GetNodes() {
		output = append(output, n.Name)
	}
	return output
}

func (e *Environment) GetNodes() []Node {
	var output []Node
	output = append(output, e.Nodes...)

	if e.HasSnip() {
		output = append(output, e.Snip)
	}
	return output
}

func (e *Environment) GetPrimaryNitroClient() (*nitro.Client, error) {
	var (
		err    error
		client *nitro.Client
	)

	client, _ = e.GetNitroClientForSnip()
	if client != nil {
		return client, nil
	}

	if !e.HasNodes() {
		return nil, fmt.Errorf("no individual nodes defined for environment %s", e.Name)
	}

	// Loop over individual nodes to determine which one is primary
	// It is not guaranteed in a High-Available environment that one of the nodes is automatically primary, so we he have to iterate over both
	for _, n := range e.Nodes {
		client, err = e.GetNitroClientForNode(n.Name)
		if err != nil {
			return nil, fmt.Errorf("could not create client for node %s to determine status as primary node for environment %s with error %w", n.Name, e.Name, err)
		}

		var isPrimary bool
		isPrimary, err = client.IsPrimaryNode()
		if isPrimary {
			return client, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not determine status for node %s for environment %s with error %w", n.Name, e.Name, err)
		}
	}

	return nil, fmt.Errorf("could not find a primary node for environment %s", e.Name)
}

func (e *Environment) HasNodes() bool {
	if len(e.Nodes) == 0 {
		return false
	}
	return true
}

func (e *Environment) HasSnip() bool {
	emptyNode := Node{}
	if e.Snip != emptyNode {
		return true
	}
	return false
}
