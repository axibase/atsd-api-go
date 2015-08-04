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
	"strings"
)

type Query struct {
	hasStartTime bool
	startTime    int64
	hasEndTime   bool
	endTime      int64
	hasStartDate bool
	startDate    string
	hasEndDate   bool
	endDate      string
	hasLimit     bool
	limit        uint64
	entity       string
	metric       string
	tags         map[string][]string
}

func NewQuery() *Query {
	return &Query{tags: make(map[string][]string)}
}
func (self *Query) SetEntity(entityName string) *Query {
	self.entity = strings.ToLower(entityName)
	return self
}
func (self *Query) SetMetric(metricName string) *Query {
	self.metric = strings.ToLower(metricName)
	return self
}
func (self *Query) SetLimit(limit uint64) *Query {
	self.hasLimit = true
	self.limit = limit
	return self
}
func (self *Query) SetEndTime(endTime int64) *Query {
	self.hasEndTime = true
	self.endTime = endTime
	return self
}
func (self *Query) SetStartTime(startTime int64) *Query {
	self.hasStartTime = true
	self.startTime = startTime
	return self
}
func (self *Query) SetEndDate(endDate string) *Query {
	self.hasEndDate = true
	self.endDate = endDate
	return self
}
func (self *Query) SetStartDate(startDate string) *Query {
	self.hasStartDate = true
	self.startDate = startDate
	return self
}
func (self *Query) SetTag(name, value string) *Query {
	self.tags[strings.ToLower(name)] = append(self.tags[name], value)
	return self
}
func (self *Query) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"entity": self.entity,
		"metric": self.metric,
		"tags":   self.tags,
	}

	if self.hasStartTime {
		m["startTime"] = self.startTime
	}
	if self.hasEndTime {
		m["endTime"] = self.endTime
	}
	if self.hasStartDate {
		m["startDate"] = self.startDate
	}
	if self.hasEndDate {
		m["endDate"] = self.endDate
	}
	if self.hasLimit {
		m["limit"] = self.limit
	}

	return json.Marshal(m)
}
