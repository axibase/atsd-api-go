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
	"fmt"
)

type Property struct {
	propType     string
	entity       string
	key          map[string]string
	tags         map[string]string
	hasTimestamp bool
	timestamp    uint64
}

func NewProperty(propType, entity string) *Property {
	return &Property{propType: propType, entity: entity, key: map[string]string{}, tags: map[string]string{}}
}
func (self *Property) SetKeyPart(name, value string) *Property {
	self.key[name] = value
	return self
}
func (self *Property) SetKey(key map[string]string) *Property {
	self.key = key
	return self
}

func (self *Property) SetAllTags(tags map[string]string) *Property {
	self.tags = tags
	return self
}
func (self *Property) SetTag(name, value string) *Property {
	self.tags[name] = value
	return self
}
func (self *Property) SetTimestamp(timestamp uint64) *Property {
	self.hasTimestamp = true
	self.timestamp = timestamp
	return self
}
func (self *Property) PropType() string {
	return self.propType
}
func (self *Property) Entity() string {
	return self.entity
}
func (self *Property) Key() map[string]string {
	copy := map[string]string{}
	for k, v := range self.key {
		copy[k] = v
	}
	return copy
}
func (self *Property) Tags() map[string]string {
	copy := map[string]string{}
	for k, v := range self.tags {
		copy[k] = v
	}
	return copy
}
func (self *Property) HasTimestamp() bool {
	return self.hasTimestamp
}
func (self *Property) Timestamp() uint64 {
	self.hasTimestamp = true
	return self.timestamp
}
func (self *Property) MarshalJSON() ([]byte, error) {
	key, _ := json.Marshal(self.key)
	tags, _ := json.Marshal(self.tags)
	return []byte(fmt.Sprintf("{\"type\":\"%v\",\"entity\":\"%v\",\"key\":%v,\"tags\":%v,\"timestamp\":%v}", self.propType, self.entity, string(key), string(tags), self.timestamp)), nil
}
func (self *Property) String() string {
	str := fmt.Sprintf("property e:%v t:%v", self.entity, self.propType)
	if self.HasTimestamp() {
		str += fmt.Sprintf(" ms:%v", self.timestamp)
	}
	keyString := ""
	for key, val := range self.key {
		keyString += fmt.Sprintf(" k:%v=%v", key, val)
	}
	str += keyString
	tagsString := ""
	for key, val := range self.tags {
		tagsString += fmt.Sprintf(" v:%v=%v", key, val)
	}
	str += tagsString

	return str
}
