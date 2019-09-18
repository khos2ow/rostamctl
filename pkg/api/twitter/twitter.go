// Copyright © 2019 Khosrow Moossavi.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package twitter contains implementation of Twitter API client
package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"

	resty "github.com/go-resty/resty/v2"
	"github.com/khos2ow/rostamctl/pkg/api"
)

type blocked struct {
	Blocked []account `json:"blockedAccounts"`
}

type account struct {
	ID   int64  `json:"twitterUserId"`
	Name string `json:"twitterScreenName"`
}

// twitter is the implementation of api.Client for Twitter
// API endpoint
type twitter struct {
	baseurl string
}

// NewClient returns a new API Client for Twitter endpoint
func NewClient() api.Client {
	return &twitter{
		baseurl: fmt.Sprintf("%s/twitter", api.APIBaseURL),
	}
}

// Get returns a blocked Twitter account based on the provided
// name, if found.
func (t *twitter) Check(name string) (*api.Account, error) {
	response, err := resty.New().R().Get(fmt.Sprintf("%s/check/%s", t.baseurl, name))
	if err != nil {
		return nil, err
	}
	value, _ := strconv.ParseBool(string(response.Body()))
	account := &api.Account{
		Blocked: value,
	}
	return account, nil
}

// Get returns a blocked Twitter account based on the provided
// name, if found.
func (t *twitter) Get(name string) (*api.Account, error) {
	data, err := t.listAll()
	if err != nil {
		return nil, err
	}

	for _, item := range data.Blocked {
		if item.Name == name {
			account := &api.Account{
				ID:      item.ID,
				Name:    item.Name,
				Blocked: true,
			}
			return account, nil
		}
	}

	account := &api.Account{
		Name:    name,
		Blocked: false,
	}
	return account, nil
}

// List returns list of blocked Twitter accounts
func (t *twitter) List() ([]*api.Account, error) {
	data, err := t.listAll()
	if err != nil {
		return nil, err
	}

	list := []*api.Account{}
	for _, item := range data.Blocked {
		list = append(list, &api.Account{
			ID:      item.ID,
			Name:    item.Name,
			Blocked: true,
		})
	}

	return list, nil
}

func (t *twitter) listAll() (*blocked, error) {
	response, err := resty.New().R().Get(fmt.Sprintf("%s/list", t.baseurl))
	if err != nil {
		return nil, err
	}

	data := &blocked{}
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
