/*
 * router API
 *
 * router API generated from router.yang
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package router

type Route struct {
	// Destination network IP
	Network string `json:"network,omitempty"`
	// Next hop; if destination is local will be shown 'local' instead of the ip address
	Nexthop string `json:"nexthop,omitempty"`
	// Outgoing interface
	Interface_ string `json:"interface,omitempty"`
	// Cost of this route
	Pathcost int32 `json:"pathcost,omitempty"`
}
