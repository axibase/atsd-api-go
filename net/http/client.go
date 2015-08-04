// Copyright 2014 Google Inc. All Rights Reserved.
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

package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"github.com/google/cadvisor/storage/atsd/net/http/model"
	"net/url"
)

const (
	seriesQueryPath  = "/api/v1/series"
	seriesInsertPath = "/api/v1/series/insert"

	propertiesInsertPath = "/api/v1/properties/insert"

	entitiesPath = "/api/v1/entities"
)

type Client struct {
	url      string
	username string
	password string

	Series     *Series
	Properties *Properties
	Entities   *Entities

	httpClient *http.Client
}

func New(url, username, password string) *Client {
	var client = Client{url: url, username: username, password: password}
	client.Series = &Series{&client}
	client.Properties = &Properties{&client}
	client.Entities = &Entities{&client}
	client.httpClient = &http.Client{}
	return &client
}

func (self *Client) request(reqType, apiUrl string, reqJson []byte) (string, error) {
	req, err := http.NewRequest(reqType, self.url, bytes.NewReader(reqJson))
	req.URL.Opaque = req.URL.Path + apiUrl
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(self.username, self.password)
	res, err := self.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var error struct {
		Error string `json:"error"`
	}

	_ = json.Unmarshal(jsonData, &error)

	if error.Error != "" {
		return string(jsonData), errors.New(error.Error)
	}

	return string(jsonData), nil
}

type Series struct {
	client *Client
}

func (self *Series) Query(queries []*model.Query) ([]*model.Series, error) {

	request := struct {
		Queries []*model.Query `json:"queries"`
	}{queries}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	jsonData, err := self.client.request("POST", seriesQueryPath, jsonRequest)
	if err != nil {
		return nil, err
	}
	var series struct {
		Series []*model.Series `json:"series"`
	}
	err = json.Unmarshal([]byte(jsonData), &series)
	if err != nil {
		panic(err)
	}
	for _, s := range series.Series {
		if s.Warning != "" {
			glog.Warning(s.Warning)
		}
	}
	return series.Series, nil
}

func (self *Series) Insert(series []*model.Series) error {
	jsonSeries, err := json.Marshal(series)
	if err != nil {
		panic(err)
	}
	_, err = self.client.request("POST", seriesInsertPath, jsonSeries)
	if err != nil {
		return err
	}

	return nil
}

type Properties struct {
	client *Client
}

func (self *Properties) Insert(properties []*model.Property) error {
	jsonProperties, err := json.Marshal(properties)
	if err != nil {
		panic(err)
	}
	_, err = self.client.request("POST", propertiesInsertPath, jsonProperties)
	if err != nil {
		return err
	}

	return nil
}

type Entities struct {
	client *Client
}

func (self *Entities) Create(entity *model.Entity) error {
	jsonRequest, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}
	_, err = self.client.request("PUT", entitiesPath+"/"+url.QueryEscape(entity.Name()), jsonRequest)
	if err != nil {
		return err
	}
	return nil
}
func (self *Entities) Update(entity *model.Entity) error {
	jsonRequest, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}
	entityName := url.QueryEscape(entity.Name())
	_, err = self.client.request("PATCH", entitiesPath+"/"+entityName, jsonRequest)
	if err != nil {
		return err
	}
	return nil
}
