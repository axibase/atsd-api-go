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

package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	netModel "github.com/axibase/atsd-api-go/net/model"
	"strconv"
	"strings"
)

type Sample struct {
	T netModel.Millis `json:"t"`
	V netModel.Number `json:"v"`
}

func (self *Sample) UnmarshalJSON(data []byte) error {
	var jsonMap map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	if err := dec.Decode(&jsonMap); err != nil {
		return err
	}
	self.T, _ = jsonMap["t"].(netModel.Millis)
	fmt.Println(jsonMap)
	switch value := jsonMap["v"].(type) {
	case json.Number:
		strRep := value.String()
		if strings.Contains(strRep, ".") {
			temp, _ := value.Float64()
			self.V = netModel.Float64(temp)
		} else {
			temp, _ := value.Int64()
			self.V = netModel.Int64(temp)
		}
	default:
		panic(value)
	}
	return nil
}

type Series struct {
	Entity  string            `json:"entity"`
	Metric  string            `json:"metric"`
	Tags    map[string]string `json:"tags"`
	Warning string            `json:"warning"`
	Data    []*Sample         `json:"data"`
}

func NewSeries(entity, metric string) *Series {
	return &Series{
		Entity: strings.ToLower(entity),
		Metric: strings.ToLower(metric),
		Tags:   map[string]string{},
		Data:   []*Sample{}}
}
func (self *Series) AddMetricPrefix(prefix string) *Series {
	self.Metric = prefix + "." + self.Metric
	return self
}
func (self *Series) MarshalJSON() ([]byte, error) {
	tags, _ := json.Marshal(self.Tags)
	data, _ := json.Marshal(self.Data)
	return []byte("{\"entity\":\"" + self.Entity + "\",\"metric\":\"" + self.Metric + "\",\"tags\":" + string(tags) + ",\"data\":" + string(data) + "}"), nil
}
func (self *Series) TagValue(name string) string {
	return self.Tags[name]
}
func (self *Series) SetTags(tags map[string]string) *Series {
	for name, val := range tags {
		self.Tags[strings.ToLower(name)] = val
	}
	return self
}
func (self *Series) SetTag(name, value string) *Series {
	self.Tags[strings.ToLower(name)] = value
	return self
}
func (self *Series) AddSample(t netModel.Millis, v netModel.Number) *Series {
	self.Data = append(self.Data, &Sample{T: t, V: v})
	return self
}
func (self *Series) String() string {
	tags := "{"
	for name, value := range self.Tags {
		tags = tags + name + ": \"" + value + "\", "
	}
	tags = tags + "}"
	data := "["
	for _, sample := range self.Data {
		data = data + "{t: " + strconv.FormatUint(uint64(sample.T), 10) + ", v: " + sample.V.String() + "}, "
	}
	data = data + "]"
	return "{entity: \"" + self.Entity + "\", metric: \"" + self.Metric + "\", tags: " + tags + ", data: " + data + "}"
}
