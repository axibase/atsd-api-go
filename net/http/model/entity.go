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
)

type Entity struct {
	name string

	active     *bool
	expression *string

	tags map[string]string

	limit *uint64
}

func NewEntity(name string) *Entity {
	return &Entity{name: name, tags: make(map[string]string)}
}

func (self *Entity) Name() string {
	return self.name
}
func (self *Entity) Active() *bool {
	return self.active
}
func (self *Entity) Expression() *string {
	return self.expression
}
func (self *Entity) Tags() map[string]string {
	copy := map[string]string{}
	for k, v := range self.tags {
		copy[k] = v
	}
	return copy
}
func (self *Entity) Limit() *uint64 {
	return self.limit
}

func (self *Entity) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"tags": self.tags,
	}
	if self.active != nil {
		m["active"] = *self.active
	}

	if self.expression != nil {
		m["expression"] = *self.expression
	}

	if self.limit != nil {
		m["limit"] = *self.limit
	}
	return json.Marshal(m)
}
func (self *Entity) String() string {
	obj, _ := self.MarshalJSON()
	return string(obj)
}

func (self *Entity) SetName(entityName string) *Entity {
	self.name = entityName
	return self
}
func (self *Entity) SetActive(isActive bool) *Entity {
	self.active = &isActive
	return self
}
func (self *Entity) SetExpression(expression string) *Entity {
	self.expression = &expression
	return self
}
func (self *Entity) SetTag(key, val string) *Entity {
	self.tags[strings.ToLower(key)] = val
	return self
}
func (self *Entity) SetLimit(limit uint64) *Entity {
	self.limit = &limit
	return self
}
