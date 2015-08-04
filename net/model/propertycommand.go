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

type PropertyCommand struct {
	propType     string
	entity       string
	key          map[string]string
	tags         map[string]string
	hasTimestamp bool
	timestamp    uint64
}

func NewPropertyCommand(propType, entity string) *PropertyCommand {
	return &PropertyCommand{propType: propType, entity: entity, key: map[string]string{}, tags: map[string]string{}}
}
func (self *PropertyCommand) SetKey(key map[string]string) *PropertyCommand {
	self.key = key
	return self
}
func (self *PropertyCommand) SetKeyPart(name, value string) *PropertyCommand {
	self.key[name] = value
	return self
}
func (self *PropertyCommand) SetAllTags(tags map[string]string) *PropertyCommand {
	self.tags = tags
	return self
}
func (self *PropertyCommand) SetTag(name, value string) *PropertyCommand {
	self.tags[name] = value
	return self
}
func (self *PropertyCommand) SetTimestamp(timestamp uint64) *PropertyCommand {
	self.hasTimestamp = true
	self.timestamp = timestamp
	return self
}
func (self *PropertyCommand) PropType() string {
	return self.propType
}
func (self *PropertyCommand) Entity() string {
	return self.entity
}
func (self *PropertyCommand) Key() map[string]string {
	copy := map[string]string{}
	for k, v := range self.key {
		copy[k] = v
	}
	return copy
}
func (self *PropertyCommand) Tags() map[string]string {
	copy := map[string]string{}
	for k, v := range self.tags {
		copy[k] = v
	}
	return copy
}
func (self *PropertyCommand) HasTimestamp() bool {
	return self.hasTimestamp
}
func (self *PropertyCommand) Timestamp() uint64 {
	self.hasTimestamp = true
	return self.timestamp
}
func (self *PropertyCommand) String() string {
	str := bytes.NewBufferString("")
	fmt.Fprintf(str, "property e:%v t:%v", self.entity, self.propType)
	if self.HasTimestamp() {
		fmt.Fprintf(str, " ms:%v", self.timestamp)
	}
	for key, val := range self.key {
		fmt.Fprintf(str, " k:%v=%v", key, val)
	}
	for key, val := range self.tags {
		fmt.Fprintf(str, " v:%v=%v", key, val)
	}

	return str.String()
}
