/*
 * GUS (gokrazy update service)
 *
 * OpenAPI for GUS (gokrazy update service)
 *
 * API version: 1.4.0
 * Contact: michael+gokrazy@stapelberg.ch
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gusapi

// TODO
type HeartbeatRequestSbom struct {
	ConfigHash *PathHash `json:"config_hash,omitempty"`
	// TODO
	GoModHashes []PathHash `json:"go_mod_hashes,omitempty"`
}