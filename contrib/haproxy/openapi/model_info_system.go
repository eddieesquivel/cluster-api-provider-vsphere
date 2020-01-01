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
