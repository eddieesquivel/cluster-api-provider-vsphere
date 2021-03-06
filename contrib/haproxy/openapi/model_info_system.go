/*
Copyright 2020 The Kubernetes Authors.

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

/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InfoSystem struct for InfoSystem
type InfoSystem struct {
	CpuInfo InfoSystemCpuInfo `json:"cpu_info,omitempty"`
	// Hostname where the HAProxy is running
	Hostname string            `json:"hostname,omitempty"`
	MemInfo  InfoSystemMemInfo `json:"mem_info,omitempty"`
	// OS string
	OsString string `json:"os_string,omitempty"`
	// Current time in milliseconds since Epoch.
	Time int32 `json:"time,omitempty"`
	// System uptime
	Uptime *int32 `json:"uptime,omitempty"`
}
