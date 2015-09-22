/*
* Copyright 2015 Axibase Corporation or its affiliates. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License").
* You may not use this file except in compliance with the License.
* A copy of the License is located at
*
* https://www.axibase.com/atsd/axibase-apache-2.0.pdf
*
* or in the "license" file accompanying this file. This file is distributed
* on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
* express or implied. See the License for the specific language governing
* permissions and limitations under the License.
 */

package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/axibase/atsd-api-go/net/http/model"
	"github.com/golang/glog"
	"net/url"
	"strconv"
)

const (
	seriesQueryPath  = "/api/v1/series"
	seriesInsertPath = "/api/v1/series/insert"

	messagesQueryPath  = "/api/v1/messages"
	messagesInsertPath = "/api/v1/messages/insert"

	propertiesInsertPath = "/api/v1/properties/insert"

	entitiesPath = "/api/v1/entities"

	metricsPath = "/api/v1/metrics"
	commandPath = "/api/v1/command"
)

type Client struct {
	url      *url.URL
	username string
	password string

	Series     *Series
	Properties *Properties
	Entities   *Entities
	Messages   *Messages

	Metric *Metric

	httpClient *http.Client
}

func New(mUrl url.URL, username, password string) *Client {
	var client = Client{url: &mUrl, username: username, password: password}
	client.Series = &Series{&client}
	client.Properties = &Properties{&client}
	client.Entities = &Entities{&client}
	client.Messages = &Messages{&client}
	client.Metric = &Metric{&client}
	client.httpClient = &http.Client{}
	return &client
}

func (self *Client) Url() url.URL {
	return *self.url
}
func (self *Client) request(reqType, apiUrl string, reqJson []byte) (string, error) {
	req, err := http.NewRequest(reqType, self.url.String(), bytes.NewReader(reqJson))
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

func (self *Series) Query(queries []*model.SeriesQuery) ([]*model.Series, error) {

	request := struct {
		Queries []*model.SeriesQuery `json:"queries"`
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
func (self *Entities) List(active bool, expression string, tags []string, limit uint64) ([]*model.Entity, error) {
	tagsParams := ""
	if len(tags) == 1 && tags[0] == "*" {
		tagsParams = url.QueryEscape("*")
	} else {
		for i, tag := range tags {
			if i == 0 {
				tagsParams += url.QueryEscape(tag)
			} else {
				tagsParams += "," + url.QueryEscape(tag)
			}
		}
	}

	jsonData, err := self.client.request("GET", entitiesPath+"?"+
		"tags="+ tagsParams+"&"+
		"active=" + url.QueryEscape(strconv.FormatBool(active)) + "&"+
		"expression=" + url.QueryEscape(expression)+"&"+
	    "limit=" + url.QueryEscape(strconv.FormatUint(limit, 10)), []byte{})
	if err != nil {
		return nil, err
	}

	var entities []*model.Entity
	err = json.Unmarshal([]byte(jsonData), &entities)
	if err != nil {
		panic(err)
	}

	return entities, nil
}

type Metric struct {
	client *Client
}

func (self *Metric) CreateOrReplace(metric *model.Metric) error {
	jsonRequest, err := json.Marshal(metric)
	if err != nil {
		panic(err)
	}
	metricName := url.QueryEscape(metric.Name())
	_, err = self.client.request("PUT", metricsPath+"/"+metricName, jsonRequest)
	if err != nil {
		return err
	}
	return nil
}

type Messages struct {
	client *Client
}

func (self Messages) Insert(messages []*model.Message) error {
	jsonRequest, err := json.Marshal(messages)
	if err != nil {
		panic(err)
	}
	_, err = self.client.request("POST", messagesInsertPath, jsonRequest)
	if err != nil {
		return err
	}
	return nil
}
func (self Messages) Query(query *model.MessagesQuery) ([]*model.Message, error) {
	jsonRequest, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}
	jsonData, err := self.client.request("POST", messagesQueryPath, jsonRequest)
	if err != nil {
		return nil, err
	}
	var messages []*model.Message
	err = json.Unmarshal([]byte(jsonData), &messages)
	if err != nil {
		panic(err)
	}
	return messages, nil
}
