/*
 * gokrazy on-device API
 *
 * OpenAPI for the gokrazy on-device API
 *
 * API version: 1.4.0
 * Contact: michael+gokrazy@stapelberg.ch
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ondeviceapi

type Service struct {
	// Path to the service program
	Path string `json:"Path,omitempty"`
	// Command-line arguments with which to start the service
	Args []string `json:"Args,omitempty"`
	// TODO
	Diverted string `json:"Diverted,omitempty"`
	// Whether the service is stopped (gokrazy will not attempt to start it) or supervised
	Stopped bool `json:"Stopped,omitempty"`
	// TODO
	StartTime string `json:"StartTime,omitempty"`
}
