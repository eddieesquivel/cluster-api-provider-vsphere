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

// InlineResponse20010 struct for InlineResponse20010
type InlineResponse20010 struct {
	Version int32 `json:"_version,omitempty"`
	// HAProxy backend servers array
	Data []Server `json:"data"`
}
