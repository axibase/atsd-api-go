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

package model

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Sample struct {
	T uint64  `json:"t"`
	V float64 `json:"v"`
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
func (self *Series) AddTags(tags map[string]string) *Series {
	for name, val := range tags {
		self.Tags[name] = val
	}
	return self
}
func (self *Series) AddTag(name, value string) *Series {
	self.Tags[strings.ToLower(name)] = value
	return self
}
func (self *Series) AddSample(t uint64, v float64) *Series {
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
		data = data + "{t: " + strconv.FormatUint(sample.T, 10) + ", v: " + strconv.FormatFloat(sample.V, 'f', -1, 64) + "}, "
	}
	data = data + "]"
	return "{entity: \"" + self.Entity + "\", metric: \"" + self.Metric + "\", tags: " + tags + ", data: " + data + "}"
}
