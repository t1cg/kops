/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ""fitask" -type=RouteTable"; DO NOT EDIT

package azuretasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// RouteTable

// JSON marshaling boilerplate
type realRouteTable RouteTable

// UnmarshalJSON implements conversion to JSON, supporting an alternate specification of the object as a string
func (o *RouteTable) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realRouteTable
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = RouteTable(r)
	return nil
}

var _ fi.HasLifecycle = &RouteTable{}

// GetLifecycle returns the Lifecycle of the object, implementing fi.HasLifecycle
func (o *RouteTable) GetLifecycle() *fi.Lifecycle {
	return o.Lifecycle
}

// SetLifecycle sets the Lifecycle of the object, implementing fi.SetLifecycle
func (o *RouteTable) SetLifecycle(lifecycle fi.Lifecycle) {
	o.Lifecycle = &lifecycle
}

var _ fi.HasName = &RouteTable{}

// GetName returns the Name of the object, implementing fi.HasName
func (o *RouteTable) GetName() *string {
	return o.Name
}

// SetName sets the Name of the object, implementing fi.SetName
func (o *RouteTable) SetName(name string) {
	o.Name = &name
}

// String is the stringer function for the task, producing readable output using fi.TaskAsString
func (o *RouteTable) String() string {
	return fi.TaskAsString(o)
}
