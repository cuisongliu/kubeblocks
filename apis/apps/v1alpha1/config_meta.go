/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

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

package v1alpha1

import (
	"encoding/json"
)

// MarshalJSON implements the Marshaler interface.
func (c *Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Data)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (c *Payload) UnmarshalJSON(data []byte) error {
	var out map[string]interface{}
	err := json.Unmarshal(data, &out)
	if err != nil {
		return err
	}
	c.Data = out
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
// This exists here to work around https://github.com/kubernetes/code-generator/issues/50
func (c *Payload) DeepCopyInto(out *Payload) {
	bytes, err := json.Marshal(c.Data)
	if err != nil {
		// TODO how to process error: panic or ignore
		return // ignore
	}
	var clone map[string]interface{}
	err = json.Unmarshal(bytes, &clone)
	if err != nil {
		// TODO how to process error: panic or ignore
		return // ignore
	}
	out.Data = clone
}

func (configuration *ConfigurationSpec) GetConfigurationItem(name string) *ConfigurationItemDetail {
	for i := range configuration.ConfigItemDetails {
		configItem := &configuration.ConfigItemDetails[i]
		if configItem.Name == name {
			return configItem
		}
	}
	return nil
}

func (configuration *ConfigurationSpec) GetConfigSpec(configSpecName string) *ComponentConfigSpec {
	if configItem := configuration.GetConfigurationItem(configSpecName); configItem != nil {
		return configItem.ConfigSpec
	}
	return nil
}

func (status *ConfigurationStatus) GetItemStatus(name string) *ConfigurationItemDetailStatus {
	for i := range status.ConfigurationItemStatus {
		itemStatus := &status.ConfigurationItemStatus[i]
		if itemStatus.Name == name {
			return itemStatus
		}
	}
	return nil
}

func (configSpec *ComponentConfigSpec) InjectEnvEnabled() bool {
	return len(configSpec.AsEnvFrom) > 0 || len(configSpec.InjectEnvTo) > 0
}

func (configSpec *ComponentConfigSpec) ToSecret() bool {
	return configSpec.AsSecret != nil && *configSpec.AsSecret
}

func (configSpec *ComponentConfigSpec) ContainersInjectedTo() []string {
	if len(configSpec.InjectEnvTo) != 0 {
		return configSpec.InjectEnvTo
	}
	return configSpec.AsEnvFrom
}
