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

import "encoding/json"

type Entity struct {
	name string

	hasActive bool
	active    bool

	hasExpression bool
	expression    string

	tags map[string]string

	hasLimit bool
	limit    uint64
}

func NewEntity(name string) *Entity {
	return &Entity{name: name, tags: make(map[string]string)}
}

func (self *Entity) Name() string {
	return self.name
}

func (self *Entity) HasActive() bool {
	return self.hasActive
}
func (self *Entity) Active() bool {
	return self.active
}

func (self *Entity) HasExpression() bool {
	return self.hasExpression
}
func (self *Entity) Expression() string {
	return self.expression
}

func (self *Entity) Tags() map[string]string {
	copy := map[string]string{}
	for k, v := range self.tags {
		copy[k] = v
	}
	return copy
}

func (self *Entity) HasLimit() bool {
	return self.hasLimit
}
func (self *Entity) Limit() uint64 {
	return self.limit
}

func (self *Entity) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"tags": self.tags,
	}

	if self.HasActive() {
		m["active"] = self.active
	}

	if self.HasExpression() {
		m["expression"] = self.expression
	}

	if self.HasLimit() {
		m["limit"] = self.limit
	}
	return json.Marshal(m)
}

func (self *Entity) SetName(entityName string) *Entity {
	self.name = entityName
	return self
}
func (self *Entity) SetActive(isActive bool) *Entity {
	self.hasActive = true
	self.active = isActive
	return self
}
func (self *Entity) SetExpression(expression string) *Entity {
	self.hasExpression = true
	self.expression = expression
	return self
}
func (self *Entity) SetTag(key, val string) *Entity {
	self.tags[key] = val
	return self
}
func (self *Entity) SetLimit(limit uint64) *Entity {
	self.hasLimit = true
	self.limit = limit
	return self
}
