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

// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	Version int32 `json:"_version,omitempty"`
	// HAProxy frontend binds array (corresponds to bind directives)
	Data []Bind `json:"data"`
}
