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

// Package api contains the definition of API client used in rostamctl
package api

// nolint
const (
	APIBaseURL = "https://rostambot.com/api/v1"
)

// Account is a normalized representation of blocked accounts.
type Account struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Blocked bool   `json:"blocked"`
}

// Client abstracts over different API Client implementations
type Client interface {
	Check(name string) (*Account, error)
	Get(name string) (*Account, error)
	List() ([]*Account, error)
}
