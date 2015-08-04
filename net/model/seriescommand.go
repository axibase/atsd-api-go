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
	"bytes"
	"fmt"
)

type SeriesCommand struct {
	hasTimestamp bool
	timestamp    uint64
	entity       string
	metricValues map[string]float64
	tags         map[string]string
}

func NewSeriesCommand(entity, metricName string, metricValue float64) *SeriesCommand {
	return &SeriesCommand{entity: entity, metricValues: map[string]float64{metricName: metricValue}, tags: map[string]string{}}
}
func (self *SeriesCommand) Metrics() map[string]float64 {
	copy := map[string]float64{}
	for k, v := range self.metricValues {
		copy[k] = v
	}
	return copy
}
func (self *SeriesCommand) Entity() string {
	return self.entity
}
func (self *SeriesCommand) HasTimestamp() bool {
	return self.hasTimestamp
}
func (self *SeriesCommand) Timestamp() uint64 {
	return self.timestamp
}
func (self *SeriesCommand) Tags() map[string]string {
	copy := map[string]string{}
	for k, v := range self.tags {
		copy[k] = v
	}
	return copy
}
func (self *SeriesCommand) SetTimestamp(millis uint64) *SeriesCommand {
	self.timestamp = millis
	self.hasTimestamp = true
	return self
}
func (self *SeriesCommand) SetMetricValue(metric string, value float64) *SeriesCommand {
	self.metricValues[metric] = value
	return self
}
func (self *SeriesCommand) SetTag(tag, value string) *SeriesCommand {
	self.tags[tag] = value
	return self
}
func (self *SeriesCommand) String() string {

	msg := bytes.NewBufferString("")
	fmt.Fprintf(msg, "series e:%v", self.entity)
	if self.hasTimestamp {
		fmt.Fprintf(msg, " ms:%v", self.timestamp)
	}
	for key, val := range self.tags {
		fmt.Fprintf(msg, " t:%v=%v", key, val)
	}
	for key, val := range self.metricValues {
		fmt.Fprintf(msg, " m:%v=%v", key, val)
	}
	return msg.String()
}
