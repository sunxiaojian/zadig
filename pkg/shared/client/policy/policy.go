/*
Copyright 2021 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package policy

import (
	"fmt"

	"github.com/koderover/zadig/pkg/tool/httpclient"
	"github.com/koderover/zadig/pkg/tool/log"
)

type Policy struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rules       []*Rule `json:"rules"`
}

type Rule struct {
	Verbs           []string         `json:"verbs"`
	Resources       []string         `json:"resources"`
	Kind            string           `json:"kind"`
	MatchAttributes []MatchAttribute `json:"match_attributes"`
}

type MatchAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreatePoliciesArgs struct {
	Policies []*Policy `json:"policies"`
}

type DeletePoliciesArgs struct {
	Names []string `json:"names"`
}

func (c *Client) CreatePolicies(request CreatePoliciesArgs) error {
	url := "/policies"

	_, err := c.Post(url, httpclient.SetBody(request))
	if err != nil {
		log.Errorf("Failed to add audit log, error: %s", err)
		return err
	}

	return nil
}

func (c *Client) DeletePolicies(ns string, request DeletePoliciesArgs) error {
	url := fmt.Sprintf("/policies?projectName=%s", ns)

	_, err := c.Delete(url, httpclient.SetBody(request))
	if err != nil {
		log.Errorf("Failed to add audit log, error: %s", err)
		return err
	}

	return nil
}

func (c *Client) UpdatePolicy(ns string, policy *Policy) error {
	url := fmt.Sprintf("/policies/%s?projectName=%s", policy.Name, ns)
	_, err := c.Put(url, httpclient.SetBody(policy))
	return err
}