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

const (
	commandName = "entity-tag"
)

type EntityTagCommand struct {
	entity string
	tags   map[string]string
}

func NewEntityTagCommand(entity, tagName, tagValue string) *EntityTagCommand {
	return &EntityTagCommand{entity: entity, tags: map[string]string{tagName: tagValue}}
}

func (self *EntityTagCommand) Entity() string {
	return self.entity
}
func (self *EntityTagCommand) Tags() map[string]string {
	copy := map[string]string{}
	for k, v := range self.tags {
		copy[k] = v
	}
	return copy
}

func (self *EntityTagCommand) AddTag(name, value string) *EntityTagCommand {
	self.tags[name] = value
	return self
}
func (self *EntityTagCommand) String() string {
	result := bytes.NewBufferString(commandName)
	fmt.Fprintf(result, " e:%v", self.entity)
	for name, value := range self.tags {
		fmt.Fprintf(result, " t:%v=%v", name, value)
	}
	return result.String()
}
