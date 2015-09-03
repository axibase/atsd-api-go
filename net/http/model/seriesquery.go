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
	"encoding/json"
	"strings"
	"time"
)

type SeriesQuery struct {
	startTime *time.Time
	endTime   *time.Time
	limit     *uint64
	entity    string
	metric    string
	tags      map[string][]string
}

func NewSeriesQuery() *SeriesQuery {
	return &SeriesQuery{tags: make(map[string][]string)}
}
func (self *SeriesQuery) SetEntity(entityName string) *SeriesQuery {
	self.entity = strings.ToLower(entityName)
	return self
}
func (self *SeriesQuery) SetMetric(metricName string) *SeriesQuery {
	self.metric = strings.ToLower(metricName)
	return self
}
func (self *SeriesQuery) SetLimit(limit uint64) *SeriesQuery {
	self.limit = &limit
	return self
}
func (self *SeriesQuery) SetStartTime(startTime time.Time) *SeriesQuery {
	self.startTime = &startTime
	return self
}
func (self *SeriesQuery) SetEndTime(endTime time.Time) *SeriesQuery {
	self.endTime = &endTime
	return self
}
func (self *SeriesQuery) SetTag(name, value string) *SeriesQuery {
	lowCaseName := strings.ToLower(name)
	self.tags[lowCaseName] = append(self.tags[lowCaseName], value)
	return self
}
func (self *SeriesQuery) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"entity": self.entity,
		"metric": self.metric,
		"tags":   self.tags,
	}

	if self.startTime != nil {
		m["startTime"] = self.startTime.UnixNano() / 1e6
	}
	if self.endTime != nil {
		m["endTime"] = self.endTime.UnixNano() / 1e6
	}
	if self.limit != nil {
		m["limit"] = self.limit
	}
	return json.Marshal(m)
}
